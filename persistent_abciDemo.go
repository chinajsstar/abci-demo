package abci-demo

import (
	"bytes"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/tendermint/abci/types"
	wire "github.com/tendermint/go-wire"
	"github.com/tendermint/merkleeyes/iavl"
	cmn "github.com/tendermint/tmlibs/common"
	dbm "github.com/tendermint/tmlibs/db"
	"github.com/tendermint/tmlibs/log"
	"fmt"
)

const (
	ValidatorSetChangePrefix string = "val:"
)

//-----------------------------------------

type PersistentABCIDemoApplication struct {
	app *ABCIDemoApplication
	db  dbm.DB

	// latest received
	// TODO: move to merkle tree?
	blockHeader *types.Header

	// validator set
	changes []*types.Validator

	logger log.Logger
}

func NewPersistentABCIDemoApplication(dbDir string) *PersistentABCIDemoApplication {
	db := dbm.NewDB("account", "leveldb", dbDir)
	lastBlock := LoadLastBlock(db)

	stateTree := iavl.NewIAVLTree(0, db)
	stateTree.Load(lastBlock.AppHash)

	// log.Notice("Loaded state", "block", lastBlock.Height, "root", stateTree.Hash())

	return &PersistentABCIDemoApplication{
		app:    &ABCIDemoApplication{state: stateTree},
		db:     db,
		logger: log.NewNopLogger(),
	}
}

func (app *PersistentABCIDemoApplication) SetLogger(l log.Logger) {
	app.logger = l
}

func (app *PersistentABCIDemoApplication) Info() (resInfo types.ResponseInfo) {
	resInfo = app.app.Info()
	lastBlock := LoadLastBlock(app.db)
	resInfo.LastBlockHeight = lastBlock.Height
	resInfo.LastBlockAppHash = lastBlock.AppHash
	return resInfo
}

func (app *PersistentABCIDemoApplication) SetOption(key string, value string) (log string) {
	return app.app.SetOption(key, value)
}

// tx is either "key=value" or just arbitrary bytes
func (app *PersistentABCIDemoApplication) DeliverTx(tx []byte) types.Result {
	if isValidatorTx(tx) {
		return app.execValidatorTx(tx)
	}

	// otherwise, update the key-value store
	return app.app.DeliverTx(tx)
}

func (app *PersistentABCIDemoApplication) CheckTx(tx []byte) types.Result {
	return app.app.CheckTx(tx)
}

func (app *PersistentABCIDemoApplication) Commit() types.Result {
	// Save
	appHash := app.app.state.Save()
	app.logger.Info("Saved state", "root", fmt.Sprintf("%X", appHash))

	lastBlock := LastBlockInfo{
		Height:  app.blockHeader.Height,
		AppHash: appHash, // this hash will be in the next block header
	}

	app.logger.Info("Saving block", "height", lastBlock.Height, "root", fmt.Sprintf("%X", lastBlock.AppHash))
	SaveLastBlock(app.db, lastBlock)

	return types.NewResultOK(appHash, "")
}

func (app *PersistentABCIDemoApplication) Query(reqQuery types.RequestQuery) types.ResponseQuery {
	return app.app.Query(reqQuery)
}

// Save the validators in the merkle tree
func (app *PersistentABCIDemoApplication) InitChain(validators []*types.Validator) {
	for _, v := range validators {
		r := app.updateValidator(v)
		if r.IsErr() {
			app.logger.Error("Error updating validators", "r", r)
		}
	}
}

// Track the block hash and header information
func (app *PersistentABCIDemoApplication) BeginBlock(hash []byte, header *types.Header) {
	// update latest block info
	app.blockHeader = header

	// reset valset changes
	app.changes = make([]*types.Validator, 0)
}

// Update the validator set
func (app *PersistentABCIDemoApplication) EndBlock(height uint64) (resEndBlock types.ResponseEndBlock) {
	return types.ResponseEndBlock{Diffs: app.changes}
}

//-----------------------------------------
// persist the last block info

var lastBlockKey = []byte("lastblock")

type LastBlockInfo struct {
	Height  uint64
	AppHash []byte
}

// Get the last block from the db
func LoadLastBlock(db dbm.DB) (lastBlock LastBlockInfo) {
	buf := db.Get(lastBlockKey)
	if len(buf) != 0 {
		r, n, err := bytes.NewReader(buf), new(int), new(error)
		wire.ReadBinaryPtr(&lastBlock, r, 0, n, err)
		if *err != nil {
			cmn.PanicCrisis(errors.Wrap(*err, "cannot load last block (data has been corrupted or its spec has changed)"))
		}
		// TODO: ensure that buf is completely read.
	}

	return lastBlock
}

func SaveLastBlock(db dbm.DB, lastBlock LastBlockInfo) {
	buf, n, err := new(bytes.Buffer), new(int), new(error)
	wire.WriteBinary(lastBlock, buf, n, err)
	if *err != nil {
		// TODO
		cmn.PanicCrisis(errors.Wrap(*err, "cannot save last block"))
	}
	db.Set(lastBlockKey, buf.Bytes())
}

//---------------------------------------------
// update validators

func (app *PersistentABCIDemoApplication) Validators() (validators []*types.Validator) {
	app.app.state.Iterate(func(key, value []byte) bool {
		if isValidatorTx(key) {
			validator := new(types.Validator)
			err := types.ReadMessage(bytes.NewBuffer(value), validator)
			if err != nil {
				panic(err)
			}
			validators = append(validators, validator)
		}
		return false
	})
	return
}

func MakeValSetChangeTx(pubkey []byte, power uint64) []byte {
	return []byte(cmn.Fmt("val:%X/%d", pubkey, power))
}

func isValidatorTx(tx []byte) bool {
	if strings.HasPrefix(string(tx), ValidatorSetChangePrefix) {
		return true
	}
	return false
}

// format is "val:pubkey1/power1,addr2/power2,addr3/power3"tx
func (app *PersistentABCIDemoApplication) execValidatorTx(tx []byte) types.Result {
	tx = tx[len(ValidatorSetChangePrefix):]
	pubKeyAndPower := strings.Split(string(tx), "/")
	if len(pubKeyAndPower) != 2 {
		return types.ErrEncodingError.SetLog(cmn.Fmt("Expected 'pubkey/power'. Got %v", pubKeyAndPower))
	}
	pubkeyS, powerS := pubKeyAndPower[0], pubKeyAndPower[1]
	pubkey, err := hex.DecodeString(pubkeyS)
	if err != nil {
		return types.ErrEncodingError.SetLog(cmn.Fmt("Pubkey (%s) is invalid hex", pubkeyS))
	}
	power, err := strconv.Atoi(powerS)
	if err != nil {
		return types.ErrEncodingError.SetLog(cmn.Fmt("Power (%s) is not an int", powerS))
	}

	// update
	return app.updateValidator(&types.Validator{pubkey, uint64(power)})
}

// add, update, or remove a validator
func (app *PersistentABCIDemoApplication) updateValidator(v *types.Validator) types.Result {
	key := []byte("val:" + string(v.PubKey))
	if v.Power == 0 {
		// remove validator
		if !app.app.state.Has(key) {
			return types.ErrUnauthorized.SetLog(cmn.Fmt("Cannot remove non-existent validator %X", key))
		}
		app.app.state.Remove(key)
	} else {
		// add or update validator
		value := bytes.NewBuffer(make([]byte, 0))
		if err := types.WriteMessage(v, value); err != nil {
			return types.ErrInternalError.SetLog(cmn.Fmt("Error encoding validator: %v", err))
		}
		app.app.state.Set(key, value.Bytes())
	}

	// we only update the changes array if we succesfully updated the tree
	app.changes = append(app.changes, v)

	return types.OK
}


package abci-demo

import (
	"github.com/tendermint/abci/types"
	"github.com/tendermint/merkleeyes/iavl"
	cmn "github.com/tendermint/tmlibs/common"
	"github.com/tendermint/tmlibs/merkle"
)

type ABCIDemoApplication struct {
	types.BaseApplication

	state merkle.Tree
}

func NewABCIDemoApplication() *ABCIDemoApplication {
	state := iavl.NewIAVLTree(0, nil)
	return &ABCIDemoApplication{state: state}
}

func (app *ABCIDemoApplication) Info() (resInfo types.ResponseInfo) {
	return types.ResponseInfo{Data: cmn.Fmt("{\"size\":%v}", app.state.Size())}
}

// tx is either "key=value" or just arbitrary bytes
func (app *ABCIDemoApplication) DeliverTx(tx []byte) types.Result {
	result, err := app.Exec(tx)
	if err != nil {
		println("app.DeliverTx", err.Error())
		return types.NewResult(types.CodeType_InternalError, nil, err.Error())
	}
	return types.NewResultOK(result, "")
}

func (app *ABCIDemoApplication) CheckTx(tx []byte) types.Result {
	err := app.Check(tx)
	if err != nil {
		println("app.CheckTx:", err.Error())
		return types.NewResult(types.CodeType_InternalError, nil, err.Error())
	}
	return types.OK
}

func (app *ABCIDemoApplication) Commit() types.Result {
	hash := app.state.Hash()
	return types.NewResultOK(hash, "")
}

func (app *ABCIDemoApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {
	if reqQuery.Prove {
		value, proof, exists := app.state.Proof(reqQuery.Data)
		resQuery.Index = -1 // TODO make Proof return index
		resQuery.Key = reqQuery.Data
		resQuery.Value = value
		resQuery.Proof = proof
		if exists {
			resQuery.Log = "exists"
		} else {
			resQuery.Log = "does not exist"
		}
		return
	} else {
		index, value, exists := app.state.Get(reqQuery.Data)
		resQuery.Index = int64(index)
		resQuery.Value = value
		if exists {
			resQuery.Log = "exists"
		} else {
			resQuery.Log = "does not exist"
		}
		return
	}
}

func (app *ABCIDemoApplication) Check(tx []byte) error {
	var req Request
	err := UnmarshalMessage(tx, &req)
	if err != nil {
		return err
	}

	err = req.CheckSign()
	if err != nil {
		return err
	}

	err = app.checkInstructionId(req.GetInstructionId())
	if err != nil {
		return err
	}

	err = app.doCheck(&req)
	if err != nil {
		return err
	}
	return nil
}

func (app *ABCIDemoApplication) Exec(tx []byte) ([]byte, error) {
	var req Request
	var resp *Response
	err := UnmarshalMessage(tx, &req)
	if err != nil {
		return nil, err
	}

	resp, err = app.doRequest(&req)
	if err != nil {
		return nil, err
	} else {
		return MarshalMessage(resp)
	}
}

func (app *ABCIDemoApplication) doCheck(req *Request) error {
	var err error

	switch req.GetActionId() {
	case MessageType_MsgInitPlatform:
		err = app.checkInitPlatform(req)
	case MessageType_MsgUserInfoOnChain:
		err = app.checkUserInfoOnChain(req)
	default:
		err = ErrWrongMessageType
	}
	return err
}

func (app *ABCIDemoApplication) doRequest(req *Request) (*Response, error) {
	var resp *Response
	var err error
	switch req.GetActionId() {
	case MessageType_MsgInitPlatform:
		resp, err = app.initPlatform(req)
	case MessageType_MsgUserInfoOnChain:
		resp, err = app.userInfoOnChain(req)
	default:
		err = ErrWrongMessageType
	}

	receipt := &Receipt{}
	if err != nil {
		receipt.Err = []byte(err.Error())
	}
	receipt.IsOk = (err == nil)
	app.saveReceipt(req.GetInstructionId(), receipt)
	return resp, err
}

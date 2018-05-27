package abci-demo_test

import (
	"testing"

	"time"
	"math/rand"
	"encoding/hex"
	"encoding/base64"

	abci "github.com/tendermint/abci/types"
	"github.com/tendermint/abci/example/abci-demo"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	rpc "github.com/tendermint/tendermint/rpc/lib/client"

	"github.com/stretchr/testify/require"

)

var (
	sdkKey = ""
)

func genInitPlatform(t *testing.T) []byte {
	sdkPriv := GetPrivateKeyByPwd("sdkuser", "123456")
	sdkKey = hex.EncodeToString(sdkPriv)

	req := &abci-demo.RequestInitPlatform{}
	req.UserName = "admin"
	req.UserPublicKey= GetPrivateKeyByPwd("admin", "123456")[32:]


	request := &abci-demo.Request{}
	request.Value = &abci-demo.Request_InitPlatform{req}

	request.InstructionId = time.Now().Unix()*1000 + int64(rand.Intn(1000))
	request.Pubkey = sdkPriv[32:]
	request.ActionId = abci-demo.MessageType_MsgInitPlatform


	data, err := abci-demo.MarshalMessage(request)
	if err != nil {
		t.Fatal(err)
	}
	request.Sign = abci-demo.Signdata(sdkPriv, data)

	data, err = abci-demo.MarshalMessage(request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(data))
	return data
}

func TestInitPlatform(t *testing.T) {
	data := genInitPlatform(t)
	testBroadcastTxCommit(t, data, clientJSON)

	data = genInitPlatform(t)
	testBroadcastTxCommitErr(t, data, clientURI)
}

func genUserInfoOnChain(t *testing.T, userName string,password string) []byte {
	//priv := GetPrivateKeyByPwd(userName,password)
	operPriv,err := hex.DecodeString(sdkKey)
	if err != nil {
		t.Fatal(err)
	}

	req := &abci-demo.RequestUserInfoOnChain{}
	req.Username = userName
	req.Age = 34
	req.Educate = "本科"
	req.Workstation = "guangdong"


	request := &abci-demo.Request{}
	request.Value = &abci-demo.Request_UserInfoOnChain{req}

	request.InstructionId = time.Now().Unix()*1000 + int64(rand.Intn(1000))
	request.Pubkey = operPriv[32:]
	request.ActionId = abci-demo.MessageType_MsgUserInfoOnChain


	data, err := abci-demo.MarshalMessage(request)
	if err != nil {
		t.Fatal(err)
	}
	request.Sign = abci-demo.Signdata(operPriv, data)

	data, err = abci-demo.MarshalMessage(request)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(data))
	return data
}

func TestUserInfoOnChain(t *testing.T) {
	data := genUserInfoOnChain(t,"user1","123456")
	testBroadcastTxCommit(t, data, clientJSON)

	//data = genUserInfoOnChain(t,"user2","kkjfhs")
	//testBroadcastTxCommit(t, data, clientJSON)

	//data = genUserInfoOnChain(t,"user1","123456")
	//testBroadcastTxCommitErr(t, data, clientURI)

}


func testBroadcastTxCommitErr(t *testing.T, tx []byte, client rpc.HTTPClient) {
	require := require.New(t)

	result := new(ctypes.ResultBroadcastTxCommit)
	_, err := client.Call("broadcast_tx_commit", map[string]interface{}{"tx": tx}, result)
	require.Nil(err)

	checkTx := result.CheckTx
	require.Equal(abci.CodeType_InternalError, checkTx.Code)
}

func testBroadcastTxCommit(t *testing.T, tx []byte, client rpc.HTTPClient) {
	require := require.New(t)

	result := new(ctypes.ResultBroadcastTxCommit)
	_, err := client.Call("broadcast_tx_commit", map[string]interface{}{"tx": tx}, result)
	require.Nil(err)

	checkTx := result.CheckTx
	require.Equal(abci.CodeType_OK, checkTx.Code)
	deliverTx := result.DeliverTx
	require.Equal(abci.CodeType_OK, deliverTx.Code)
	mem := node.MempoolReactor().Mempool
	require.Equal(0, mem.Size())
}

func GetPrivateKeyByPwd(userName, pwd string) []byte {
	prv := abci-demo.GetPrivateKey(pwd)
	prv = abci-demo.GetPrivateKey(userName + "_" + hex.EncodeToString(prv))
	pub := abci-demo.GetPublicKey(prv)
	var p [64]byte
	copy(p[:32], prv[:])
	copy(p[32:], pub[:])
	return p[:]
}





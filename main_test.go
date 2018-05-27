package abci-demo_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/tendermint/abci/example/abci-demo"
	nm "github.com/tendermint/tendermint/node"
	client "github.com/tendermint/tendermint/rpc/lib/client"
)

var node *nm.Node

var clientURI *client.URIClient
var clientJSON *client.JSONRPCClient

func TestMain(m *testing.M) {
	// start a tendermint node (and merkleeyes) in the background to test against
	dir, err := ioutil.TempDir("/tmp", "abci-demo-test") // TODO
	if err != nil {
		panic(err)
	}
	app := abci-demo.NewPersistentABCIDemoApplication(dir)

	node = StartTendermint(app)
	clientURI = GetURIClient()
	clientJSON = GetJSONClient()
	code := m.Run()

	// and shut down proper at the end
	node.Stop()
	node.Wait()
	os.Exit(code)
}

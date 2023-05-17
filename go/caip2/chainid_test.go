package caip2_test

import (
	"fmt"
	"github.com/aviate-labs/agent-go"
	"github.com/fxamacker/cbor/v2"
	"github.com/icvc/icp-namespace/go/caip2"
	"net/url"
	"os/exec"
	"testing"
)

var icp0, _ = url.Parse("https://icp0.io/")

func TestChainId_curl(t *testing.T) {
	path, err := exec.LookPath("curl")
	if err != nil {
		t.Skip("curl not found")
	}
	cborResponse, err := exec.Command(path, "https://icp0.io/api/v2/status").Output()
	if err != nil {
		t.Fatal(err)
	}
	var raw map[string]any
	if err := cbor.Unmarshal(cborResponse, &raw); err != nil {
		t.Fatal(err)
	}

	expected := "icp:737ba355e855bd4b61279056603e0550"
	if chainId := caip2.ChainId(raw["root_key"].([]byte)); chainId != expected {
		t.Error(chainId)
	}
}

func Example() {
	c := agent.NewClient(agent.ClientConfig{Host: icp0})
	status, _ := c.Status()
	fmt.Println(caip2.ChainId(status.RootKey))
	// Output:
	// icp:737ba355e855bd4b61279056603e0550
}

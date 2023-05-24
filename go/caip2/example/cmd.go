package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/aviate-labs/agent-go"
	"net/url"
)

var icp0, _ = url.Parse("https://icp0.io/")

func main() {
	c := agent.NewClient(agent.ClientConfig{Host: icp0})
	status, _ := c.Status()
	hash := sha256.Sum256(status.RootKey)
	hashHexString := hex.EncodeToString(hash[:])
	fmt.Printf("icp:%s", hashHexString[:32])
}

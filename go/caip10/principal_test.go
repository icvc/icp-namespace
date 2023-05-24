package caip10_test

import (
	"fmt"
	"github.com/aviate-labs/agent-go/principal"
	"github.com/icvc/icp-namespace/go/caip10"
)

func Example_textual() {
	p := caip10.TextualEncode([]byte{0x00})
	rp, _ := caip10.TextualDecode(p)
	fmt.Println(p)
	fmt.Printf("%02x\n", rp)
	// Output:
	// 2ibo7-dia
	// 00
}

func Example_binary() {
	fmt.Println(principal.AnonymousID.String())
	// Output:
	// 2ibo7-dia
}

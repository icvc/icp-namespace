package caip10

import "github.com/aviate-labs/agent-go/principal"

func TextualEncode(raw []byte) string {
	return principal.Principal{Raw: raw}.String()
}

func TextualDecode(str string) ([]byte, error) {
	p, err := principal.Decode(str)
	return p.Raw, err
}

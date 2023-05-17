package caip2

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Namespace() string {
	return "icp"
}

func Reference(rootKey []byte) string {
	hash := sha256.Sum256(rootKey)
	hashHexString := hex.EncodeToString(hash[:])
	return hashHexString[:32]
}

func ChainId(rootKey []byte) string {
	return fmt.Sprintf("%s:%s", Namespace(), Reference(rootKey))
}

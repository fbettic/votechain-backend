package verification

import "github.com/ethereum/go-ethereum/crypto"

func GetHash() string {

	return string(crypto.Keccak256())
}

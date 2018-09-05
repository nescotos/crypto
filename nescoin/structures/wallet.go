package structures

import "crypto/ecdsa"

//Wallet : Digital Wallet for nescoin
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

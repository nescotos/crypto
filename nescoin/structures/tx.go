package structures

//TX : Transaction representation for nescoin
type TX struct {
	TxID      []byte
	Signature []byte
	PublicKey []byte
	Value     int
	Payload   []byte
}

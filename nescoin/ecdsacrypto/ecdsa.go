package ecdsacrypto

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
)

//KeyGen : Generates a public/pair key using Elliptic Curves Digital Signature Algorithm
func KeyGen() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	checkError(err, "Something went wrong generating keys")
	publicKey := &privateKey.PublicKey
	savePEMFileForPrivateKey("private.pem", privateKey)
	savePEMFileForPublicKey("public.pem", publicKey)
}

//Sign : Generate a signature using an specific private key
func Sign(message string) {
	hash := sha256.Sum256([]byte(message))
	privateKey := readKey("private.pem")
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hash[:])
	signature := append(r.Bytes(), s.Bytes()...)
	pubKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)
	checkError(err, "Error signing message")
	log.Println("R: ", r)
	log.Println("S: ", s)
	fmt.Printf("{Signature: [% x]}", string(signature))
	fmt.Printf("{PubKey: [% x]}", string(pubKey))
	Verify("Perro", signature, pubKey)
}

//Verify : Verify the Digital Signature using the public key
func Verify(message string, signature []byte, publicKey []byte) {
	hash := sha256.Sum256([]byte(message))
	r := big.Int{}
	s := big.Int{}
	sigLen := len(signature)
	r.SetBytes(signature[:(sigLen / 2)])
	s.SetBytes(signature[(sigLen / 2):])
	x := big.Int{}
	y := big.Int{}
	keyLen := len(publicKey)
	x.SetBytes(publicKey[:(keyLen / 2)])
	y.SetBytes(publicKey[(keyLen / 2):])
	key := ecdsa.PublicKey{elliptic.P256(), &x, &y}
	//key := readPublicKey(publicKey)
	log.Println(ecdsa.Verify(&key, hash[:], &r, &s))
}

func readKey(fileName string) *ecdsa.PrivateKey {
	privateKeyFile, err := os.Open(fileName)
	checkError(err, "Error Opening "+fileName)
	defer privateKeyFile.Close()
	pemFileInfo, _ := privateKeyFile.Stat()
	size := pemFileInfo.Size()
	pemBytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pemBytes)
	checkError(err, "Error decoding"+fileName)
	data, _ := pem.Decode([]byte(pemBytes))
	privateKey, err := x509.ParseECPrivateKey(data.Bytes)
	checkError(err, "Error Parsing "+fileName)
	return privateKey
}

func savePEMFileForPrivateKey(fileName string, key *ecdsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err, "There was an error creating file: "+fileName)
	defer outFile.Close()
	privateKeyEncoded, _ := x509.MarshalECPrivateKey(key)
	pemKey := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyEncoded,
	}
	err = pem.Encode(outFile, pemKey)
	checkError(err, "There was an error encoding PEM into "+fileName)
}

func savePEMFileForPublicKey(fileName string, key *ecdsa.PublicKey) {
	outFile, err := os.Create(fileName)
	checkError(err, "There was an error creating file: "+fileName)
	defer outFile.Close()
	publicKeyEncoded, _ := x509.MarshalPKIXPublicKey(key)
	pemKey := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyEncoded,
	}
	err = pem.Encode(outFile, pemKey)
	checkError(err, "There was an error encoding PEM into "+fileName)
}

func readPublicKey(key []byte) *ecdsa.PublicKey {
	publicKeyBlock, _ := pem.Decode(key)
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	checkError(err, "Something went wrong parsing the public key")
	return publicKey.(*ecdsa.PublicKey)
}

func checkError(err error, message string) {
	if err != nil {
		log.Println(message, err.Error())
		os.Exit(1)
	}
}

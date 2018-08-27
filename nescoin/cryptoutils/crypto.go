package cryptoutils

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

//GeneratePair : Generates a public/private key pairs!
func GeneratePair() {
	bitSize := 4096
	log.Println("Generating Pair")
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	printError(err, "Error Generating Pair: ")
	publicKey := privateKey.PublicKey
	saveKey("private.key", privateKey)
	saveKey("public.key", publicKey)
	savePEM("private.pem", x509.MarshalPKCS1PrivateKey(privateKey), "PRIVATE KEY", privateKey)
	publicByte, err := asn1.Marshal(publicKey)
	printError(err, "Unable to encode Public Key: ")
	savePEM("public.pem", publicByte, "PUBLIC KEY", publicKey)
}

//SignMessage : Sign a message using the Private Key
func SignMessage(message string) {
	key := readKey("private.pem")
	signature := sign(message, key)
	fmt.Printf("Signature: %x", signature)
}

//ValidateSignature : Validates a signature using an specific Public Key
func ValidateSignature(key []byte, signature string, message string) {
	sig, _ := hex.DecodeString(message)
	payload := sha256.Sum256([]byte(message))
	publicKey := readPublicKey(key)
	//log.Println(crypto.SHA256.Size(), " vs ", len(message))
	log.Println(rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, payload[:], sig))
}

func readKey(fileName string) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(fileName)
	printError(err, "Error Opening "+fileName)
	defer privateKeyFile.Close()
	pemFileInfo, _ := privateKeyFile.Stat()
	size := pemFileInfo.Size()
	pemBytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pemBytes)
	printError(err, "Error decoding"+fileName)
	data, _ := pem.Decode([]byte(pemBytes))
	privateKey, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	printError(err, "Error Parsing "+fileName)
	return privateKey
}

func readPublicKey(key []byte) *rsa.PublicKey {
	publicKeyBlock, _ := pem.Decode(key)
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	printError(err, "Something went wrong parsing the public key")
	return publicKey
}

func saveKey(fileName string, key interface{}) {
	log.Println("Saving ", fileName)
	outFile, err := os.Create(fileName)
	printError(err, "Unable to create file "+fileName)
	defer outFile.Close()
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	printError(err, "Error encoding file "+fileName)
	log.Println(fileName, "saved")
}

func savePEM(fileName string, pemBytes []byte, pemType string, key interface{}) {
	log.Println("Saving Private Key PEM File ", fileName)
	outFile, err := os.Create(fileName)
	printError(err, "Unable to create file"+fileName)
	defer outFile.Close()

	pemKey := &pem.Block{
		Type:  pemType,
		Bytes: pemBytes,
	}

	err = pem.Encode(outFile, pemKey)
	printError(err, "Error encoding file"+fileName)
	log.Println(fileName, "saved")
}

func sign(message string, key *rsa.PrivateKey) []byte {
	hash := sha256.Sum256([]byte(message))
	fmt.Printf("Message: %x\n", hash)
	signature, err := rsa.SignPKCS1v15(rand.Reader, key, crypto.SHA256, hash[:])
	printError(err, "Something went wrong while sign the message")
	fmt.Printf("Signature: %x\n", signature)
	return signature
}

func printError(err error, message string) {
	if err != nil {
		log.Println(message, err.Error())
		os.Exit(1)
	}
}

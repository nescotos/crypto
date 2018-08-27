package cryptoutils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/pem"
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

func printError(err error, message string) {
	if err != nil {
		log.Println(message, err.Error())
		os.Exit(1)
	}
}

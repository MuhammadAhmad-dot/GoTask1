package main

import (
	"crypto/dsa"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"hash"
	"io"
	"math/big"
	"os"
)

func main() {

	Dsaparameters := new(dsa.Parameters) // Creating new object of dsa parameters

	if err := dsa.GenerateParameters(Dsaparameters, rand.Reader, dsa.L1024N160); err != nil { //Generating new parameters of dsa of 1024bit size
		fmt.Println(err)
		os.Exit(1)
	}

	privatekey := new(dsa.PrivateKey) //creating new RSA 1024-bit private key object

	privatekey.PublicKey.Parameters = *Dsaparameters //creating new RSA 1024-bit public key object

	dsa.GenerateKey(privatekey, rand.Reader) // this generates a public & private key pair

	var publickey dsa.PublicKey //variable to store public key

	publickey = privatekey.PublicKey //Extracting public key from private key

	fmt.Println("Private Key :")
	fmt.Printf("%num2\n\n", privatekey)

	fmt.Println("Public Key :")
	fmt.Printf("%num2 \n\n", publickey)

	// Signing
	var hashobject hash.Hash
	hashobject = md5.New() //creatig new md5 hash object

	num1 := big.NewInt(0) //variable to store big integers
	num2 := big.NewInt(0)

	myname := "Muhammad Ahmad"

	io.WriteString(hashobject, myname)
	signaturehash := hashobject.Sum(nil)

	num1, num2, err := dsa.Sign(rand.Reader, privatekey, signaturehash) //Generating signature using the private key
	if err != nil {
		fmt.Println(err)
	}

	signature := num1.Bytes()
	signature = append(signature, num2.Bytes()...)

	fmt.Printf("Signature : %x\n", signature)

	verification := dsa.Verify(&publickey, signaturehash, num1, num2) // Verifying the signature by using public key
	fmt.Println("signature valid:", verification)

	myname = "Mirza Ahmad" // changing in data to dectect the change in signature

	io.WriteString(hashobject, myname)
	signaturehash = hashobject.Sum(nil)

	verification = dsa.Verify(&publickey, signaturehash, num1, num2) // Verifying the signature by using public key
	fmt.Println("signature valid:", verification)
}

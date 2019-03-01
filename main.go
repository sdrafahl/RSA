package main

import "fmt"

func main() {
	privateKey, publicKey := gen()
	fmt.Println("Private Key")
	fmt.Println(privateKey)
	fmt.Println("Public Key")
	fmt.Println(publicKey)
	//var encryptedMessage string = encrypt("secret", publicKey)
	//fmt.Println("Encrypted")
	//fmt.Println(encryptedMessage)
	//var decryptedMessage string = decrypt(encryptedMessage, privateKey)
	//fmt.Println("Decrypted")
	//fmt.Println(decryptedMessage)
}

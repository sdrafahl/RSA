package main

import (
	"fmt"
	"math"
)

func encrypt(message string, publicKey PublicKey) string {
	var newString string = ""
	for _, char := range message {
		fmt.Println(int(char))
		newString += string(int(math.Mod(math.Pow(float64(int(char)), float64(publicKey.e)), float64(publicKey.n))))
	}
	return newString
}

func decrypt(message string, privateKey PrivatKey) string {
	var newString string = ""
	for _, char := range message {
		newString += string(int(math.Mod(math.Pow(float64(int(char)), float64(privateKey.d)), float64(privateKey.n))))
		fmt.Println(int(math.Mod(math.Pow(float64(int(char)), float64(privateKey.d)), float64(privateKey.n))))
	}
	return newString
}

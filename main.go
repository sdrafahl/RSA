package main

import "fmt"

func main() {
	n, d := gen()
	fmt.Println("Private Key")
	fmt.Println(n)
	fmt.Println("Public Key")
	fmt.Println(d)
}

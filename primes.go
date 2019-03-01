package main

import (
	"math"
)

func twoPrimes(index int) (int, int) {
	prime1, prime2 := 0, 0
	first := true
	for i := 3; i < index; i++ {
		if isPrime(i) {
			if prime1 == 0 {
				prime1 = i
			} else {
				if prime2 == 0 {
					prime2 = i
				} else {
					if first {
						prime1 = i
					} else {
						prime2 = i
					}
					first = !first
				}
			}
		}
	}
	return prime1, prime2
}

func isPrime(numberInQuestion int) bool {
	for i := 2; i < numberInQuestion; i++ {
		remainder := math.Mod(float64(numberInQuestion), float64(i))
		if remainder == 0 {
			return false
		}
	}
	return true
}

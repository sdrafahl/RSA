package rsa

import "math"

func twoPrimes(index int) (int, int) {
	prime1, prime2 := 0, 0
	for i := 2; i < index; i++ {
		if isPrime(index) {
			if prime1 == 0 {
				prime1 = index
			} else {
				if prime2 == 0 {
					prime2 = index
				}
			}
		}
	}
	return prime1, prime2
}

func isPrime(numberInQuestion int) bool {
	half := numberInQuestion / 2
	for i := 2; i < half; i++ {
		remainder := math.Mod(float64(numberInQuestion), float64(i))
		if remainder > 1 {
			return false
		}
	}
	return true
}

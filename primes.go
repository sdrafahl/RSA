package rsa

import "math"

func twoPrimes(index int) (int, int) {

}

func isPrime(numberInQuestion int) bool {
	half := numberInQuestion
	for i := 2; i < half; i++ {
		remainder := math.Mod(float64(numberInQuestion), float64(i))
		if remainder > 1 {
			return false
		}
	}
	return true
}

package rsa

import "math"

func gen() {
	p, q := twoPrimes(100)
	n := p * q
	totient := (p - 1) * (q - 1)
	e := find_e(totient)
}

func find_e(totient int) int {
	var e int
	for i := 2; i < totient; i++ {
		if int(math.Mod(float64(totient), float64(i))) != 0 {
			e = i
		}
	}
	return e
}

func euclid(totient int, e int) int {
	var polys []Polynomial
	answer := totient
	factor1 := e
	factor2 := 0
	remainder := int(math.Mod(float64(totient), float64(e)))
	for true {
		factor2 = (answer - remainder) / factor1
		polys = append(polys, Polynomial{answer, factor1, factor2, remainder})
		if remainder == 1 {
			break
		}
		answer = factor1
		factor1 = remainder
		remainder = int(math.Mod(float64(answer), float64(factor1)))
	}
	for i := 0; i < len(polys); i++ {

	}
}

type Polynomial struct {
	answer    int
	factor1   int
	factor2   int
	remainder int
}

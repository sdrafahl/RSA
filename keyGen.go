package main

import (
	"math"
)

func gen() (PrivatKey, PublicKey) {
	p, q := twoPrimes(10)
	n := p * q
	totient := (p - 1) * (q - 1)
	e := find_e(totient)
	d := computeD(totient, e)
	var privateKey PrivatKey = PrivatKey{p, q, totient, d}
	var publicKey PublicKey = PublicKey{n, e}
	return privateKey, publicKey
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

func getFactor(factor int, polymer Polynomial, polymers []Polynomial, totient int) int {
	var factorValue1 int = 0
	if polymer.answer == factor {
		factorValue1 = 1
	} else {
		var polymer *Polynomial = searchPolynomials(factor, polymers)
		if polymer != nil {
			factorValue1 = getFactor(factor, *polymer, polymers, totient)
		}
	}
	var factorValue2 int = 0
	if polymer.factor1 == factor {
		factorValue2 = polymer.factor2 * -1
	} else {
		var polymer *Polynomial = searchPolynomials(factor, polymers)
		if polymer != nil {
			factorValue2 = getFactor(factor, *polymer, polymers, totient) * polymer.factor2 * -1
		}
	}
	d := factorValue1 + factorValue2
	if d < 0 {
		return int(totient - int(math.Abs(float64(d))))
	} else {
		return d
	}
}

func searchPolynomials(remainder int, polynomials []Polynomial) *Polynomial {
	for i, s := range polynomials {
		if s.remainder == remainder {
			i++
			return &s
		}
	}
	return nil
}

func computeD(totient int, e int) int {
	var polys []Polynomial
	answer := totient
	factor1 := 1
	factor1 = e
	factor2 := 0
	remainder := int(math.Mod(float64(totient), float64(e)))
	incrament := 0
	for true {
		incrament++
		factor2 = (answer - remainder) / factor1
		polys = append(polys, Polynomial{answer, factor1, factor2, remainder})
		if remainder == 1 {
			break
		}
		answer = factor1
		factor1 = remainder
		remainder = int(math.Mod(float64(answer), float64(factor1)))
	}
	var polyNomial Polynomial = (polys[incrament-1 : incrament])[0]
	var d int = getFactor(e, polyNomial, polys, totient)
	return d
}

type Polynomial struct {
	answer    int
	factor1   int
	factor2   int
	remainder int
}

type PrivatKey struct {
	p       int
	q       int
	totient int
	d       int
}

type PublicKey struct {
	n int
	e int
}

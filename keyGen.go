package main

import (
	"fmt"
	"math"
)

func gen() (PrivatKey, PublicKey) {
	//p, q := twoPrimes(10)
	p := 61
	q := 53
	n := p * q
	totient := (p - 1) * (q - 1)
	//e := find_e(totient)
	e := 17
	d := computeD(totient, e)
	var privateKey PrivatKey = PrivatKey{p, q, totient, d, n}
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
		var searchedPolymer *Polynomial = searchPolynomials(polymer.answer, polymers)
		if searchedPolymer != nil {
			factorValue1 = getFactor(factor, *searchedPolymer, polymers, totient)
		}
	}
	var factorValue2 int = 0
	if polymer.factor1 == factor {
		factorValue2 = polymer.factor2 * -1
	} else {
		var searchedPolymer *Polynomial = searchPolynomials(polymer.factor1, polymers)
		if searchedPolymer != nil {
			factorValue2 = getFactor(factor, *searchedPolymer, polymers, totient) * polymer.factor2 * -1
		}
	}
	//fmt.Println("total")
	//fmt.Println(polymer)
	//fmt.Println(factorValue1)
	//fmt.Println(factorValue2)
	//fmt.Println(factorValue1 + factorValue2)
	return factorValue1 + factorValue2
}

func searchPolynomials(remainder int, polynomials []Polynomial) *Polynomial {
	for _, s := range polynomials {
		if s.remainder == remainder {
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
	fmt.Println("d:")
	fmt.Println(d)

	if d < 0 {
		return totient + d
	}
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
	n       int
}

type PublicKey struct {
	n int
	e int
}

package rsa

import "math"

func gen() (int, int) {
	p, q := twoPrimes(100)
	n := p * q
	totient := (p - 1) * (q - 1)
	e := find_e(totient)
	d := computeD(totient, e)
	return (n, d)
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

func getFactor(factor int, polymer Polynomial, polymers []Polynomial) int {
	var factorValue1 int = 0
	if polymer.answer == factor {
		factorValue1 = 1
	} else {
		var polymer Polynomial = searchPolynomials(factor, polymers)
		if polymer != nil {
			factorValue1 = getFactor(factor, polymer, polymers)
		}
	}
	var factorValue2 int = 0
	if polymer.factor1 == factor {
		factorValue2 = polymer.factor2 * -1
	} else {
		var polymer Polynomial = searchPolynomials(factor, polymers)
		if polymer != nil {
			factorValue2 = getFactor(factor, polymer, polymers) * polymer.factor2 * -1
		}
	}
	return factorValue1 + factorValue2
}

func searchPolynomials(remainder int, polynomials []Polynomial) Polynomial {
	for i, s := range polynomials {
		if s.remainder == remainder {
			return s
		}
	}
	return nil
}

func computeD(totient int, e int) int {
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
	var polyNomial Polynomial = (polys[i:i+1])[0]
	var d int = getFactor(polyNomial, polys)
	return d
}

type Polynomial struct {
	answer    int
	factor1   int
	factor2   int
	remainder int
}

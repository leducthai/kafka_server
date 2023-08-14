package primenum

import "math"

func FindNextPrime(p int) int {
	rt := p

	for {
		if isPrime(rt) {
			break
		}
		rt += 1
	}

	return rt
}

func isPrime(p int) bool {
	if p <= 1 {
		return false
	}

	if p%2 == 0 {
		return false
	}
	if p%3 == 0 {
		return false
	}

	for i := 5; i <= int(math.Sqrt(float64(p))); i += 6 {
		if p%i == 0 || p%(i+2) == 0 {
			return false
		}
	}

	return true
}

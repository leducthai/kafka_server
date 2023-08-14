package primenum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsPrime(t *testing.T) {
	assertion := assert.New(t)

	// good case
	{ // is a prime  number
		rp := isPrime(7)
		assertion.True(rp)
	}
	{ // is not prime
		rp := isPrime(9)
		assertion.False(rp)
	}

}

func Test_FindNextPrime(t *testing.T) {
	assertion := assert.New(t)

	np := FindNextPrime(9)
	assertion.Equal(11, np)

	np = FindNextPrime(758)
	assertion.Equal(761, np)

	np = FindNextPrime(757)
	assertion.Equal(757, np)
}

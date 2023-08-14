package terminator

import (
	"math/rand"
	"producer/algorithms/hash"
)

type Terminator[T comparable] interface {
	AddEle(T) bool
}

type HashFunc interface {
	Hash(int, int) int
	GetHashInfor() []int
} 

type PairWise struct {
	a int
	b int
	set bool
}

type FieldOptions func(*PairWise)

func WithDefaultHash(a int, b int) FieldOptions {
	return func(p *PairWise) {
		p.a = a
		p.b = b
		p.set = true
	}
}

func (p PairWise) Hash(x int, m int) int {
	return (p.a*x + p.b) % m
}

func NewPairWise(p int, opts []FieldOptions) HashFunc {
	rt := PairWise{}

	for _, opt := range opts {
		opt(&rt)
	}

	if !rt.set {
		rt.a = rand.Intn(p)
		rt.b = rand.Intn(p)
	}

	return rt
}

type HashTable[T comparable] interface {
	ChangeFirstHash()
	ChangeSecondHash()
	AddSpace()
	DownSpace()
	AddEle(T)
	GetP()
}

type DoubleHashTable[T comparable] struct {
	FirstHashFunc  HashFunc
	SecondHashFunc HashFunc
	Terminate      Terminator[T]
	TableOne       map[int]string
	TableTwo       map[int]string
	NumEle         int
	P              int
}

func NewDoubleHashTable[T comparable](init int, t Terminator[T], withBuiltHashs bool) HashTable[T] {
	return DoubleHashTable[T]{
		FirstHashFunc: aaa,
		SecondHashFunc: bbb,
		Terminate: t,
		TableOne:  map[int]string{},
		TableTwo:  map[int]string{},
		NumEle:    0,
		P:         init,
	}
}

func (d DoubleHashTable[T]) ChangeHash() {
	d.changeHashFuncs()
}

func (d DoubleHashTable[T]) AddEle(ele T) {
	d.addEle(ele)
}

func (d DoubleHashTable[T]) DownSpace() {

}

func (d DoubleHashTable[T]) AddSpace() {

}

func (d *DoubleHashTable[T]) changeHashFuncs() {
	
}

func (d *DoubleHashTable[T]) addEle(ele T) {,. M?NN
	eleToNum := hash.HashToInt(ele)

	pos1 := 
}

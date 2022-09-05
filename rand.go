package crand

import (
	crand "crypto/rand"
	"math"
	"math/big"
	"math/rand"
)

func NewRand() *Rand {
	s := NewSource()

	return &Rand{rand.New(s), s}
}

// AddressLength represents Ethereum address length
const AddressLength = 20

func (r *Rand) Address() [AddressLength]byte {
	bytes := make([]byte, AddressLength)
	_, _ = r.Read(bytes)

	var a [AddressLength]byte

	copy(a[:], bytes)

	return a
}

func BigInt(max *big.Int) *big.Int {
	n, _ := crand.Int(crand.Reader, max)
	return n
}

// Source doesn't respect a seed
type Source struct{}

func (s Source) Int63() int64 {
	intN, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	return intN.Int64()
}

func (s Source) Seed(_ int64) {}

func NewSource() rand.Source {
	return Source{}
}

type Source64 struct {
	Source
}

func (s Source64) Uint64() uint64 {
	n, _ := crand.Int(crand.Reader, big.NewInt(0).SetUint64(math.MaxUint64))
	return n.Uint64()
}

func NewSource64() rand.Source64 {
	return Source64{}
}

type Rand struct {
	*rand.Rand
	rand.Source
}

package datt

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrBitsetOfDifferentSize     = errors.New("bitsets must be of the same size")
	ErrBitsetOperationOutOfRance = errors.New("out of range operation")
	ErrBitsetWrongSize           = errors.New("the size of the bitset must be > 0")
)

type Bitset struct {
	bytes   []uint8
	numByte int
	numBits int
}

func NewBitset(numBits int) (*Bitset, error) {
	if numBits < 0 {
		return nil, ErrBitsetWrongSize
	}

	numByte := ((numBits - 1) / 8) + 1

	return &Bitset{
		bytes:   make([]uint8, numByte),
		numByte: numByte,
		numBits: numBits,
	}, nil
}

func (b *Bitset) Set(n int, val bool) error {
	if n < 0 || n >= b.numBits {
		return ErrBitsetOperationOutOfRance
	}

	byteIndex := n / 8
	bitPosition := n % 8

	if val {
		b.bytes[byteIndex] |= (1 << bitPosition)
	} else {
		b.bytes[byteIndex] &= (0b11111111 - (1 << bitPosition))
	}
	return nil
}

func (b *Bitset) Get(n int) bool {
	if n < 0 || n >= b.numBits {
		return false
	}

	byteIndex := n / 8
	bitPosition := n % 8

	return (b.bytes[byteIndex] & (1 << bitPosition)) != 0
}

func (b *Bitset) Any() bool {
	for _, b := range b.bytes {
		if b != 0 {
			return true
		}
	}

	return false
}

func (b *Bitset) None() bool {
	for _, b := range b.bytes {
		if b != 0 {
			return false
		}
	}

	return true
}

func (b *Bitset) All() bool {
	for _, b := range b.bytes {
		if b != 0b11111111 {
			return false
		}
	}

	return true
}

func (b *Bitset) Contain(other *Bitset) bool {
	if b.numBits != other.numBits {
		return false
	}

	for i := 0; i < b.numByte; i++ {
		res := (b.bytes[i] & other.bytes[i])

		if res != other.bytes[i] {
			return false
		}
	}

	return true
}

func (b *Bitset) Equal(other *Bitset) bool {
	if b.numBits != other.numBits {
		return false
	}

	for i := 0; i < b.numByte; i++ {
		if b.bytes[i] != other.bytes[i] {
			return false
		}
	}

	return true
}

func (b *Bitset) And(other *Bitset) error {
	if b.numBits != other.numBits {
		return ErrBitsetOfDifferentSize
	}

	for i := 0; i < b.numByte; i++ {
		b.bytes[i] &= other.bytes[i]
	}

	return nil
}

func (b *Bitset) Or(other *Bitset) error {
	if b.numBits != other.numBits {
		return ErrBitsetOfDifferentSize
	}

	for i := 0; i < b.numByte; i++ {
		b.bytes[i] |= other.bytes[i]
	}

	return nil
}

func (b *Bitset) Xor(other *Bitset) error {
	if b.numBits != other.numBits {
		return ErrBitsetOfDifferentSize
	}

	for i := 0; i < b.numByte; i++ {
		current := b.bytes[i]
		other := other.bytes[i]

		sameBits := current & other
		invert := (0b11111111 - sameBits)
		sum := current | other

		b.bytes[i] = (sum & invert)
	}

	return nil
}

func (b *Bitset) Flip() {
	for i, bits := range b.bytes {
		b.bytes[i] = (0b11111111 - bits)
	}
}

func (b *Bitset) Clone() *Bitset {
	newBiteset := Bitset{
		numByte: b.numByte,
		numBits: b.numBits,
		bytes:   make([]uint8, b.numByte),
	}

	copy(newBiteset.bytes, b.bytes)

	return &newBiteset
}

func (b *Bitset) Size() int {
	return b.numBits
}

func (b *Bitset) Clear() {
	for i := 0; i < b.numByte; i++ {
		b.bytes[i] = 0
	}
}

func (b *Bitset) String() string {
	var sb strings.Builder

	for i := b.numByte - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%08b ", b.bytes[i]))
	}

	return sb.String()
}

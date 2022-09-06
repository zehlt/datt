package datt

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrBitsetOfDifferentSize = errors.New("bitsets must be of the same size")
)

type Bitset struct {
	bytes []uint8
	size  int
}

func NewBitset(byteCapacity int) *Bitset {
	return &Bitset{
		bytes: make([]uint8, byteCapacity),
	}
}

func (b *Bitset) Set(n int, val bool) {
	if n < 0 {
		return
	}

	byteIndex := n / 8
	bitPosition := n % 8

	if byteIndex >= b.size {
		diff := byteIndex + 1 - b.size
		for i := 0; i < diff; i++ {
			b.bytes = append(b.bytes, 0)
			b.size++
		}
	}

	if val {
		b.bytes[byteIndex] |= (1 << bitPosition)
	} else {
		b.bytes[byteIndex] &= (0b11111111 - (1 << bitPosition))
	}
}

func (b *Bitset) Get(n int) bool {
	if n < 0 {
		return false
	}

	if n >= b.size*8 {
		return false
	}

	byteIndex := n / 8
	bitPosition := n % 8

	return (b.bytes[byteIndex] & (1 << bitPosition)) != 0
}

func (b *Bitset) Include(other *Bitset) bool {
	if other.size > b.size {
		return false
	}

	for i := 0; i < other.size; i++ {
		if b.bytes[i] != other.bytes[i] {
			return false
		}
	}

	return true
}

func (b *Bitset) Match(matcher *Bitset) bool {
	for i := 0; i < matcher.size; i++ {
		match_byte := matcher.bytes[i]

		if match_byte == 0 {
			continue
		}

		if i > b.size {
			return false
		}

		if (b.bytes[i] & match_byte) != match_byte {
			return false
		}
	}

	return true
}

func (b *Bitset) Equal(other *Bitset) bool {
	if b.size != other.size {
		return false
	}

	for i := 0; i < b.size; i++ {
		if b.bytes[i] != other.bytes[i] {
			return false
		}
	}

	return true
}

func (b *Bitset) And(other *Bitset) error {
	if b.size != other.size {
		return ErrBitsetOfDifferentSize
	}

	for i := 0; i < b.size; i++ {
		b.bytes[i] &= other.bytes[i]
	}

	return nil
}

func (b *Bitset) Or(other *Bitset) error {
	if b.size != other.size {
		return ErrBitsetOfDifferentSize
	}

	for i := 0; i < b.size; i++ {
		b.bytes[i] |= other.bytes[i]
	}

	return nil
}

func (b *Bitset) Xor(other *Bitset) error {
	if b.size != other.size {
		return ErrBitsetOfDifferentSize
	}

	for i := 0; i < b.size; i++ {
		current := b.bytes[i]
		other := other.bytes[i]

		sameBits := current & other
		invert := (0b11111111 - sameBits)
		sum := current | other

		b.bytes[i] = (sum & invert)
	}

	return nil
}

func (b *Bitset) Lenght() int {
	return b.size
}

func (b *Bitset) Clear() {
	b.bytes = make([]uint8, 0)
	b.size = 0
}

func (b *Bitset) String() string {
	var sb strings.Builder

	for i := b.size - 1; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%08b ", b.bytes[i]))
	}

	return sb.String()
}

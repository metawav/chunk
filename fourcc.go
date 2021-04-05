package chunk

import (
	"encoding/binary"
	"math"
)

// FourCC is a sequence of four bytes (Four-character-code) used for data format identification.
type FourCC [4]byte

// CreateFourCC creates a four byte sequence from provided string value. Overlapping bytes will be cut.
func CreateFourCC(value string) FourCC {
	size := math.Min(4, float64(len(value)))
	var fourCC FourCC
	copy(fourCC[:4], value[:int(size)])

	return fourCC
}

// String returns string value.
func (f *FourCC) String() string {
	return string(f[:])
}

// ToUint32 returns uint32 value.
func (f *FourCC) ToUint32() uint32 {
	return binary.BigEndian.Uint32(f[:])
}

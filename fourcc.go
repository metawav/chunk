package chunk

import "encoding/binary"

// FourCC is a sequence of four bytes (Four-character-code) used for data format identification.
type FourCC [4]byte

// String returns string value.
func (f *FourCC) String() string {
	return string(f[:])
}

// ToUint32 returns uint32 value.
func (f *FourCC) ToUint32() uint32 {
	return binary.BigEndian.Uint32(f[:])
}

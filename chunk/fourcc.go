package chunk

import "encoding/binary"

type FourCC [4]byte

func (f *FourCC) String() string {
	return string(f[:])
}

func (f *FourCC) ToUint32() uint32 {
	return binary.BigEndian.Uint32(f[:])
}

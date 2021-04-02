package chunk

import (
	"encoding/binary"
	"fmt"

	"github.com/pmoule/wav/strconv"
)

// RiffHeader is a Header carrying additional format information.
type ContainerHeader struct {
	*Header
	format uint32
}

// Format is a 4-letter format description.
func (rh *ContainerHeader) Format() string {
	val := make([]byte, 32)
	binary.BigEndian.PutUint32(val, rh.format)

	return strconv.Trim(val)
}

// String returns a string representation of header.
func (rh *ContainerHeader) String() string {
	return fmt.Sprintf("ID: %s Size: %d FullSize: %d StartPos: %d Format: %s", rh.ID(), rh.Size(), rh.FullSize(), rh.StartPos(), rh.Format())
}

// EncodeContainerHeader encodes provided id, size and format to ContainerHeader.
func EncodeContainerHeader(id FourCC, size uint32, format *FourCC, byteOrder binary.ByteOrder) *ContainerHeader {
	chunkHeader := EncodeChunkHeader(id, size, byteOrder)
	formatVal := format.ToUint32()

	return &ContainerHeader{Header: chunkHeader, format: formatVal}
}

// DecodeContainerHeader decodes container header from bytes.
func DecodeContainerHeader(bytes [ContainerHeaderSizeBytes]byte, byteOrder binary.ByteOrder) *ContainerHeader {
	var headerBytes [HeaderSizeBytes]byte
	copy(headerBytes[:], bytes[:HeaderSizeBytes])
	chunk := DecodeChunkHeader(headerBytes, 0, byteOrder)
	format := binary.BigEndian.Uint32(bytes[HeaderSizeBytes : HeaderSizeBytes+FormatSizeBytes])

	return &ContainerHeader{Header: chunk, format: format}
}

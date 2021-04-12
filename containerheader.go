package chunk

import (
	"encoding/binary"
	"errors"
	"fmt"
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

	return trim(val)
}

// String returns a string representation of header.
func (rh *ContainerHeader) String() string {
	return fmt.Sprintf("ID: %s Size: %d FullSize: %d StartPos: %d Format: %s", rh.ID(), rh.Size(), rh.FullSize(), rh.StartPos(), rh.Format())
}

// Bytes converts ContainerHeader to byte array.
// An amount of 12 bytes is returned.
func (rh *ContainerHeader) Bytes() []byte {
	bytes := rh.Header.Bytes()
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data[0:4], rh.format)
	bytes = append(bytes, data...)

	return bytes
}

// EncodeContainerHeader encodes provided id, size and format to ContainerHeader.
func EncodeContainerHeader(id FourCC, size uint32, format FourCC, byteOrder binary.ByteOrder) *ContainerHeader {
	chunkHeader := EncodeChunkHeader(id, size, byteOrder)
	formatVal := format.ToUint32()

	return &ContainerHeader{Header: chunkHeader, format: formatVal}
}

// DecodeContainerHeader decodes container header from bytes.
func DecodeContainerHeader(data []byte, byteOrder binary.ByteOrder) (*ContainerHeader, error) {
	if len(data) < int(ContainerHeaderSizeBytes) {
		msg := fmt.Sprintf("data slice requires a minimim lenght of %d", HeaderSizeBytes)
		return nil, errors.New(msg)
	}

	chunkHeader := decodeChunkHeader(data[:HeaderSizeBytes], 0, byteOrder)
	format := binary.BigEndian.Uint32(data[HeaderSizeBytes : HeaderSizeBytes+FormatSizeBytes])

	return &ContainerHeader{Header: chunkHeader, format: format}, nil
}

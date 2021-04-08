package chunk

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestEncodeContainerHeader(t *testing.T) {
	header := EncodeContainerHeader(CreateFourCC("RIFF"), 4, CreateFourCC("WAVE"), binary.LittleEndian)

	assertEqual(t, header.ID(), "RIFF", "ID")
	assertEqual(t, header.Size(), uint32(4), "Size")
	assertEqual(t, header.Format(), "WAVE", "Format")
}

func TestDecodeContainerHeader(t *testing.T) {
	header, err := DecodeContainerHeader(nil, binary.LittleEndian)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, header, "header should be nil")

	data := make([]byte, ContainerHeaderSizeBytes-1)
	header, err = DecodeContainerHeader(data, binary.LittleEndian)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, header, "header should be nil")

	data = make([]byte, ContainerHeaderSizeBytes)
	header, err = DecodeContainerHeader(data, binary.LittleEndian)

	assertNil(t, err, "err should be nil")
	assertNotNil(t, header, "header should not be nil")
	assertEqual(t, header.Format(), "", "Format")

	data = make([]byte, ContainerHeaderSizeBytes)

	format := CreateFourCC("WAVE")
	copy(data[HeaderSizeBytes:], format[:])
	header, err = DecodeContainerHeader(data, binary.LittleEndian)

	assertEqual(t, header.Format(), "WAVE", "Format")
}

func TestContainerHeaderBytes(t *testing.T) {
	data := make([]byte, 12)
	chunk, _ := DecodeContainerHeader(data, binary.LittleEndian)

	assertEqual(t, bytes.Compare(chunk.Bytes(), data), 0, "bytes")
}

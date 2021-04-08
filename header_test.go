package chunk

import (
	"encoding/binary"
	"testing"
)

func TestDecodeChunkHeader(t *testing.T) {
	var bytes [8]byte
	var offset uint32 = 0
	header := DecodeChunkHeader(bytes, offset, binary.LittleEndian)

	assertNotNil(t, header, "header should not be nil")
	assertEqual(t, header.StartPos(), uint32(0), "StartPos")

	offset = 12
	header = DecodeChunkHeader(bytes, offset, binary.LittleEndian)
	assertEqual(t, header.StartPos(), uint32(offset), "StartPos")
}

func TestEncodeChunkHeader(t *testing.T) {
	id := "Test"
	var size uint32 = 12
	header := createHeader(id, size, binary.LittleEndian)

	assertNotNil(t, header, "header should not be nil")
	assertEqual(t, header.ID(), id, "ID")
	assertEqual(t, header.Size(), size, "Size")
	assertEqual(t, header.FullSize(), HeaderSizeBytes+size, "FullSize")
	assertEqual(t, header.StartPos(), uint32(0), "StartPos")

	header = createHeader(id, size, binary.BigEndian)

	assertEqual(t, header.Size(), size, "Size")

	header = createHeader(id, 13, binary.BigEndian)

	assertEqual(t, header.FullSize(), HeaderSizeBytes+13+1, "FullSize")
}

func TestHeaderBytes(t *testing.T) {
	header := createHeader("Test", 12, binary.LittleEndian)
	headerBytes := header.Bytes()

	assertEqual(t, len(headerBytes), int(HeaderSizeBytes), "header bytes length")

	var bytes [HeaderSizeBytes]byte
	copy(bytes[:], headerBytes)
	decodedHeader := DecodeChunkHeader(bytes, 0, binary.LittleEndian)

	assertEqual(t, decodedHeader.ID(), header.ID(), "ID")
	assertEqual(t, decodedHeader.Size(), header.Size(), "Size")
}

func createHeader(idVal string, size uint32, byteOrder binary.ByteOrder) *Header {
	var id FourCC
	copy(id[:], idVal)

	return EncodeChunkHeader(id, size, byteOrder)
}

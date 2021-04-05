package chunk

import (
	"encoding/binary"
	"testing"
)

func TestDecodeChunkHeader(t *testing.T) {
	var bytes [8]byte
	var offset uint32 = 0
	header := DecodeChunkHeader(bytes, offset, binary.LittleEndian)

	if header == nil {
		t.Errorf("header should not be nil")
	}

	if header.StartPos() != 0 {
		t.Errorf("header StartPos is %d, want %d", header.StartPos(), 0)
	}

	offset = 12
	header = DecodeChunkHeader(bytes, offset, binary.LittleEndian)

	if header.StartPos() != offset {
		t.Errorf("header StartPos is %d, want %d", header.StartPos(), offset)
	}
}

func TestEncodeChunkHeader(t *testing.T) {
	id := "Test"
	var size uint32 = 12
	header := createHeader(id, size, binary.LittleEndian)

	if header == nil {
		t.Errorf("header should not be nil")
	}

	if header.ID() != id {
		t.Errorf("header ID is %s, want %s", header.ID(), id)
	}

	if header.Size() != size {
		t.Errorf("header size is %d, want %d", header.Size(), size)
	}

	if header.FullSize() != HeaderSizeBytes+size {
		t.Errorf("header FullSize is %d, want %d", header.FullSize(), HeaderSizeBytes+size)
	}

	if header.StartPos() != 0 {
		t.Errorf("header start pos is %d, want %d", header.StartPos(), 0)
	}

	header = createHeader(id, size, binary.BigEndian)

	if header.Size() != size {
		t.Errorf("header size is %d, want %d", header.Size(), size)
	}

	header = createHeader(id, 13, binary.BigEndian)

	if header.FullSize() != HeaderSizeBytes+13+1 {
		t.Errorf("header full size is %d, want %d", header.FullSize(), HeaderSizeBytes+13+1)
	}
}

func TestHeaderBytes(t *testing.T) {
	header := createHeader("Test", 12, binary.LittleEndian)
	headerBytes := header.Bytes()

	if len(headerBytes) != int(HeaderSizeBytes) {
		t.Errorf("header bytes length is %d, want %d", len(headerBytes), HeaderSizeBytes)
	}

	var bytes [HeaderSizeBytes]byte
	copy(bytes[:], headerBytes)
	decodedHeader := DecodeChunkHeader(bytes, 0, binary.LittleEndian)

	if decodedHeader.ID() != header.ID() {
		t.Errorf("header ID is %s, want %s", header.ID(), header.ID())
	}

	if decodedHeader.Size() != header.Size() {
		t.Errorf("header Size is %d, want %d", decodedHeader.Size(), header.Size())
	}
}

func createHeader(idVal string, size uint32, byteOrder binary.ByteOrder) *Header {
	var id FourCC
	copy(id[:], idVal)

	return EncodeChunkHeader(id, size, byteOrder)
}

package wav

import "testing"

func TestDecodeChunkHeader(t *testing.T) {
	var bytes [8]byte
	var offset uint32 = 0
	header := DecodeChunkHeader(bytes, offset)

	if header == nil {
		t.Errorf("header should not be nil")
	}

	if header.StartPos() != 0 {
		t.Errorf("header StartPos is %d, want %d", header.StartPos(), 0)
	}

	offset = 12
	header = DecodeChunkHeader(bytes, offset)

	if header.StartPos() != offset {
		t.Errorf("header StartPos is %d, want %d", header.StartPos(), offset)
	}
}

func TestEncodeChunkHeader(t *testing.T) {
	idVal := "Test"
	var id [IDSizeBytes]byte
	copy(id[:], idVal)
	var size uint32 = 12
	header := EncodeChunkHeader(id, size)

	if header == nil {
		t.Errorf("header should not be nil")
	}

	if header.ID() != idVal {
		t.Errorf("header ID is %s, want %s", header.ID(), idVal)
	}

	if header.Size() != size {
		t.Errorf("header Size is %d, want %d", header.Size(), size)
	}

	if header.FullSize() != HeaderSizeBytes+size {
		t.Errorf("header FullSize is %d, want %d", header.FullSize(), HeaderSizeBytes+size)
	}

	if header.StartPos() != 0 {
		t.Errorf("header StartPos is %d, want %d", header.StartPos(), 0)
	}
}

func TestBytes(t *testing.T) {
	idVal := "Test"
	var id [IDSizeBytes]byte
	copy(id[:], idVal)
	var size uint32 = 12
	header := EncodeChunkHeader(id, size)
	headerBytes := header.Bytes()

	if len(headerBytes) != int(HeaderSizeBytes) {
		t.Errorf("header bytes length is %d, want %d", len(headerBytes), HeaderSizeBytes)
	}

	var bytes [HeaderSizeBytes]byte
	copy(bytes[:], headerBytes)
	decodedHeader := DecodeChunkHeader(bytes, 0)

	if decodedHeader.ID() != idVal {
		t.Errorf("header ID is %s, want %s", header.ID(), idVal)
	}

	if decodedHeader.Size() != size {
		t.Errorf("header Size is %d, want %d", header.Size(), size)
	}
}

package wav

import "testing"

func TestDecodeChunkHeader(t *testing.T) {
	var bytes [8]byte
	var offset uint32 = 0
	header := DecodeChunkHeader(bytes, offset)

	if header == nil {
		t.Errorf("header should not be nil")
	}
}

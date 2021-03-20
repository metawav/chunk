package wav

import "testing"

func TestDecodeChunkHeader(t *testing.T) {
	var bytes []byte = nil
	var offset uint32 = 0
	header, err := DecodeChunkHeader(bytes, offset)

	if err == nil {
		t.Errorf("error should be returned with size %d", len(bytes))
	}

	if header != nil {
		t.Errorf("header should not be valid with size %d", len(bytes))
	}

	bytes = make([]byte, HeaderSizeBytes-1)
	header, err = DecodeChunkHeader(bytes, offset)

	if err == nil {
		t.Errorf("error should be returned with size %d", len(bytes))
	}

	if header != nil {
		t.Errorf("header should not be valid with size %d", len(bytes))
	}

	bytes = make([]byte, HeaderSizeBytes)
	header, err = DecodeChunkHeader(bytes, offset)

	if err != nil {
		t.Errorf("error should not be returned with size %d", len(bytes))
	}

	if header == nil {
		t.Errorf("header should be valid with size %d", len(bytes))
	}
}

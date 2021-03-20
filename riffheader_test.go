package wav

import "testing"

func TestDecodeRiffHeader(t *testing.T) {
	var bytes []byte = nil
	header, err := DecodeRiffHeader(bytes)

	if err == nil {
		t.Errorf("error should be returned with size %d", len(bytes))
	}

	if header != nil {
		t.Errorf("header should not be valid with size %d", len(bytes))
	}

	bytes = make([]byte, RiffHeaderSizeBytes-1)
	header, err = DecodeRiffHeader(bytes)

	if err == nil {
		t.Errorf("error should be returned with size %d", len(bytes))
	}

	if header != nil {
		t.Errorf("header should not be valid with size %d", len(bytes))
	}

	bytes = make([]byte, RiffHeaderSizeBytes)
	header, err = DecodeRiffHeader(bytes)

	if err != nil {
		t.Errorf("error should not be returned with size %d", len(bytes))
	}

	if header == nil {
		t.Errorf("header should be valid with size %d", len(bytes))
	}
}

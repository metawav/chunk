package wav

import "testing"

func TestDecodeRiffHeader(t *testing.T) {
	var bytes [RiffHeaderSizeBytes]byte
	header := DecodeRiffHeader(bytes)

	if header == nil {
		t.Errorf("header should not be nil")
	}
}

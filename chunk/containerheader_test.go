package chunk

import (
	"encoding/binary"
	"testing"
)

func TestDecodeContainerHeader(t *testing.T) {
	var bytes [ContainerHeaderSizeBytes]byte
	header := DecodeContainerHeader(bytes, binary.LittleEndian)

	if header == nil {
		t.Errorf("header should not be nil")
	}
}

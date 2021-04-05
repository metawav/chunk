package chunk

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestEncodeContainerHeader(t *testing.T) {
	header := EncodeContainerHeader(CreateFourCC("RIFF"), 4, CreateFourCC("WAVE"), binary.LittleEndian)

	if header.ID() != "RIFF" {
		t.Errorf("ID is %s, want %s", header.ID(), "RIFF")
	}

	if header.Size() != 4 {
		t.Errorf("size is %d, want %d", header.Size(), 4)
	}

	if header.Format() != "WAVE" {
		t.Errorf("format is %s, want %s", header.Format(), "WAVE")
	}
}

func TestDecodeContainerHeader(t *testing.T) {
	header, err := DecodeContainerHeader(nil, binary.LittleEndian)

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if header != nil {
		t.Errorf("header should be be nil")
	}

	data := make([]byte, ContainerHeaderSizeBytes-1)
	header, err = DecodeContainerHeader(data, binary.LittleEndian)

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if header != nil {
		t.Errorf("header should be nil")
	}

	data = make([]byte, ContainerHeaderSizeBytes)
	header, err = DecodeContainerHeader(data, binary.LittleEndian)

	if err != nil {
		t.Errorf("err should be nil")
	}

	if header == nil {
		t.Errorf("header should not be nil")
	}

	if header.Format() != "" {
		t.Errorf("format is %s, want %s", header.Format(), "")
	}

	data = make([]byte, ContainerHeaderSizeBytes)

	format := CreateFourCC("WAVE")
	copy(data[HeaderSizeBytes:], format[:])
	header, err = DecodeContainerHeader(data, binary.LittleEndian)

	if header.Format() != "WAVE" {
		t.Errorf("format is %s, want %s", header.Format(), "WAVE")
	}
}

func TestContainerHeaderBytes(t *testing.T) {
	data := make([]byte, 12)
	chunk, _ := DecodeContainerHeader(data, binary.LittleEndian)

	if bytes.Compare(chunk.Bytes(), data) != 0 {
		t.Errorf("bytes is %v, want %v", chunk.Bytes(), data)
	}
}

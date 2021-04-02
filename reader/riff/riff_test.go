package riff

import (
	"encoding/binary"
	"strings"
	"testing"

	"github.com/pmoule/wav/chunk"
)

func TestRead(t *testing.T) {
	fileName := "testFile"
	content := make([]byte, chunk.ContainerHeaderSizeBytes-1)
	reader := strings.NewReader(string(content))
	riff, err := Read(fileName, reader)

	if err == nil {
		t.Errorf("error should be returned with file %s", fileName)
	}

	if riff != nil {
		t.Errorf("riff file should not be returned with file %s", fileName)
	}

	content = make([]byte, chunk.ContainerHeaderSizeBytes)
	reader = strings.NewReader(string(content))
	riff, err = Read(fileName, reader)

	if err != nil {
		t.Errorf("error should not be returned with file %s", fileName)
	}

	if riff == nil {
		t.Errorf("riff file should be returned with file %s", fileName)
	}

	if len(riff.Headers) != 0 {
		t.Errorf("headers have a length of %d, wanted 0", len(riff.Headers))
	}

	if riff.Headers != nil {
		t.Errorf("headers should be nil")
	}

	content = make([]byte, chunk.ContainerHeaderSizeBytes+chunk.HeaderSizeBytes)
	reader = strings.NewReader(string(content))
	riff, err = Read(fileName, reader)

	if len(riff.Headers) != 1 {
		t.Errorf("headers have a length of %d, wanted 1", len(riff.Headers))
	}

	headerID := createID("test")
	var headerSize int = 12
	header := chunk.EncodeChunkHeader(headerID, uint32(headerSize), binary.LittleEndian)
	headerBytes := header.Bytes()
	headerBytes = append(headerBytes, make([]byte, headerSize)...)
	content = make([]byte, chunk.ContainerHeaderSizeBytes)
	content = append(content, headerBytes...)
	content = append(content, headerBytes...)
	reader = strings.NewReader(string(content))
	riff, err = Read(fileName, reader)

	if len(riff.Headers) != 2 {
		t.Errorf("headers have a length of %d, wanted 2", len(riff.Headers))
	}
}

func createID(idVal string) [chunk.IDSizeBytes]byte {
	var id [chunk.IDSizeBytes]byte
	copy(id[:], idVal)

	return id
}

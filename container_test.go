package chunk

import (
	"encoding/binary"
	"fmt"
	"strings"
	"testing"
)

func TestFindHeaders(t *testing.T) {
	riffFile := &Container{}

	for _, h := range riffFile.Headers {
		fmt.Printf("%v\n", h)
	}

	headers := riffFile.FindHeaders("")

	if headers != nil {
		t.Errorf("headers should not be returned")
	}

	headerID := createID("test")
	header := EncodeChunkHeader(headerID, 0, binary.LittleEndian)

	riffFile.Headers = append(riffFile.Headers, header)
	headers = riffFile.FindHeaders(string(headerID[:]))

	if headers == nil {
		t.Errorf("headers should be returned")
	}

	if len(headers) != 1 {
		t.Errorf("headers size is %d, want %d", len(headers), 1)
	}

	riffFile.Headers = append(riffFile.Headers, header)
	headers = riffFile.FindHeaders(string(headerID[:]))

	if len(headers) != 2 {
		t.Errorf("headers size is %d, want %d", len(headers), 2)
	}
}

func TestReadRiff(t *testing.T) {
	fileName := "testFile"
	content := make([]byte, ContainerHeaderSizeBytes-1)
	reader := strings.NewReader(string(content))
	riff, err := ReadRiff(fileName, reader)

	if err == nil {
		t.Errorf("error should be returned with file %s", fileName)
	}

	if riff != nil {
		t.Errorf("riff file should not be returned with file %s", fileName)
	}

	content = make([]byte, ContainerHeaderSizeBytes)
	reader = strings.NewReader(string(content))
	riff, err = ReadRiff(fileName, reader)

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

	content = make([]byte, ContainerHeaderSizeBytes+HeaderSizeBytes)
	reader = strings.NewReader(string(content))
	riff, err = ReadRiff(fileName, reader)

	if len(riff.Headers) != 1 {
		t.Errorf("headers have a length of %d, wanted 1", len(riff.Headers))
	}

	headerID := createID("test")
	var headerSize int = 12
	header := EncodeChunkHeader(headerID, uint32(headerSize), binary.LittleEndian)
	headerBytes := header.Bytes()
	headerBytes = append(headerBytes, make([]byte, headerSize)...)
	content = make([]byte, ContainerHeaderSizeBytes)
	content = append(content, headerBytes...)
	content = append(content, headerBytes...)
	reader = strings.NewReader(string(content))
	riff, err = ReadRiff(fileName, reader)

	if len(riff.Headers) != 2 {
		t.Errorf("headers have a length of %d, wanted 2", len(riff.Headers))
	}
}

func createID(idVal string) [IDSizeBytes]byte {
	var id [IDSizeBytes]byte
	copy(id[:], idVal)

	return id
}

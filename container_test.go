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

	assertEqual(t, len(headers), 1, "headers length")

	riffFile.Headers = append(riffFile.Headers, header)
	headers = riffFile.FindHeaders(string(headerID[:]))

	assertEqual(t, len(headers), 2, "headers length")
}

func TestReadRiff(t *testing.T) {
	fileName := "testFile"
	content := make([]byte, ContainerHeaderSizeBytes-1)
	reader := strings.NewReader(string(content))
	riff, err := ReadRiff(fileName, reader)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, riff, "riff should be nil")

	content = make([]byte, ContainerHeaderSizeBytes)
	reader = strings.NewReader(string(content))
	riff, err = ReadRiff(fileName, reader)

	assertNil(t, err, "err should be nil")
	assertNotNil(t, riff, "riff should not be nil")
	assertEqual(t, len(riff.Headers), 0, "headers length")

	if riff.Headers != nil {
		t.Errorf("headers should be nil")
	}

	content = make([]byte, ContainerHeaderSizeBytes+HeaderSizeBytes)
	reader = strings.NewReader(string(content))
	riff, err = ReadRiff(fileName, reader)

	assertEqual(t, len(riff.Headers), 1, "headers length")

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

	assertEqual(t, len(riff.Headers), 2, "headers length")
}

func createID(idVal string) [IDSizeBytes]byte {
	var id [IDSizeBytes]byte
	copy(id[:], idVal)

	return id
}

package wav

import (
	"encoding/binary"
	"fmt"
	"strings"
	"testing"
)

func TestCreateRiffFile(t *testing.T) {
	fileName := "testFile"
	content := make([]byte, RiffHeaderSizeBytes-1)
	reader := strings.NewReader(string(content))
	riff, err := CreateRiffFile(fileName, reader)

	if err == nil {
		t.Errorf("error should be returned with file %s", fileName)
	}

	if riff != nil {
		t.Errorf("riff file should not be returned with file %s", fileName)
	}

	content = make([]byte, RiffHeaderSizeBytes)
	reader = strings.NewReader(string(content))
	riff, err = CreateRiffFile(fileName, reader)

	if err != nil {
		t.Errorf("error should not be returned with file %s", fileName)
	}

	if riff == nil {
		t.Errorf("riff file should be returned with file %s", fileName)
	}

	if len(riff.Headers) != 0 {
		t.Errorf("headers have a length of %d, , wanted 0", len(riff.Headers))
	}

	if riff.Headers != nil {
		t.Errorf("headers should be nil")
	}

	content = make([]byte, RiffHeaderSizeBytes+HeaderSizeBytes)
	reader = strings.NewReader(string(content))
	riff, err = CreateRiffFile(fileName, reader)

	if len(riff.Headers) != 1 {
		t.Errorf("headers have a length of %d, , wanted 1", len(riff.Headers))
	}

	var headerSize uint32 = 12
	header := &Header{size: headerSize}
	headerBytes := header.Bytes()
	headerBytes = append(headerBytes, make([]byte, 12)...)
	content = make([]byte, RiffHeaderSizeBytes)
	content = append(content, headerBytes...)
	content = append(content, make([]byte, HeaderSizeBytes)...)
	reader = strings.NewReader(string(content))
	riff, err = CreateRiffFile(fileName, reader)

	if len(riff.Headers) != 2 {
		t.Errorf("headers have a length of %d, , wanted 2", len(riff.Headers))
	}
}

func TestGetHeaderByID(t *testing.T) {
	riffFile := &RiffFile{}

	for _, h := range riffFile.Headers {
		fmt.Printf("%v\n", h)
	}

	header := riffFile.GetHeaderByID("")

	if header != nil {
		t.Errorf("header should not be returned")
	}

	headerID := "test"
	id := binary.BigEndian.Uint32([]byte(headerID)[:4])
	header = &Header{id: id}
	riffFile.Headers = append(riffFile.Headers, header)
	header = riffFile.GetHeaderByID(headerID)

	if header == nil {
		t.Errorf("header should be returned")
	}
}

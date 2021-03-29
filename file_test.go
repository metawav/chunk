package wav

import (
	"encoding/binary"
	"fmt"
	"strings"
	"testing"
)

func TestCreateRiffFile(t *testing.T) {
	fileName := "testFile"
	reader := strings.NewReader("This ab")
	riff, err := CreateRiffFile(fileName, reader)

	if err != nil {
		t.Errorf("error should bot be returned with file %s", fileName)
	}

	if riff == nil {
		t.Errorf("riff file should be returned with file %s", fileName)
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

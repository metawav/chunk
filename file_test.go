package wav

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"
)

func TestCreateRiffFile(t *testing.T) {
	var file *os.File
	riffFile, err := CreateRiffFile(file)

	if err == nil {
		t.Errorf("error should be returned with file %v", file)
	}

	if riffFile != nil {
		t.Errorf("riff file should not be returned with file %v", file)
	}

	fileName := "sine.1sec.wav"
	file, _ = OpenFile(fileName)

	riffFile, err = CreateRiffFile(file)

	if err != nil {
		t.Errorf("error should be returned with file %s", file.Name())
	}

	if riffFile == nil {
		t.Errorf("riff file should be returned with file %s", file.Name())
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

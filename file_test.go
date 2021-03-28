package wav

import (
	"encoding/binary"
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
	riffFile := NewRiffFile("", nil, []*Header{})

	header, err := riffFile.GetHeaderByID("")

	if err == nil {
		t.Errorf("error should be returned")
	}

	if header != nil {
		t.Errorf("header should not be returned")
	}

	headerID := "test"
	id := binary.BigEndian.Uint32([]byte(headerID)[:4])
	header = &Header{id: id}
	riffFile.headers = append(riffFile.headers, header)
	header, err = riffFile.GetHeaderByID(headerID)

	if err != nil {
		t.Errorf("error should not be returned, %s", err)
	}

	if header == nil {
		t.Errorf("header should be returned")
	}
}

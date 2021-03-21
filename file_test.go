package wav

import (
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

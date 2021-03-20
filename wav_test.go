package wav

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	fileName := ""
	file, err := ReadFile(fileName)

	if err == nil {
		t.Errorf("should return error")
	}

	if file != nil {
		t.Errorf("file should be nil")
	}

	fileName = "sine.1sec.wav"
	file, err = ReadFile(fileName)

	if err != nil {
		t.Errorf("should not return error: %+v", err)
	}

	if file == nil {
		t.Errorf("file should not be nil")
	}
}

package wav

import (
	"errors"
	"fmt"
	"os"
)

// ReadFile
func ReadFile(name string) (*RiffFile, error) {
	file, err := openFile(name)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	riffFile, err := CreateRiffFile(file)

	return riffFile, err
}

func openFile(name string) (*os.File, error) {
	if !isFileExist(name) {
		msg := fmt.Sprintf("file does not exist: %s", name)
		return nil, errors.New(msg)
	}

	file, err := os.Open(name)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func isFileExist(name string) bool {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

package aiff

import (
	"encoding/binary"
	"io"

	"github.com/pmoule/wav/chunk"
	"github.com/pmoule/wav/reader/internal"
)

// Read reads a AIFF / AIFF-C file from provided byte stream and assigns provided name.
func Read(name string, reader io.ReadSeeker) (*chunk.Container, error) {
	return internal.Read(name, reader, binary.BigEndian)
}

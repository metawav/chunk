package wav

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// RiffHeader is a Header carrying additional format information
type RiffHeader struct {
	*Header
	format uint32
}

// NewRiffHeader creates a new RIFF header
func NewRiffHeader(header *Header, format uint32) *RiffHeader {
	riffHeader := &RiffHeader{Header: header, format: format}
	riffHeader.format = format

	return riffHeader
}

// DecodeRiffHeader decodes RIFF header from bytes
func DecodeRiffHeader(bytes []byte) (*RiffHeader, error) {
	if !isValidRiffHeader(bytes) {
		msg := fmt.Sprintf("invalid riff header")
		return nil, errors.New(msg)
	}

	chunk, err := DecodeChunkHeader(bytes, 0)

	if err != nil {
		return nil, err
	}

	format := binary.BigEndian.Uint32(bytes[HeaderSizeBytes : HeaderSizeBytes+FormatSizeBytes])

	return NewRiffHeader(chunk, format), nil
}

// Format is a 4-letter format description
func (rh *RiffHeader) Format() string {
	val := make([]byte, 32)
	binary.BigEndian.PutUint32(val, rh.format)

	return trim(val)
}

// Size is the chunk size in bytes
// not including: id (4 bytes), size (4 bytes) and format (4 bytes)
func (rh *RiffHeader) Size() uint32 {
	return rh.size - FormatSizeBytes
}

// FullSize is the chunk size in bytes
func (rh *RiffHeader) FullSize() uint32 {
	return rh.Header.Size() + RiffHeaderSizeBytes
}

// String returns a stringrepresentation of header
func (rh *RiffHeader) String() string {
	return fmt.Sprintf("ID: %s Size: %d FullSize: %d StartPos: %d Format: %s", rh.ID(), rh.Size(), rh.FullSize(), rh.StartPos(), rh.Format())
}

func isValidRiffHeader(bytes []byte) bool {
	length := len(bytes)

	if uint32(length) < RiffHeaderSizeBytes {
		return false
	}

	return true
}

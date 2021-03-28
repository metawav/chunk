package wav

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Header carries the following information of a chunk: ID, size and start position in a RIFF file
type Header struct {
	id       uint32
	size     uint32
	startPos uint32
}

// NewHeader creates a new header
func NewHeader(id uint32, size uint32, startPos uint32) *Header {
	header := &Header{id: id, size: size, startPos: startPos}

	return header
}

// EncodeChunkHeader encodes
func EncodeChunkHeader(id [4]byte, size uint32, startPos uint32) *Header {
	idVal := binary.BigEndian.Uint32(id[:])
	header := NewHeader(idVal, size, startPos)

	return header
}

// DecodeChunkHeader decodes chunk header from bytes
func DecodeChunkHeader(bytes []byte, startPos uint32) (*Header, error) {
	if !isValidChunkHeader(bytes) {
		msg := fmt.Sprintf("invalid header")
		return nil, errors.New(msg)
	}

	id := binary.BigEndian.Uint32(bytes[:IDSizeBytes])
	size := binary.LittleEndian.Uint32(bytes[IDSizeBytes : IDSizeBytes+SizeSizeBytes])

	return NewHeader(id, size, startPos), nil
}

// ID is a 4-letter chunk identifier
func (h *Header) ID() string {
	val := make([]byte, 4)
	binary.BigEndian.PutUint32(val, h.id)

	return trim(val)
}

// StartPos is the starting position of this chunk in the byte sequence
func (h *Header) StartPos() uint32 {
	return h.startPos
}

// Size is the chunk size in bytes
// not including: id (4 bytes) and size (4 bytes)
func (h *Header) Size() uint32 {
	return h.size
}

// FullSize is the chunk size in bytes
func (h *Header) FullSize() uint32 {
	return h.size + HeaderSizeBytes
}

func isValidChunkHeader(bytes []byte) bool {
	length := len(bytes)

	if uint32(length) < HeaderSizeBytes {
		return false
	}

	return true
}

// Bytes converts Header to  byte array
func (h *Header) Bytes() []byte {
	bytes := make([]byte, HeaderSizeBytes)
	binary.BigEndian.PutUint32(bytes[:IDSizeBytes], h.id)
	binary.LittleEndian.PutUint32(bytes[IDSizeBytes:IDSizeBytes+SizeSizeBytes], h.size)

	return bytes
}

// String returns a stringrepresentation of header
func (h *Header) String() string {
	return fmt.Sprintf("ID: %s Size: %d FullSize: %d StartPos: %d", h.ID(), h.Size(), h.FullSize(), h.StartPos())
}

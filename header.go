package wav

import (
	"encoding/binary"
	"fmt"
)

// Header carries the following information of a chunk: ID, size and start position in a RIFF file
type Header struct {
	id       uint32
	size     uint32
	startPos uint32
}

// SortHeadersByStartPosAsc sorts headers by start position in ascending order.
type SortHeadersByStartPosAsc []*Header

func (a SortHeadersByStartPosAsc) Len() int           { return len(a) }
func (a SortHeadersByStartPosAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortHeadersByStartPosAsc) Less(i, j int) bool { return a[i].StartPos() < a[j].StartPos() }

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
// not including: id (4 bytes) and size (4 bytes).
func (h *Header) Size() uint32 {
	return h.size
}

// FullSize is the chunk size in bytes.
func (h *Header) FullSize() uint32 {
	return h.size + HeaderSizeBytes
}

// String returns a stringrepresentation of header
func (h *Header) String() string {
	return fmt.Sprintf("ID: %s Size: %d FullSize: %d StartPos: %d", h.ID(), h.Size(), h.FullSize(), h.StartPos())
}

// Bytes converts Header to  byte array.
func (h *Header) Bytes() []byte {
	bytes := make([]byte, HeaderSizeBytes)
	binary.BigEndian.PutUint32(bytes[:IDSizeBytes], h.id)
	binary.LittleEndian.PutUint32(bytes[IDSizeBytes:IDSizeBytes+SizeSizeBytes], h.size)

	return bytes
}

// EncodeChunkHeader encodes provided id and size to Header.
func EncodeChunkHeader(id [IDSizeBytes]byte, size uint32) *Header {
	idVal := binary.BigEndian.Uint32(id[:])

	return &Header{id: idVal, size: size}
}

// DecodeChunkHeader decodes chunk header from bytes.
func DecodeChunkHeader(bytes [HeaderSizeBytes]byte, startPos uint32) *Header {
	id := binary.BigEndian.Uint32(bytes[:IDSizeBytes])
	size := binary.LittleEndian.Uint32(bytes[IDSizeBytes : IDSizeBytes+SizeSizeBytes])

	return &Header{id: id, size: size, startPos: startPos}
}

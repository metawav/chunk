package chunk

import (
	"encoding/binary"
	"fmt"
)

// Header carries the following information of a chunk: ID, size and start position in a RIFF file
type Header struct {
	id        uint32
	size      uint32
	startPos  uint32
	byteOrder binary.ByteOrder
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
// not including: id (4 bytes) and size (4 bytes).
func (h *Header) Size() uint32 {
	return h.size
}

// FullSize is the chunk size in bytes including optional padding byte.
func (h *Header) FullSize() uint32 {
	size := h.size + HeaderSizeBytes

	if size%2 != 0 {
		size++
	}

	return size
}

// HasPadding returns true when padding byte is added as chunk must be even sized and must start at an even position.
func (h *Header) HasPadding() bool {
	isOdd := h.Size()%2 != 0

	return isOdd
}

// String returns a stringrepresentation of header
func (h *Header) String() string {
	return fmt.Sprintf("ID: %s Size: %d FullSize: %d StartPos: %d HasPadding: %t", h.ID(), h.Size(), h.FullSize(), h.StartPos(), h.HasPadding())
}

// Bytes converts Header to byte array.
// An amount of 8 bytes is returned.
func (h *Header) Bytes() []byte {
	bytes := make([]byte, HeaderSizeBytes)
	binary.BigEndian.PutUint32(bytes[:IDSizeBytes], h.id)
	h.byteOrder.PutUint32(bytes[IDSizeBytes:IDSizeBytes+SizeSizeBytes], h.size)

	return bytes
}

// EncodeChunkHeader encodes provided id and size to Header.
func EncodeChunkHeader(id FourCC, size uint32, byteOrder binary.ByteOrder) *Header {
	idVal := id.ToUint32()

	return &Header{id: idVal, size: size, byteOrder: byteOrder}
}

// DecodeChunkHeader decodes chunk header from bytes.
func DecodeChunkHeader(bytes [HeaderSizeBytes]byte, startPos uint32, byteOrder binary.ByteOrder) *Header {
	id := binary.BigEndian.Uint32(bytes[:IDSizeBytes])
	size := byteOrder.Uint32(bytes[IDSizeBytes : IDSizeBytes+SizeSizeBytes])

	return &Header{id: id, size: size, startPos: startPos, byteOrder: byteOrder}
}

func decodeChunkHeader(bytes []byte, startPos uint32, byteOrder binary.ByteOrder) *Header {
	var headerBytes [HeaderSizeBytes]byte
	copy(headerBytes[:], bytes[:HeaderSizeBytes])
	chunkHeader := DecodeChunkHeader(headerBytes, startPos, byteOrder)

	return chunkHeader
}

// pad adds a pad byte with value 0 if data size is odd.
func pad(data []byte) []byte {
	if len(data)%2 == 0 {
		return data
	}

	var pad byte
	return append(data, pad)
}

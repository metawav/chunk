package chunk

import (
	"encoding/binary"
	"testing"
)

func TestEncodeFMThunk(t *testing.T) {
	chunk := EncodeFMTChunk(14, 0, 1, 44100, 1000, 2)

	assertEqual(t, chunk.ID(), FMTID, "ID")
	assertEqual(t, chunk.Size(), uint32(14), "Size")
	assertEqual(t, chunk.Format(), 0, "Format")
	assertEqual(t, chunk.Channels(), 1, "Channels")
	assertEqual(t, chunk.SamplesPerSec(), 44100, "SamplesPerSec")
	assertEqual(t, chunk.BytesPerSec(), 1000, "BytesPerSec")
	assertEqual(t, chunk.BlockAlign(), 2, "BlockAlign")
}

func TestDecodeFMTChunk(t *testing.T) {
	chunk, err := DecodeFMTChunk(nil)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, chunk, "chunk should be nil")

	data := make([]byte, HeaderSizeBytes-1)
	chunk, err = DecodeFMTChunk(data)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, chunk, "chunk should be nil")

	data = make([]byte, HeaderSizeBytes)
	chunk, err = DecodeFMTChunk(data)

	assertNotNil(t, err, "err should not be nil when DecodeFMTChunk with not enough data")
	assertNotNil(t, chunk.Header, "header should not be nil")
	assertEqual(t, chunk.Format(), 0, "Format")
	assertEqual(t, chunk.Channels(), 0, "Channels")
	assertEqual(t, chunk.SamplesPerSec(), 0, "SamplesPerSec")
	assertEqual(t, chunk.BytesPerSec(), 0, "BytesPerSec")
	assertEqual(t, chunk.BlockAlign(), 0, "BlockAlign")

	data2 := make([]byte, 22)
	var format uint16 = 1
	binary.LittleEndian.PutUint16(data2[HeaderSizeBytes:HeaderSizeBytes+2], format)
	var channels uint16 = 2
	binary.LittleEndian.PutUint16(data2[HeaderSizeBytes+2:HeaderSizeBytes+4], channels)
	var samplesPerSec uint32 = 2
	binary.LittleEndian.PutUint32(data2[HeaderSizeBytes+4:HeaderSizeBytes+8], samplesPerSec)
	var bytesPerSec uint32 = 2
	binary.LittleEndian.PutUint32(data2[HeaderSizeBytes+8:HeaderSizeBytes+12], bytesPerSec)
	var blockAlign uint16 = 5
	binary.LittleEndian.PutUint16(data2[HeaderSizeBytes+12:HeaderSizeBytes+14], blockAlign)
	chunk, _ = DecodeFMTChunk(data2)

	assertEqual(t, chunk.Format(), int(format), "Format")
	assertEqual(t, chunk.Channels(), int(channels), "Channels")
	assertEqual(t, chunk.SamplesPerSec(), int(samplesPerSec), "SamplesPerSec")
	assertEqual(t, chunk.BytesPerSec(), int(bytesPerSec), "BytesPerSec")
	assertEqual(t, chunk.BlockAlign(), int(blockAlign), "BlockAlign")
}

func TestFMTBytes(t *testing.T) {
	const formatSize = 22
	data := make([]byte, formatSize)
	var format uint16 = 1
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes:HeaderSizeBytes+2], format)
	var channels uint16 = 2
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes+2:HeaderSizeBytes+4], channels)
	var samplesPerSec uint32 = 2
	binary.LittleEndian.PutUint32(data[HeaderSizeBytes+4:HeaderSizeBytes+8], samplesPerSec)
	var bytesPerSec uint32 = 2
	binary.LittleEndian.PutUint32(data[HeaderSizeBytes+8:HeaderSizeBytes+12], bytesPerSec)
	var blockAlign uint16 = 5
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes+12:HeaderSizeBytes+14], blockAlign)
	chunk, _ := DecodeFMTChunk(data)

	assertEqual(t, len(chunk.Bytes()), formatSize, "bytes length after DecodeFMTChunk")

	data = make([]byte, 8)
	chunk, _ = DecodeFMTChunk(data)

	assertEqual(t, len(chunk.Bytes()), formatSize, "bytes length after DecodeFMTChunk")
}

func TestEncodePCMFormathunk(t *testing.T) {
	chunk := EncodePCMFormatChunk(14, 0, 1, 44100, 1000, 2, 24)

	assertEqual(t, chunk.ID(), FMTID, "ID")
	assertEqual(t, chunk.Size(), uint32(14), "Size")
	assertEqual(t, chunk.Format(), 0, "Format")
	assertEqual(t, chunk.Channels(), 1, "Channels")
	assertEqual(t, chunk.SamplesPerSec(), 44100, "SamplesPerSec")
	assertEqual(t, chunk.BytesPerSec(), 1000, "BytesPerSec")
	assertEqual(t, chunk.BlockAlign(), 2, "BlockAlign")
	assertEqual(t, chunk.BitsPerSample(), 24, "BitsPerSample")
}

func TestDecodePCMFormatChunk(t *testing.T) {
	chunk, err := DecodePCMFormatChunk(nil)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, chunk, "chunk should be nil")

	dataMinSize := 22
	data := make([]byte, dataMinSize-1)
	chunk, err = DecodePCMFormatChunk(data)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, chunk, "chunk should be nil")

	fmtData := make([]byte, 22)
	var format uint16 = 1
	binary.LittleEndian.PutUint16(fmtData[HeaderSizeBytes:HeaderSizeBytes+2], format)
	var channels uint16 = 2
	binary.LittleEndian.PutUint16(fmtData[HeaderSizeBytes+2:HeaderSizeBytes+4], channels)
	var samplesPerSec uint32 = 2
	binary.LittleEndian.PutUint32(fmtData[HeaderSizeBytes+4:HeaderSizeBytes+8], samplesPerSec)
	var bytesPerSec uint32 = 2
	binary.LittleEndian.PutUint32(fmtData[HeaderSizeBytes+8:HeaderSizeBytes+12], bytesPerSec)
	var blockAlign uint16 = 5
	binary.LittleEndian.PutUint16(fmtData[HeaderSizeBytes+12:HeaderSizeBytes+14], blockAlign)
	fmt, _ := DecodeFMTChunk(fmtData)

	data = make([]byte, len(fmt.Bytes())-1)
	chunk, err = DecodePCMFormatChunk(data)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, chunk, "chunk should be nil")

	data = make([]byte, len(fmt.Bytes()))
	chunk, err = DecodePCMFormatChunk(data)

	assertNotNil(t, err, "err should not be nil")
	assertNotNil(t, chunk, "chunk should not be nil")
	assertEqual(t, chunk.BitsPerSample(), 0, "BitsPerSample")

	data = make([]byte, 2)
	var bitsPerSample uint16 = 24
	binary.LittleEndian.PutUint16(data[0:2], bitsPerSample)
	data = append(fmtData, data...)
	chunk, err = DecodePCMFormatChunk(data)

	assertNil(t, err, "err should be nil")
	assertEqual(t, chunk.BitsPerSample(), int(bitsPerSample), "BitsPerSample")
}

func TestPCMFormatBytes(t *testing.T) {
	const pcmFormatSize = 24
	data := make([]byte, pcmFormatSize)
	var format uint16 = 1
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes:HeaderSizeBytes+2], format)
	var channels uint16 = 2
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes+2:HeaderSizeBytes+4], channels)
	var samplesPerSec uint32 = 2
	binary.LittleEndian.PutUint32(data[HeaderSizeBytes+4:HeaderSizeBytes+8], samplesPerSec)
	var bytesPerSec uint32 = 2
	binary.LittleEndian.PutUint32(data[HeaderSizeBytes+8:HeaderSizeBytes+12], bytesPerSec)
	var blockAlign uint16 = 5
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes+12:HeaderSizeBytes+14], blockAlign)
	var bitsPerSample uint16 = 24
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes+14:HeaderSizeBytes+16], bitsPerSample)
	chunk, _ := DecodePCMFormatChunk(data)

	assertEqual(t, len(chunk.Bytes()), pcmFormatSize, "bytes length after DecodePCMFormatChunk")

	data = make([]byte, 22)
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes:HeaderSizeBytes+2], format)
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes+2:HeaderSizeBytes+4], channels)
	binary.LittleEndian.PutUint32(data[HeaderSizeBytes+4:HeaderSizeBytes+8], samplesPerSec)
	binary.LittleEndian.PutUint32(data[HeaderSizeBytes+8:HeaderSizeBytes+12], bytesPerSec)
	binary.LittleEndian.PutUint16(data[HeaderSizeBytes+12:HeaderSizeBytes+14], blockAlign)
	chunk, _ = DecodePCMFormatChunk(data)

	assertEqual(t, len(chunk.Bytes()), pcmFormatSize, "bytes length after DecodePCMFormatChunk")
}

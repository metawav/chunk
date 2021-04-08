package chunk

import (
	"encoding/binary"
	"testing"
)

func TestEncodeCommChunk(t *testing.T) {
	chunk := EncodeCOMMChunk(32, 2, 100, 200, 44100, CreateFourCC("NONE"), "no compression")

	assertEqual(t, chunk.ID(), COMMID, "ID")
	assertEqual(t, chunk.Size(), uint32(32), "Size")
	assertEqual(t, chunk.Channels(), 2, "Channels")
	assertEqual(t, chunk.SampleFrames(), 100, "SampleFrames")
	assertEqual(t, chunk.SampleSize(), 200, "SampleSize")
	assertEqual(t, chunk.SampleRate(), uint(44100), "SampleRate")
	assertEqual(t, chunk.CompressionType(), "NONE", "CompressionType")
	assertEqual(t, chunk.CompressionName(), "no compression", "CompressionName")

	var longNameBytes = make([]byte, 255)

	for i := 0; i < len(longNameBytes); i++ {
		longNameBytes[i] = 65
	}

	chunk = EncodeCOMMChunk(32, 2, 100, 200, 44100, CreateFourCC("NONE"), string(longNameBytes))

	assertEqual(t, chunk.CompressionName(), string(longNameBytes), "CompressionName")

	longNameBytes = make([]byte, 257)

	for i := 0; i < len(longNameBytes); i++ {
		longNameBytes[i] = 65
	}

	chunk = EncodeCOMMChunk(32, 2, 100, 200, 44100, CreateFourCC("NONE"), string(longNameBytes))

	assertEqual(t, len([]byte(chunk.CompressionName())), 255, "CompressionName length")
}
func TestDecodeCommChunk(t *testing.T) {
	chunk, err := DecodeCOMMChunk(nil)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, chunk, "chunk should be nil")

	data := make([]byte, HeaderSizeBytes-1)
	chunk, err = DecodeCOMMChunk(data)

	assertNotNil(t, err, "err should not be nil")
	assertNil(t, chunk, "chunk should be nil")

	data = make([]byte, HeaderSizeBytes)
	chunk, err = DecodeCOMMChunk(data)

	assertNotNil(t, err, "err should not be nil when DecodeCOMMChunk with not enough data")
	assertNotNil(t, chunk.Header, "header is nil")
	assertEqual(t, chunk.Channels(), 0, "Channels")
	assertEqual(t, chunk.SampleFrames(), 0, "SampleFrames")
	assertEqual(t, chunk.SampleSize(), 0, "SampleSize")
	assertEqual(t, chunk.CompressionType(), "", "CompressionType")
	assertEqual(t, chunk.CompressionName(), "", "CompressionName")

	data2 := make([]byte, 45)
	var channels uint16 = 1
	binary.BigEndian.PutUint16(data2[HeaderSizeBytes:HeaderSizeBytes+2], channels)
	var sampleFrames uint32 = 2
	binary.BigEndian.PutUint32(data2[HeaderSizeBytes+2:HeaderSizeBytes+6], sampleFrames)
	var samplesSize uint16 = 3
	binary.BigEndian.PutUint16(data2[HeaderSizeBytes+6:HeaderSizeBytes+8], samplesSize)
	var sampleRate [10]byte = [10]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	copy(data2[HeaderSizeBytes+8:], sampleRate[:])
	b := FourCC{'N', 'O', 'N', 'E'}
	var compressionType uint32 = b.ToUint32()
	binary.BigEndian.PutUint32(data2[HeaderSizeBytes+18:HeaderSizeBytes+22], compressionType)
	nameSize := [1]byte{14}
	copy(data2[HeaderSizeBytes+22:], nameSize[:])
	cName := []byte("no compression")
	copy(data2[HeaderSizeBytes+23:], cName[:])

	chunk, _ = DecodeCOMMChunk(data2)

	assertEqual(t, chunk.Channels(), 1, "Channels")
	assertEqual(t, chunk.SampleFrames(), 2, "SampleFrames")
	assertEqual(t, chunk.SampleSize(), 3, "SampleSize")
	assertEqual(t, chunk.CompressionType(), "NONE", "CompressionType")
	assertEqual(t, chunk.CompressionName(), "no compression", "CompressionName")
}

func TestCommBytes(t *testing.T) {
	data := make([]byte, 8)
	expectedData := make([]byte, 30)

	chunk, _ := DecodeCOMMChunk(data)
	bytes := chunk.Bytes()

	assertEqual(t, len(bytes), len(expectedData), "bytes length after EncodeCOMMChunk")

	compressionName := createTestString(14)
	expectedData = make([]byte, 30+14+2)
	chunk = EncodeCOMMChunk(32, int16(chunk.Channels()), uint32(chunk.SampleFrames()), int16(chunk.SampleSize()), chunk.SampleRate(), CreateFourCC("NONE"), compressionName)
	bytes = chunk.Bytes()

	assertEqual(t, len(bytes), len(expectedData), "bytes length after EncodeCOMMChunk")

	compressionName = createTestString(255)
	expectedData = make([]byte, 30+256)
	chunk = EncodeCOMMChunk(32, int16(chunk.Channels()), uint32(chunk.SampleFrames()), int16(chunk.SampleSize()), chunk.SampleRate(), CreateFourCC("NONE"), compressionName)
	bytes = chunk.Bytes()

	assertEqual(t, len(bytes), len(expectedData), "bytes length after EncodeCOMMChunk")

	compressionName = createTestString(256)
	expectedData = make([]byte, 30+256)
	chunk = EncodeCOMMChunk(32, int16(chunk.Channels()), uint32(chunk.SampleFrames()), int16(chunk.SampleSize()), chunk.SampleRate(), CreateFourCC("NONE"), compressionName)
	bytes = chunk.Bytes()

	assertEqual(t, len(bytes), len(expectedData), "bytes length after EncodeCOMMChunk")
}

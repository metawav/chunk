package chunk

import (
	"encoding/binary"
	"testing"
)

func TestEncodeCommChunk(t *testing.T) {
	chunk := EncodeCOMMChunk(32, 2, 100, 200, 44100, CreateFourCC("NONE"), "no compression")

	if chunk.ID() != "COMM" {
		t.Errorf("ID is %s, want %s", chunk.ID(), "COMM")
	}

	if chunk.Size() != 32 {
		t.Errorf("size is %d, want %d", chunk.Size(), 32)
	}

	if chunk.Channels() != 2 {
		t.Errorf("channels is %d, want %d", chunk.Channels(), 2)
	}

	if chunk.SampleFrames() != 100 {
		t.Errorf("sample frames is %d, want %d", chunk.SampleFrames(), 100)
	}

	if chunk.SampleSize() != 200 {
		t.Errorf("sample size is %d, want %d", chunk.SampleSize(), 200)
	}

	if chunk.SampleRate() != 44100 {
		t.Errorf("sample rate is %d, want %d", chunk.SampleRate(), 44100)
	}

	if chunk.CompressionType() != "NONE" {
		t.Errorf("compression type is %s, want %s", chunk.CompressionType(), "NONE")
	}

	if chunk.CompressionName() != "no compression" {
		t.Errorf("compression name is %s, want %s", chunk.CompressionName(), "no compression")
	}

	var longNameBytes = make([]byte, 255)

	for i := 0; i < len(longNameBytes); i++ {
		longNameBytes[i] = 65
	}

	chunk = EncodeCOMMChunk(32, 2, 100, 200, 44100, CreateFourCC("NONE"), string(longNameBytes))

	if chunk.CompressionName() != string(longNameBytes) {
		t.Errorf("compression name is %s, want %s", chunk.CompressionName(), string(longNameBytes))
	}

	longNameBytes = make([]byte, 257)

	for i := 0; i < len(longNameBytes); i++ {
		longNameBytes[i] = 65
	}

	chunk = EncodeCOMMChunk(32, 2, 100, 200, 44100, CreateFourCC("NONE"), string(longNameBytes))

	if len([]byte(chunk.CompressionName())) != 255 {
		t.Errorf("compression name length is %d, want %d", len([]byte(chunk.CompressionName())), 256)
	}
}
func TestDecodeCommChunk(t *testing.T) {
	chunk, err := DecodeCOMMChunk(nil)

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if chunk != nil {
		t.Errorf("chunk should be be nil")
	}

	data := make([]byte, HeaderSizeBytes-1)
	chunk, err = DecodeCOMMChunk(data)

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if chunk != nil {
		t.Errorf("chunk should be be nil")
	}

	data = make([]byte, HeaderSizeBytes)
	chunk, err = DecodeCOMMChunk(data)

	if err != nil {
		t.Errorf("err should be nil")
	}

	if chunk.Header == nil {
		t.Errorf("header is nil")
	}

	if chunk.Channels() != 0 {
		t.Errorf("channels is %d, want %d", chunk.Channels(), 0)
	}

	if chunk.SampleFrames() != 0 {
		t.Errorf("sample frames is %d, want %d", chunk.SampleFrames(), 0)
	}

	if chunk.SampleSize() != 0 {
		t.Errorf("sample size is %d, want %d", chunk.SampleSize(), 0)
	}

	if chunk.CompressionType() != "" {
		t.Errorf("compression type length is %s, want %s", chunk.CompressionType(), "")
	}

	if chunk.CompressionName() != "" {
		t.Errorf("compression name is %s, want %s", chunk.CompressionName(), "")
	}

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

	if chunk.Channels() != 1 {
		t.Errorf("channels is %d, want %d", chunk.Channels(), 0)
	}

	if chunk.SampleFrames() != 2 {
		t.Errorf("sample frames is %d, want %d", chunk.SampleFrames(), 2)
	}

	if chunk.SampleSize() != 3 {
		t.Errorf("sample size is %d, want %d", chunk.SampleSize(), 3)
	}

	if chunk.CompressionType() != "NONE" {
		t.Errorf("compression type is %s, want %s", chunk.CompressionType(), "NONE")
	}

	if chunk.CompressionName() != "no compression" {
		t.Errorf("compression name is %s, want %s", chunk.CompressionName(), cName)
	}
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

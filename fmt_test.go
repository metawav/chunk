package chunk

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestEncodeFMThunk(t *testing.T) {
	chunk := EncodeFMTChunk([4]byte{'f', 'm', 't', ' '}, 14, 0, 1, 44100, 1000, 2)

	if chunk.ID() != "fmt " {
		t.Errorf("ID is %s, want %s", chunk.ID(), "fmt ")
	}

	if chunk.Size() != 14 {
		t.Errorf("size is %d, want %d", chunk.Size(), 14)
	}

	if chunk.Channels() != 1 {
		t.Errorf("channels is %d, want %d", chunk.Channels(), 1)
	}

	if chunk.SamplesPerSec() != 44100 {
		t.Errorf("samples per sec is %d, want %d", chunk.SamplesPerSec(), 44100)
	}

	if chunk.BytesPerSec() != 1000 {
		t.Errorf("bytes per sec is %d, want %d", chunk.BytesPerSec(), 1000)
	}

	if chunk.BlockAlign() != 2 {
		t.Errorf("block align is %d, want %d", chunk.BlockAlign(), 2)
	}
}

func TestDecodeFMTChunk(t *testing.T) {
	chunk, err := DecodeFMTChunk(nil)

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if chunk != nil {
		t.Errorf("chunk should be be nil")
	}

	data := make([]byte, HeaderSizeBytes-1)
	chunk, err = DecodeFMTChunk(data)

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if chunk != nil {
		t.Errorf("chunk should be be nil")
	}

	data = make([]byte, HeaderSizeBytes)
	chunk, err = DecodeFMTChunk(data)

	if err != nil {
		t.Errorf("err should be nil")
	}

	if chunk.Header == nil {
		t.Errorf("header is nil")
	}

	if chunk.Format() != 0 {
		t.Errorf("format is %d, want %d", chunk.Format(), 0)
	}

	if chunk.Channels() != 0 {
		t.Errorf("channels is %d, want %d", chunk.Channels(), 0)
	}

	if chunk.SamplesPerSec() != 0 {
		t.Errorf("samples per sec is %d, want %d", chunk.SamplesPerSec(), 0)
	}

	if chunk.BytesPerSec() != 0 {
		t.Errorf("bytes per sec is %d, want %d", chunk.BytesPerSec(), 0)
	}

	if chunk.BlockAlign() != 0 {
		t.Errorf("block align is %d, want %d", chunk.BlockAlign(), 0)
	}

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

	if chunk.Format() != int(format) {
		t.Errorf("format is %d, want %d", chunk.Format(), int(format))
	}

	if chunk.Channels() != int(channels) {
		t.Errorf("format is %d, want %d", chunk.Channels(), int(channels))
	}

	if chunk.SamplesPerSec() != int(samplesPerSec) {
		t.Errorf("samples per sec is %d, want %d", chunk.SamplesPerSec(), int(samplesPerSec))
	}

	if chunk.BytesPerSec() != int(bytesPerSec) {
		t.Errorf("bytes per sec is %d, want %d", chunk.BytesPerSec(), int(bytesPerSec))
	}

	if chunk.BlockAlign() != int(blockAlign) {
		t.Errorf("block align is %d, want %d", chunk.BlockAlign(), int(blockAlign))
	}
}

func TestFMTBytes(t *testing.T) {
	data := make([]byte, 22)
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

	if bytes.Compare(chunk.Bytes(), data) != 0 {
		t.Errorf("bytes is %v, want %v", chunk.Bytes(), data)
	}
}

func TestEncodePCMFormathunk(t *testing.T) {
	chunk := EncodePCMFormatChunk([4]byte{'f', 'm', 't', ' '}, 14, 0, 1, 44100, 1000, 2, 24)

	if chunk.ID() != "fmt " {
		t.Errorf("ID is %s, want %s", chunk.ID(), "fmt ")
	}

	if chunk.Size() != 14 {
		t.Errorf("size is %d, want %d", chunk.Size(), 14)
	}

	if chunk.Channels() != 1 {
		t.Errorf("channels is %d, want %d", chunk.Channels(), 1)
	}

	if chunk.SamplesPerSec() != 44100 {
		t.Errorf("samples per sec is %d, want %d", chunk.SamplesPerSec(), 44100)
	}

	if chunk.BytesPerSec() != 1000 {
		t.Errorf("bytes per sec is %d, want %d", chunk.BytesPerSec(), 1000)
	}

	if chunk.BlockAlign() != 2 {
		t.Errorf("block align is %d, want %d", chunk.BlockAlign(), 2)
	}

	if chunk.BitsPerSample() != 24 {
		t.Errorf("bits per sample is %d, want %d", chunk.BitsPerSample(), 24)
	}
}

func TestDecodePCMFormatChunk(t *testing.T) {
	chunk, err := DecodePCMFormatChunk(nil)

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if chunk != nil {
		t.Errorf("chunk should be be nil")
	}

	dataMinSize := 22
	data := make([]byte, dataMinSize-1)
	chunk, err = DecodePCMFormatChunk(data)

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if chunk != nil {
		t.Errorf("chunk should be be nil")
	}

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

	if err == nil {
		t.Errorf("err should not be nil")
	}

	if chunk != nil {
		t.Errorf("chunk should be be nil")
	}

	data = make([]byte, 2)
	var bitsPerSample uint16 = 24
	binary.LittleEndian.PutUint16(data[0:2], bitsPerSample)
	data = append(fmtData, data...)
	chunk, err = DecodePCMFormatChunk(data)

	if err != nil {
		t.Errorf("err should be nil")
	}

	if chunk.BitsPerSample() != int(bitsPerSample) {
		t.Errorf("bits per sample is %d, want %d", chunk.BitsPerSample(), int(bitsPerSample))
	}
}

func TestPCMFormatBytes(t *testing.T) {
	data := make([]byte, 24)
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

	if bytes.Compare(chunk.Bytes(), data) != 0 {
		t.Errorf("bytes is %v, want %v", chunk.Bytes(), data)
	}
}

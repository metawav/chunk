package chunk

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

// FMT is RIFF format chunk 'fmt ' describing sampled sound in data chunk 'data'.
type FMT struct {
	*Header
	format        uint16
	channels      uint16
	samplesPerSec uint32
	bytesPerSec   uint32
	blockAlign    uint16
}

// Format
func (fc *FMT) Format() int {
	return int(fc.format)
}

// Channels
func (fc *FMT) Channels() int {
	return int(fc.channels)
}

// SamplesPerSec
func (fc *FMT) SamplesPerSec() int {
	return int(fc.samplesPerSec)
}

// BytesPerSec
func (fc *FMT) BytesPerSec() int {
	return int(fc.bytesPerSec)
}

// BlockAlign
func (fc *FMT) BlockAlign() int {
	return int(fc.blockAlign)
}

// String returns string represensation of chunk.
func (fc *FMT) String() string {
	return fmt.Sprintf("Format: %d\nChannels: %d\nSample rate: %d\nByte rate: %d\nBlock align: %d", fc.Format(), fc.Channels(), fc.SamplesPerSec(), fc.BytesPerSec(), fc.BlockAlign())
}

// Bytes converts FMT to byte array. A new Header with id 'fmt ' is created.
//
// Header size is set to real data size. An amount of 22 bytes is returned.
// chunk header - 8 bytes
// format - 2 bytes
// channels - 2 bytes
// samples per sec - 4 bytes
// bytes per sec - 4 bytes
// block align - 2 bytes
//
// A padding byte is added if size is odd. This optional byte is not reflected in size.
func (fc *FMT) Bytes() []byte {
	byteOrder := binary.LittleEndian
	data := make([]byte, 14)
	byteOrder.PutUint16(data[0:2], fc.format)
	byteOrder.PutUint16(data[2:4], fc.channels)
	byteOrder.PutUint32(data[4:8], fc.samplesPerSec)
	byteOrder.PutUint32(data[8:12], fc.bytesPerSec)
	byteOrder.PutUint16(data[12:14], fc.blockAlign)
	dataSize := len(data)
	header := EncodeChunkHeader(CreateFourCC(FMTID), uint32(dataSize), byteOrder)
	bytes := append(header.Bytes(), data...)

	return pad(bytes)
}

// EncodeFMTChunk returns encoded chunk 'fmt ' by provided parameters.
func EncodeFMTChunk(size uint32, format uint16, channels uint16, samplesPerSec uint32, bytesPerSec uint32, blockAlign uint16) *FMT {
	id := CreateFourCC(FMTID)
	header := EncodeChunkHeader(id, size, binary.LittleEndian)

	return &FMT{Header: header, format: format, channels: channels, samplesPerSec: samplesPerSec, bytesPerSec: bytesPerSec, blockAlign: blockAlign}
}

// DecodeFMTChunk decodes provided byte array to FMT.
//
// Array content should be:
// chunk header - 8 bytes (min. requirement for successful decoding)
// format - 2 bytes
// channels - 2 bytes
// samples per sec - 4 bytes
// bytes per sec - 4 bytes
// block align - 2 bytes
func DecodeFMTChunk(data []byte) (*FMT, error) {
	if len(data) < int(HeaderSizeBytes) {
		msg := fmt.Sprintf("data slice requires a minimim lenght of %d", HeaderSizeBytes)
		return nil, errors.New(msg)
	}

	fc := &FMT{}
	byteOrder := binary.LittleEndian
	fc.Header = decodeChunkHeader(data[:HeaderSizeBytes], 0, byteOrder)
	buf := bytes.NewReader(data[HeaderSizeBytes:])
	fields := []interface{}{&fc.format, &fc.channels, &fc.samplesPerSec, &fc.bytesPerSec, &fc.blockAlign}

	for _, f := range fields {
		err := binary.Read(buf, byteOrder, f)

		if err != nil {
			return fc, err
		}
	}

	return fc, nil
}

// PCMFormat is RIFF format chunk 'fmt ' describing sampled sound in data chunk 'data'.
type PCMFormat struct {
	*FMT
	bitsPerSample uint16
}

// BitsPerSample
func (pfc *PCMFormat) BitsPerSample() int {
	return int(pfc.bitsPerSample)
}

// String returns string represensation of chunk.
func (pfc *PCMFormat) String() string {
	return fmt.Sprintf("%s\nBits per sample: %d", pfc.FMT, pfc.BitsPerSample())
}

// Bytes converts PCMFormat to byte array. A new Header with id 'fmt ' is created.
//
// Header size is set to real data size. An amount of 24 bytes is returned.
// FMT - 22 bytes (see FMT)
// bits per sample - 2 bytes
//
// A padding byte is added if size is odd. This optional byte is not reflected in size.
func (pfc *PCMFormat) Bytes() []byte {
	byteOrder := binary.LittleEndian
	data := make([]byte, 2)
	byteOrder.PutUint16(data[0:2], pfc.bitsPerSample)
	fmtBytes := pfc.FMT.Bytes()
	data = append(fmtBytes[HeaderSizeBytes:], data...)
	dataSize := len(data)
	header := EncodeChunkHeader(CreateFourCC(FMTID), uint32(dataSize), byteOrder)
	bytes := append(header.Bytes(), data...)

	return pad(bytes)
}

// EncodePCMFormatChunk returns encoded chunk 'fmt ' by provided parameters.
func EncodePCMFormatChunk(size uint32, format uint16, channels uint16, samplesPerSec uint32, bytesPerSec uint32, blockAlign uint16, bitPerSample uint16) *PCMFormat {
	fmt := EncodeFMTChunk(size, format, channels, samplesPerSec, bytesPerSec, blockAlign)

	return &PCMFormat{FMT: fmt, bitsPerSample: bitPerSample}
}

// DecodePCMFormatChunk decodes provided byte array to PCMFormat.
//
// Array content should be:
// FMT - 22 bytes (see FMT)
// bits per sample - 2 bytes
func DecodePCMFormatChunk(data []byte) (*PCMFormat, error) {
	fc, err := DecodeFMTChunk(data)

	if err != nil {
		return nil, err
	}

	if len(data) < len(fc.Bytes()) {
		msg := fmt.Sprintf("data slice requires a minimim lenght of %d", len(fc.Bytes()))
		return nil, errors.New(msg)
	}

	pfc := &PCMFormat{FMT: fc}
	buf := bytes.NewReader(data[len(fc.Bytes()):])
	err = binary.Read(buf, binary.LittleEndian, &pfc.bitsPerSample)

	return pfc, err
}

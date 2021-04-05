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
	return fmt.Sprintf("Format: %d\nChannels: %d\nSample rate: %d\nByte rate: %d\nBytes per sample: %d", fc.Format(), fc.Channels(), fc.SamplesPerSec(), fc.BytesPerSec(), fc.BlockAlign())
}

// Bytes converts FMT to byte array.
// An amount of 22 bytes is returned.
func (fc *FMT) Bytes() []byte {
	bytes := fc.Header.Bytes()

	data := make([]byte, 14)
	byteOrder := binary.LittleEndian

	byteOrder.PutUint16(data[0:2], fc.format)
	byteOrder.PutUint16(data[2:4], fc.channels)
	byteOrder.PutUint32(data[4:8], fc.samplesPerSec)
	byteOrder.PutUint32(data[8:12], fc.bytesPerSec)
	byteOrder.PutUint16(data[12:14], fc.blockAlign)

	bytes = append(bytes, data...)

	return bytes
}

// EncodeFMTChunk returns encoded chunk 'fmt ' by provided parameters.
func EncodeFMTChunk(size uint32, format uint16, channels uint16, samplesPerSec uint32, bytesPerSec uint32, blockAlign uint16) *FMT {
	id := CreateFourCC(FMTID)
	header := EncodeChunkHeader(id, size, binary.LittleEndian)

	return &FMT{Header: header, format: format, channels: channels, samplesPerSec: samplesPerSec, bytesPerSec: bytesPerSec, blockAlign: blockAlign}
}

// DecodeFMTChunk decodes provided byte array to FMT.
// Array content is:
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
			err = handleError(err)

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

// Bytes converts PCMFormat to byte array.
// An amount of 24 bytes is returned.
func (pfc *PCMFormat) Bytes() []byte {
	bytes := pfc.FMT.Bytes()

	data := make([]byte, 2)
	byteOrder := binary.LittleEndian

	byteOrder.PutUint16(data[0:2], pfc.bitsPerSample)

	bytes = append(bytes, data...)

	return bytes
}

// EncodePCMFormatChunk returns encoded chunk 'fmt ' by provided parameters.
func EncodePCMFormatChunk(size uint32, format uint16, channels uint16, samplesPerSec uint32, bytesPerSec uint32, blockAlign uint16, bitPerSample uint16) *PCMFormat {
	fmt := EncodeFMTChunk(size, format, channels, samplesPerSec, bytesPerSec, blockAlign)

	return &PCMFormat{FMT: fmt, bitsPerSample: bitPerSample}
}

// DecodePCMFormatChunk decodes provided byte array to PCMFormat.
// Array content is:
// FMT - 22 bytes (see FMT)
// bits per sample - 2 bytes
func DecodePCMFormatChunk(data []byte) (*PCMFormat, error) {
	fc, err := DecodeFMTChunk(data)

	if err != nil {
		return nil, err
	}

	pfc := &PCMFormat{FMT: fc}

	if len(data) < len(fc.Bytes()) {
		msg := fmt.Sprintf("data slice requires a minimim lenght of %d", len(fc.Bytes()))
		return nil, errors.New(msg)
	}

	buf := bytes.NewReader(data[len(fc.Bytes()):])
	err = binary.Read(buf, binary.LittleEndian, &pfc.bitsPerSample)

	if err != nil {
		return nil, err
	}

	return pfc, nil
}

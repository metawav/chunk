package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

// Format
type Format struct {
	*Header
	format        uint16
	channels      uint16
	samplesPerSec uint32
	bytesPerSec   uint32
	blockAlign    uint16
}

func (fc *Format) Format() string {
	return strconv.Itoa(int(fc.format))
}

func (fc *Format) Channels() int {
	return int(fc.channels)
}

func (fc *Format) SamplesPerSec() int {
	return int(fc.samplesPerSec)
}

func (fc *Format) BytesPerSec() int {
	return int(fc.bytesPerSec)
}

func (fc *Format) BlockAlign() int {
	return int(fc.blockAlign)
}

func (fc *Format) String() string {
	return fmt.Sprintf("Format: %s\nChannels: %d\nSample rate: %d\nByte rate: %d\nBytes per sample: %d", fc.Format(), fc.Channels(), fc.SamplesPerSec(), fc.BytesPerSec(), fc.BlockAlign())
}

// Bytes converts Format to  byte array.
func (fc *Format) Bytes() []byte {
	bytes := fc.Header.Bytes()

	//todo: implement

	return bytes
}

// DecodeFormatChunk
func DecodeFormatChunk(data []byte) (*Format, error) {
	fc := &Format{}
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

// PCMFormat
type PCMFormat struct {
	*Format
	bitsPerSample uint16
}

// BitsPerSample
func (pfc *PCMFormat) BitsPerSample() int {
	return int(pfc.bitsPerSample)
}

func (pfc *PCMFormat) String() string {
	return fmt.Sprintf("%s\nBits per sample: %d", pfc.Format, pfc.BitsPerSample())
}

func (pfc *PCMFormat) Bytes() []byte {
	bytes := pfc.Format.Bytes()

	//todo: implement

	return bytes
}

// DecodePCMFormatChunk
func DecodePCMFormatChunk(data []byte) (*PCMFormat, error) {
	fc, err := DecodeFormatChunk(data)

	if err != nil {
		return nil, err
	}

	pfc := &PCMFormat{Format: fc}
	//todo: replace 22 by len(fc.Bytes()
	buf := bytes.NewReader(data[22:])
	err = binary.Read(buf, binary.LittleEndian, &pfc.bitsPerSample)

	if err != nil {
		return nil, err
	}

	return pfc, nil
}

package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

// Format
type Format struct {
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

func readFormatChunk(data []byte) (*Format, error) {
	fc := &Format{}
	buf := bytes.NewReader(data[HeaderSizeBytes:])
	byteOrder := binary.LittleEndian
	err := binary.Read(buf, byteOrder, &fc.format)
	err = binary.Read(buf, byteOrder, &fc.channels)
	err = binary.Read(buf, byteOrder, &fc.samplesPerSec)
	err = binary.Read(buf, byteOrder, &fc.bytesPerSec)
	err = binary.Read(buf, byteOrder, &fc.blockAlign)

	if err != nil {
		return nil, err
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

// ReadPCMFormatChunk
func ReadPCMFormatChunk(data []byte) (*PCMFormat, error) {
	fc, err := readFormatChunk(data)

	if err != nil {
		return nil, err
	}

	pfc := &PCMFormat{Format: fc}
	buf := bytes.NewReader(data[22:])
	err = binary.Read(buf, binary.LittleEndian, &pfc.bitsPerSample)

	if err != nil {
		return nil, err
	}

	return pfc, nil
}

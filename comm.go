package chunk

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/metawav/chunk/internal"
)

// COMM is AIFF / AIFF-C common chunk 'COMM' describing sampled sound in sound chunk 'SSND'.
type COMM struct {
	*Header
	numChannels     int16
	numSampleFrames uint32
	sampleSize      int16
	sampleRate      [10]byte
	compressionType FourCC
	compressionName []byte
}

// Channels
func (c *COMM) Channels() int {
	return int(c.numChannels)
}

// SampleFrames
func (c *COMM) SampleFrames() int {
	return int(c.numSampleFrames)
}

// SampleSize
func (c *COMM) SampleSize() int {
	return int(c.sampleSize)
}

// SampleRate
func (c *COMM) SampleRate() uint {
	return internal.IEEEFloatToInt(c.sampleRate)
}

// CompressionType
func (c *COMM) CompressionType() string {
	return trim(c.compressionType[:])
}

// CompressionName
func (c *COMM) CompressionName() string {
	return string(c.compressionName[:])
}

// String returns string represensation of chunk.
func (c *COMM) String() string {
	return fmt.Sprintf("%s Channels: %d Sample frames: %d Sample size: %d Sample rate: %d Compression type: %s Compression name: %s", c.Header, c.Channels(), c.SampleFrames(), c.SampleSize(), c.SampleRate(), c.CompressionType(), c.CompressionName())
}

// Bytes converts COMM to byte array. A new Header with id 'COMM' is created.
//
// Header size is set to real data size. A minimum of 30 bytes is returned. In case a compression name
// is set a minimum of 2 and a maximum of 256 additional bytes is returned.
//
// A padding byte is added if size is odd. This optional byte is not reflected in size.
func (c *COMM) Bytes() []byte {
	byteOrder := binary.BigEndian
	data := make([]byte, 22)
	byteOrder.PutUint16(data[0:2], uint16(c.numChannels))
	byteOrder.PutUint32(data[2:6], c.numSampleFrames)
	byteOrder.PutUint16(data[6:8], uint16(c.sampleSize))
	copy(data[8:18], c.sampleRate[:])
	copy(data[18:22], []byte(c.CompressionType()))

	if c.compressionName != nil {
		compressionNameSize := len(c.compressionName)
		compressionName := c.compressionName

		if compressionNameSize > 255 {
			compressionName = compressionName[:255]
		}

		data = append(data, byte(len(compressionName)))
		data = append(data, c.compressionName[:]...)

		if len(compressionName)%2 == 0 {
			data = append(data, 0)
		}
	}

	dataSize := len(data)
	header := EncodeChunkHeader(CreateFourCC(COMMID), uint32(dataSize), byteOrder)
	bytes := append(header.Bytes(), data...)

	return pad(bytes)
}

// EncodeCOMMChunk returns encoded chunk 'COMM' by provided parameters.
func EncodeCOMMChunk(size uint32, numChannels int16, numSamplesPerFrame uint32, sampleSize int16, sampleRate uint, compressionType FourCC, compressionName string) *COMM {
	id := CreateFourCC(COMMID)
	header := EncodeChunkHeader(id, size, binary.LittleEndian)

	sRate := internal.IntToIEEE(sampleRate)
	compressionNameSize := len(compressionName)

	if compressionNameSize > 255 {
		compressionName = compressionName[:255]
	}

	return &COMM{Header: header, numChannels: numChannels, numSampleFrames: numSamplesPerFrame, sampleSize: sampleSize, sampleRate: sRate, compressionType: compressionType, compressionName: []byte(compressionName)}
}

// DecodeCOMMChunk decodes provided byte array to COMM.
//
// Array content should be:
// chunk header - 8 bytes (min. requirement for successful decoding)
// num channels - 2 bytes
// num sample frames - 4 bytes
// sample size - 2 bytes
// sample rate - 10 bytes 80 bit encoded float with extended precision (IEEE 754)
// compression type - 4 bytes
// compression name - max. 256 bytes (1 byte size, max. 255 byte name, max. 1 byte padding)
func DecodeCOMMChunk(data []byte) (*COMM, error) {
	if len(data) < int(HeaderSizeBytes) {
		msg := fmt.Sprintf("data slice requires a minimim lenght of %d", HeaderSizeBytes)
		return nil, errors.New(msg)
	}

	c := &COMM{}

	byteOrder := binary.BigEndian
	c.Header = decodeChunkHeader(data[:HeaderSizeBytes], 0, byteOrder)
	var compressionNameSize uint8

	buf := bytes.NewReader(data[HeaderSizeBytes:])
	fields := []interface{}{&c.numChannels, &c.numSampleFrames, &c.sampleSize, &c.sampleRate, &c.compressionType, &compressionNameSize}

	for _, f := range fields {
		err := binary.Read(buf, byteOrder, f)

		if err != nil {
			return c, err
		}
	}

	compressionName := make([]byte, compressionNameSize)
	err := binary.Read(buf, byteOrder, &compressionName)

	if err != nil {
		return c, err
	}

	c.compressionName = append(c.compressionName, compressionName...)

	return c, nil
}

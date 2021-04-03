package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pmoule/wav/internal"
)

// COMM
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
func (c *COMM) SampleRate() int {
	return internal.IEEEFloatToInt(c.sampleRate)
}

// COmpressionType
func (c *COMM) CompressionType() string {
	return string(c.compressionType[:])
}

// CompressionName
func (c *COMM) CompressionName() string {
	return string(c.compressionName[:])
}

// String returns string represensation of chunk.
func (c *COMM) String() string {
	return fmt.Sprintf("%s Channels: %d Sample frames: %d Sample size: %d Sample rate: %d Compression type: %s Compression name: %s", c.Header, c.Channels(), c.SampleFrames(), c.SampleSize(), c.SampleRate(), c.CompressionType(), c.CompressionName())
}

func (c *COMM) Bytes() []byte {
	bytes := c.Header.Bytes()

	//todo: implement

	return bytes
}

// DecodeCOMMChunk
func DecodeCOMMChunk(data []byte) (*COMM, error) {
	c := &COMM{}

	byteOrder := binary.BigEndian
	c.Header = decodeChunkHeader(data[:HeaderSizeBytes], 0, byteOrder)
	var compressionNameSize uint8

	buf := bytes.NewReader(data[HeaderSizeBytes:])
	fields := []interface{}{&c.numChannels, &c.numSampleFrames, &c.sampleSize, &c.sampleRate, &c.compressionType, &compressionNameSize}

	for _, f := range fields {
		err := binary.Read(buf, byteOrder, f)

		if err != nil {
			err = handleError(err)

			return c, err
		}
	}

	compressionName := make([]byte, compressionNameSize)
	err := binary.Read(buf, byteOrder, &compressionName)

	if err != nil {
		err = handleError(err)

		return c, err
	}

	c.compressionName = append(c.compressionName, compressionName...)

	return c, nil
}

func handleError(err error) error {
	if err == io.EOF {
		return nil
	}

	return err
}

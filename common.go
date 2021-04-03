package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pmoule/wav/internal"
)

type Common struct {
	*Header
	numChannels     int16
	numSampleFrames uint32
	sampleSize      int16
	sampleRate      [10]byte
	compressionType FourCC
	compressionName []byte
}

func (c *Common) Channels() int {
	return int(c.numChannels)
}

func (c *Common) SampleFrames() int {
	return int(c.numSampleFrames)
}

func (c *Common) SampleSize() int {
	return int(c.sampleSize)
}

func (c *Common) SampleRate() int {
	return internal.IEEEFloatToInt(c.sampleRate)
}

func (c *Common) CompressionType() string {
	return string(c.compressionType[:])
}

func (c *Common) CompressionName() string {
	return string(c.compressionName[:])
}

func (c *Common) String() string {
	return fmt.Sprintf("%s Channels: %d Sample frames: %d Sample size: %d Sample rate: %d Compression type: %s Compression name: %s", c.Header, c.Channels(), c.SampleFrames(), c.SampleSize(), c.SampleRate(), c.CompressionType(), c.CompressionName())
}

func (c *Common) Bytes() []byte {
	bytes := c.Header.Bytes()

	//todo: implement

	return bytes
}

// DecodeCommonChunk
func DecodeCommonChunk(data []byte) (*Common, error) {
	c := &Common{}

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

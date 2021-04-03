package chunk

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/pmoule/wav/internal"
)

type Common struct {
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

// ReadCommonChunk
func ReadCommonChunk(data []byte) (*Common, error) {
	c := &Common{}

	var id FourCC
	var size int32
	var compressionNameSize uint8
	compressionName := make([]byte, compressionNameSize)

	buf := bytes.NewReader(data[:])
	byteOrder := binary.BigEndian

	binary.Read(buf, byteOrder, &id)
	binary.Read(buf, byteOrder, &size)
	binary.Read(buf, byteOrder, &c.numChannels)
	binary.Read(buf, byteOrder, &c.numSampleFrames)
	binary.Read(buf, byteOrder, &c.sampleSize)
	binary.Read(buf, byteOrder, &c.sampleRate)
	err := binary.Read(buf, byteOrder, &c.compressionType)

	if err != nil {
		err = handleError(err)

		return c, err
	}

	err = binary.Read(buf, byteOrder, &compressionNameSize)
	err = binary.Read(buf, byteOrder, &compressionName)
	c.compressionName = append(c.compressionName, compressionName...)

	return c, nil
}

func handleError(err error) error {
	if err == io.EOF {
		return nil
	}

	return err
}

func (c *Common) String() string {
	return fmt.Sprintf("Channels: %d Sample frames: %d Sample size: %d Sample rate: %d Compression type: %s Compression name: %s", c.Channels(), c.SampleFrames(), c.SampleSize(), c.SampleRate(), c.CompressionType(), c.CompressionName())
}

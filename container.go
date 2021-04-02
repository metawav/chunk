package chunk

import (
	"encoding/binary"
	"io"
)

// Container describes a chunked data structure by name, container chunk header, contained chunk headers and byte order.
type Container struct {
	Name      string
	Header    *ContainerHeader
	Headers   []*Header
	ByteOrder binary.ByteOrder
}

// FindHeader returns headers with provided ID.
func (c *Container) FindHeaders(id string) []*Header {
	var headers []*Header

	for _, header := range c.Headers {
		if header.ID() == id {
			headers = append(headers, header)
		}
	}

	return headers
}

// ReadAiff reads a AIFF / AIFF-C file from provided byte stream and assigns provided name.
func ReadAiff(name string, reader io.ReadSeeker) (*Container, error) {
	return read(name, reader, binary.BigEndian)
}

// ReadRiff reads a RIFF file from provided byte stream and assigns provided name.
func ReadRiff(name string, reader io.ReadSeeker) (*Container, error) {
	return read(name, reader, binary.LittleEndian)
}

// Read
func read(name string, reader io.ReadSeeker, byteOrder binary.ByteOrder) (*Container, error) {
	containerHeader, err := readContainerHeader(reader, byteOrder)

	if err != nil {
		return nil, err
	}

	headers := readChunkHeaders(reader, ContainerHeaderSizeBytes, byteOrder)
	container := &Container{Name: name, Header: containerHeader, Headers: headers, ByteOrder: byteOrder}

	return container, nil
}

func readChunkHeaders(reader io.ReadSeeker, offset uint32, byteOrder binary.ByteOrder) []*Header {
	var headers []*Header

	for {
		var chunkHeaderBytes [HeaderSizeBytes]byte
		_, err := io.ReadFull(reader, chunkHeaderBytes[:])

		if err != nil {
			break
		}

		chunkHeader := DecodeChunkHeader(chunkHeaderBytes, offset, byteOrder)
		headers = append(headers, chunkHeader)
		offset += chunkHeader.FullSize()

		seek := int64(chunkHeader.Size())

		// For compatibility with EA IFF (Electronic Arts Interchange File Format)
		// chunks must be even sized and always start at an even position.
		if chunkHeader.HasPadding() {
			seek++
		}

		reader.Seek(seek, io.SeekCurrent)
	}

	return headers
}

func readContainerHeader(reader io.ReadSeeker, byteOrder binary.ByteOrder) (*ContainerHeader, error) {
	var headerBytes [ContainerHeaderSizeBytes]byte
	_, err := io.ReadFull(reader, headerBytes[:])

	if err != nil {
		return nil, err
	}

	header := DecodeContainerHeader(headerBytes, byteOrder)

	return header, nil
}

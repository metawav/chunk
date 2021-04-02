package internal

import (
	"encoding/binary"
	"io"

	"github.com/pmoule/wav/chunk"
)

// Read
func Read(name string, reader io.ReadSeeker, byteOrder binary.ByteOrder) (*chunk.Container, error) {
	containerHeader, err := readContainerHeader(reader, byteOrder)

	if err != nil {
		return nil, err
	}

	headers := readChunkHeaders(reader, chunk.ContainerHeaderSizeBytes, byteOrder)
	container := &chunk.Container{Name: name, Header: containerHeader, Headers: headers, ByteOrder: byteOrder}

	return container, nil
}

func readChunkHeaders(reader io.ReadSeeker, offset uint32, byteOrder binary.ByteOrder) []*chunk.Header {
	var headers []*chunk.Header

	for {
		var chunkHeaderBytes [chunk.HeaderSizeBytes]byte
		_, err := io.ReadFull(reader, chunkHeaderBytes[:])

		if err != nil {
			break
		}

		chunkHeader := chunk.DecodeChunkHeader(chunkHeaderBytes, offset, byteOrder)
		headers = append(headers, chunkHeader)
		offset += chunkHeader.FullSize()

		// For compatibility with EA IFF (Electronic Arts Interchange File Format)
		// chunks must be even sized and always start at an even position.
		if offset%2 != 0 {
			offset++
			reader.Seek(1, io.SeekCurrent)
		}

		reader.Seek(int64(chunkHeader.Size()), io.SeekCurrent)
	}

	return headers
}

func readContainerHeader(reader io.ReadSeeker, byteOrder binary.ByteOrder) (*chunk.ContainerHeader, error) {
	var headerBytes [chunk.ContainerHeaderSizeBytes]byte
	_, err := io.ReadFull(reader, headerBytes[:])

	if err != nil {
		return nil, err
	}

	riffHeader := chunk.DecodeContainerHeader(headerBytes, byteOrder)

	return riffHeader, nil
}

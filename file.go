package wav

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sort"
)

type RiffFile struct {
	Name    string
	Header  *RiffHeader
	Headers []*Header
}

func CreateRiffFile(fileName string, reader io.ReadSeeker) (*RiffFile, error) {
	riffHeader, err := readRiffHeader(reader)

	if err != nil {
		return nil, err
	}

	headers := readChunkHeaders(reader, RiffHeaderSizeBytes)
	riffFile := &RiffFile{Name: fileName, Header: riffHeader, Headers: headers}

	return riffFile, nil
}

func readChunkHeaders(reader io.ReadSeeker, offset uint32) []*Header {
	var headers []*Header

	for {
		var chunkHeaderBytes [HeaderSizeBytes]byte
		_, err := reader.Read(chunkHeaderBytes[:])

		if err != nil {
			break
		}

		chunkHeader := DecodeChunkHeader(chunkHeaderBytes, offset)
		headers = append(headers, chunkHeader)
		offset += chunkHeader.FullSize()

		// For compatibility with EA IFF (Electronic Arts Interchange File Format)
		// chunks must be even sized and always start at an even position.
		if offset%2 != 0 {
			offset++
		}

		reader.Seek(int64(chunkHeader.Size()), io.SeekCurrent)
	}

	return headers
}

func readRiffHeader(reader io.ReadSeeker) (*RiffHeader, error) {
	var headerBytes [RiffHeaderSizeBytes]byte
	_, err := reader.Read(headerBytes[:])

	if err != nil {
		return nil, err
	}

	riffHeader := DecodeRiffHeader(headerBytes)

	return riffHeader, nil
}

// GetHeaderByID returns header with provided ID or nil if not contained in RIFF.
func (rf *RiffFile) GetHeaderByID(headerID string) *Header {
	for _, header := range rf.Headers {
		if header.ID() == headerID {
			return header
		}
	}

	return nil
}

func (rf *RiffFile) DeleteChunk(headerID string, reader io.ReaderAt, writer io.WriterAt) (uint32, error) {
	header := rf.GetHeaderByID(headerID)

	if header == nil {
		msg := fmt.Sprintf("chunk not found: %s", headerID)
		return 0, errors.New(msg)
	}

	headers := rf.Headers
	sort.Sort(SortHeadersByStartPosAsc(headers))

	offset := header.StartPos()

	for i := 0; i < len(headers); i++ {
		if headers[i].StartPos() > header.StartPos() {
			headerBytes := make([]byte, headers[i].FullSize())
			//todo: ensure reading complete data
			n, err := reader.ReadAt(headerBytes, int64(headers[i].StartPos()))

			if err != nil {
				return 0, err
			}

			_, err = writer.WriteAt(headerBytes, int64(offset))

			if err != nil {
				return 0, err
			}

			offset += uint32(n)
		}
	}

	riffSize := rf.Header.Size() - header.FullSize()
	err := rf.UpdateSize(riffSize, writer)

	if err != nil {
		return 0, err
	}

	fileSize := rf.Header.FullSize() - header.FullSize()

	return fileSize, nil
}

// AddChunk
func (rf *RiffFile) AddChunk(reader io.Reader, writer io.WriterAt, bufferSize int) error {
	if bufferSize <= 0 {
		bufferSize = 1024
	}

	// start writing chunk to end of file
	offset := int64(rf.Header.FullSize())
	var chunkSize uint32 = 0

	for {
		b := make([]byte, bufferSize)
		// todo: how to be sure it's a valid header?
		n, err := reader.Read(b)

		if n > 0 {
			b = b[:n]
			_, err := writer.WriteAt(b, offset)

			if err != nil {
				return err
			}

			offset += int64(n)
			chunkSize += uint32(n)
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}
	}

	riffSize := rf.Header.size + chunkSize
	err := rf.UpdateSize(riffSize, writer)

	if err != nil {
		return err
	}

	return nil
}

// UpdateSize
func (rf *RiffFile) UpdateSize(size uint32, writer io.WriterAt) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, size)

	if err != nil {

		return err
	}

	b := buf.Bytes()
	writer.WriteAt(b, 4)

	return nil
}

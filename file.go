package wav

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sort"
)

// RiffFile describes a RIFF file by name, RIFF header and contained headers.
type RiffFile struct {
	Name    string
	Header  *RiffHeader
	Headers []*Header
}

// CreateRiffFile creates a RiffFile struct from provided byte stream and assigns provided name.
func CreateRiffFile(name string, reader io.ReadSeeker) (*RiffFile, error) {
	riffHeader, err := readRiffHeader(reader)

	if err != nil {
		return nil, err
	}

	headers := readChunkHeaders(reader, RiffHeaderSizeBytes)
	riffFile := &RiffFile{Name: name, Header: riffHeader, Headers: headers}

	return riffFile, nil
}

func readChunkHeaders(reader io.ReadSeeker, offset uint32) []*Header {
	var headers []*Header

	for {
		var chunkHeaderBytes [HeaderSizeBytes]byte
		_, err := io.ReadFull(reader, chunkHeaderBytes[:])

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
	_, err := io.ReadFull(reader, headerBytes[:])

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

// DeleteChunk
func (rf *RiffFile) DeleteChunk(headerID string, reader io.ReaderAt, writer io.WriterAt) (uint32, error) {
	header := rf.GetHeaderByID(headerID)

	if header == nil {
		msg := fmt.Sprintf("chunk not found: %s", headerID)
		return 0, errors.New(msg)
	}

	headers := rf.Headers
	sort.Sort(SortHeadersByStartPosAsc(headers))
	writeOffset := header.StartPos()

	for i := 0; i < len(headers); i++ {
		if headers[i].StartPos() > header.StartPos() {
			sectionReader := io.NewSectionReader(reader, int64(headers[i].StartPos()), int64(headers[i].FullSize()))
			n, err := moveChunk(writeOffset, headers[i].FullSize(), sectionReader, writer)

			if err != nil {
				return 0, err
			}

			writeOffset += uint32(n)
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

func moveChunk(writeOffset uint32, size uint32, reader io.Reader, writer io.WriterAt) (int, error) {
	bytes := make([]byte, size)
	n, err := io.ReadFull(reader, bytes[:])

	if err != nil {
		return 0, err
	}

	_, err = writer.WriteAt(bytes, int64(writeOffset))

	if err != nil {
		return 0, err
	}

	return n, nil
}

// AddChunk reads chnunk bytes from io.Reader and writes to io.Writer
func (rf *RiffFile) AddChunk(reader io.Reader, writer io.WriterAt, bufferSize int) error {
	if bufferSize <= 0 {
		bufferSize = 1024
	}

	// start writing chunk to end of file
	offset := int64(rf.Header.FullSize())
	var chunkSize uint32 = 0

	for {
		b := make([]byte, bufferSize)
		n, err := io.ReadFull(reader, b)

		if n > 0 {
			//todo: ensure writng all bytes from b
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

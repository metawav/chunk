package wav

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sort"
)

// File describes a chunked file by name, container header, contained headers and byte order.
type File struct {
	Name      string
	Header    *ContainerHeader
	Headers   []*Header
	ByteOrder binary.ByteOrder
}

// CreateRiffFile creates a RiffFile struct from provided byte stream and assigns provided name.
func CreateRiffFile(name string, reader io.ReadSeeker) (*File, error) {
	return createFile(name, reader, binary.LittleEndian)
}

// CreateAiffFile creates a AiffFile struct from provided byte stream and assigns provided name.
func CreateAiffFile(name string, reader io.ReadSeeker) (*File, error) {
	return createFile(name, reader, binary.BigEndian)
}

func createFile(name string, reader io.ReadSeeker, byteOrder binary.ByteOrder) (*File, error) {
	containerHeader, err := readContainerHeader(reader, byteOrder)

	if err != nil {
		return nil, err
	}

	headers := readChunkHeaders(reader, ContainerHeaderSizeBytes, byteOrder)
	riffFile := &File{Name: name, Header: containerHeader, Headers: headers, ByteOrder: byteOrder}

	return riffFile, nil
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

func readContainerHeader(reader io.ReadSeeker, byteOrder binary.ByteOrder) (*ContainerHeader, error) {
	var headerBytes [ContainerHeaderSizeBytes]byte
	_, err := io.ReadFull(reader, headerBytes[:])

	if err != nil {
		return nil, err
	}

	riffHeader := DecodeContainerHeader(headerBytes, byteOrder)

	return riffHeader, nil
}

// FindHeader returns headers with provided ID.
func (rf *File) FindHeaders(id string) []*Header {
	var headers []*Header

	for _, header := range rf.Headers {
		if header.ID() == id {
			headers = append(headers, header)
		}
	}

	return headers
}

// todo: chunk ids might not be unique. Better delete by offset => GetHeaderByStartPos
// DeleteChunk
func (rf *File) DeleteChunk(headerID string, reader io.ReaderAt, writer io.WriterAt) (uint32, error) {
	foundHeaders := rf.FindHeaders(headerID)

	if foundHeaders == nil {
		msg := fmt.Sprintf("chunk not found: %s", headerID)
		return 0, errors.New(msg)
	}

	//todo: find by offset
	header := foundHeaders[0]
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
func (rf *File) AddChunk(reader io.Reader, writer io.WriterAt, bufferSize int) error {
	if bufferSize <= 0 {
		bufferSize = 1024
	}

	// start writing chunk to end of file
	offset := int64(rf.Header.FullSize())
	var chunkSize uint32 = 0

	for {
		b := make([]byte, bufferSize)
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
func (rf *File) UpdateSize(size uint32, writer io.WriterAt) error {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, size)

	if err != nil {

		return err
	}

	b := buf.Bytes()
	_, err = writer.WriteAt(b, 4)

	if err != nil {
		return err
	}

	rf.Header.size = size

	return nil
}

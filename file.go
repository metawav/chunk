package wav

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

type RiffFile struct {
	Name    string
	Header  *RiffHeader
	Headers []*Header
}

//todo: remove file dependency and use reader as parameter
func CreateRiffFile(file *os.File) (*RiffFile, error) {
	fileInfo, err := file.Stat()

	if err != nil {
		return nil, err
	}

	riffHeader, err := readRiffHeader(file)

	if err != nil {
		return nil, err
	}

	headers := readChunkHeaders(RiffHeaderSizeBytes, file, fileInfo.Size())
	riffFile := &RiffFile{Name: file.Name(), Header: riffHeader, Headers: headers}

	return riffFile, nil
}

func readChunkHeaders(offset uint32, file *os.File, fileSize int64) []*Header {
	var headers []*Header

	for int64(offset) < fileSize {
		var chunkHeaderBytes [HeaderSizeBytes]byte
		n, err := file.ReadAt(chunkHeaderBytes[:], int64(offset))

		if err != nil {
			test := make([]byte, n)
			//todo: handle error
			n, _ := file.ReadAt(test, int64(offset))
			fmt.Printf("ERROR - fileName: %s fileSize: %d offset: %d read bytes: %d raw: %x error: %+v\n", file.Name(), fileSize, offset, n, test, err)
			offset += uint32(n)
			continue
		}

		chunkHeader := DecodeChunkHeader(chunkHeaderBytes, offset)
		headers = append(headers, chunkHeader)
		offset += chunkHeader.FullSize()

		// For compatibility with EA IFF (Electronic Arts Interchange File Format)
		// chunks must be even sized and always start at an even position.
		if offset%2 != 0 {
			offset++
		}
	}

	return headers
}

func readRiffHeader(file *os.File) (*RiffHeader, error) {
	var headerBytes [RiffHeaderSizeBytes]byte
	_, err := file.ReadAt(headerBytes[:], 0)

	if err != nil {
		return nil, err
	}

	riffHeader := DecodeRiffHeader(headerBytes)

	return riffHeader, nil
}

func (rf *RiffFile) GetHeaderByID(headerID string) (*Header, error) {
	for _, header := range rf.Headers {
		if header.ID() == headerID {
			return header, nil
		}
	}

	msg := fmt.Sprintf("header not found: %s", headerID)

	return nil, errors.New(msg)
}

func (rf *RiffFile) DeleteChunk(headerID string, reader io.ReaderAt, writer io.WriterAt) (uint32, error) {
	header, err := rf.GetHeaderByID(headerID)

	if err != nil {
		msg := fmt.Sprintf("chunk not found: %s", headerID)
		return 0, errors.New(msg)
	}

	//todo: implement
	// update riff header size
	// use reader and writer interface to avoid dependency
	headers := rf.Headers
	sort.Sort(SortBy(headers))

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
	err = rf.UpdateSize(riffSize, writer)

	if err != nil {
		return 0, err
	}

	fileSize := rf.Header.FullSize() - header.FullSize()

	return fileSize, nil
}

//todo: move to header file and make internal
type SortBy []*Header

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].StartPos() < a[j].StartPos() }

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

package wav

import (
	"errors"
	"fmt"
	"os"
)

type RiffFile struct {
	name    string
	header  *RiffHeader
	headers []*Header
}

// NewFile creates a new RIFF file.
func NewRiffFile(name string, header *RiffHeader, headers []*Header) *RiffFile {
	return &RiffFile{name: name, header: header, headers: headers}
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
	riffFile := NewRiffFile(file.Name(), riffHeader, headers)

	return riffFile, nil
}

// FileName returns the file name.
func (rf *RiffFile) FileName() string {
	return rf.name
}

// RiffHeader returns the RIFF header.
func (rf *RiffFile) RiffHeader() *RiffHeader {
	return rf.header
}

//todo: rename to Headers
// Headers returns all contained chunk headers.
func (rf *RiffFile) Headers() []*Header {
	return rf.headers
}

func readChunkHeaders(offset uint32, file *os.File, fileSize int64) []*Header {
	var headers []*Header

	for int64(offset) < fileSize {
		chunkHeaderBytes := make([]byte, HeaderSizeBytes)
		n, err := file.ReadAt(chunkHeaderBytes, int64(offset))

		if err != nil {
			test := make([]byte, n)
			//todo: handle error
			n, _ := file.ReadAt(test, int64(offset))
			fmt.Printf("ERROR - fileName: %s fileSize: %d offset: %d read bytes: %d raw: %x error: %+v\n", file.Name(), fileSize, offset, n, test, err)
			offset += uint32(n)
			continue
		}

		chunkHeader, err := DecodeChunkHeader(chunkHeaderBytes, offset)

		if err == nil {
			headers = append(headers, chunkHeader)
			offset += chunkHeader.FullSize()
			continue
		}

		offset += uint32(n)

		// For compatibility with EA IFF (Electronic Arts Interchange File Format)
		// chunks must be even sized and always start at an even position.
		if offset%2 != 0 {
			offset++
		}
	}

	return headers
}

func readRiffHeader(file *os.File) (*RiffHeader, error) {
	headerBytes := make([]byte, RiffHeaderSizeBytes)
	_, err := file.ReadAt(headerBytes, 0)

	if err != nil {
		return nil, err
	}

	riffHeader, err := DecodeRiffHeader(headerBytes)

	return riffHeader, err
}

func (rf *RiffFile) GetHeaderByID(headerID string) (*Header, error) {
	for _, header := range rf.headers {
		if header.ID() == headerID {
			return header, nil
		}
	}

	msg := fmt.Sprintf("header not found: %s", headerID)

	return nil, errors.New(msg)
}

func (rf *RiffFile) DeleteChunk(headerID string) error {
	//todo: implement
	header, err := rf.GetHeaderByID(headerID)

	if err != nil {
		msg := fmt.Sprintf("chunk not found: %s", headerID)
		return errors.New(msg)
	}

	fmt.Printf("Deleting chunk %s\n", header.ID())

	return nil
}

func (rf *RiffFile) AddChunk(header *Header) error {
	//todo: implement
	fmt.Printf("Adding chunk %s", header.ID())
	return nil
}

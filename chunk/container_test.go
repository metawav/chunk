package chunk

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestFindHeaders(t *testing.T) {
	riffFile := &Container{}

	for _, h := range riffFile.Headers {
		fmt.Printf("%v\n", h)
	}

	headers := riffFile.FindHeaders("")

	if headers != nil {
		t.Errorf("headers should not be returned")
	}

	headerID := createID("test")
	header := EncodeChunkHeader(headerID, 0, binary.LittleEndian)

	riffFile.Headers = append(riffFile.Headers, header)
	headers = riffFile.FindHeaders(string(headerID[:]))

	if headers == nil {
		t.Errorf("headers should be returned")
	}

	if len(headers) != 1 {
		t.Errorf("headers size is %d, want %d", len(headers), 1)
	}

	riffFile.Headers = append(riffFile.Headers, header)
	headers = riffFile.FindHeaders(string(headerID[:]))

	if len(headers) != 2 {
		t.Errorf("headers size is %d, want %d", len(headers), 2)
	}
}

func createID(idVal string) [IDSizeBytes]byte {
	var id [IDSizeBytes]byte
	copy(id[:], idVal)

	return id
}

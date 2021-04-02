package chunk

import (
	"encoding/binary"
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

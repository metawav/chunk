package chunk

import "bytes"

func trimNullChars(data []byte) []byte {
	return bytes.Trim(data, "\x00")
}

// trim converts byte array to string with null characters removed.
func trim(data []byte) string {
	trimmed := trimNullChars(data)

	return string(trimmed)
}

func terminate(value string, maxLength int) string {
	if maxLength == len(value) {
		return value
	}

	if maxLength < len(value) {
		return value[:maxLength]
	}

	i := bytes.IndexByte([]byte(value), '\x00')

	if i >= 0 {
		return value
	}

	return value + "\x00"
}

func nullTermToString(b []byte) string {
	i := bytes.IndexByte(b[:], '\x00')

	if i >= 0 {
		return string(b[:i])
	}

	return trim(b)
}

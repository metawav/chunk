package wav

import "bytes"

func trimNullChars(data []byte) []byte {
	return bytes.Trim(data, "\x00")
}

// Trim converts byte array to string with null characters removed.
func trim(data []byte) string {
	trimmed := trimNullChars(data)

	return string(trimmed)
}

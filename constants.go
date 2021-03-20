package wav

const (
	// IDSizeBytes is byte size of the id value
	IDSizeBytes uint32 = 4
	// SizeSizeBytes is byte size of the size value
	SizeSizeBytes uint32 = 4
	// HeaderSizeBytes is byte size of chunk header (id + size)
	HeaderSizeBytes uint32 = IDSizeBytes + SizeSizeBytes
	// FormatSizeBytes is byte size of the format value
	FormatSizeBytes uint32 = 4
	// RiffHeaderSizeBytes is byte size of the RIFF header
	RiffHeaderSizeBytes uint32 = HeaderSizeBytes + FormatSizeBytes
)

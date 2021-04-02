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
	// ContainerHeaderSizeBytes is byte size of the container header
	ContainerHeaderSizeBytes uint32 = HeaderSizeBytes + FormatSizeBytes
)

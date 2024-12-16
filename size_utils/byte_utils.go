package size_utils

type ByteSize uint64

const (
	KB ByteSize = 1024
	MB          = KB * 1024
	GB          = MB * 1024
	TB          = GB * 1024
	PB          = TB * 1024
)

// Converts the given number of bytes to the selected size, using size.KB etc.
// Example (2,147,483,648, size.GB) will return 2.
func ConvertByteTo(byteCount uint64, convertTo ByteSize) uint64 {
	return byteCount / uint64(convertTo)
}

// Converts the given number of magnitude bytes to the number of bytes.
// Example: (2, size.GB) will return 2,147,483,648.
func ConvertToByte(byteCount uint64, magnitude ByteSize) uint64 {
	return byteCount * uint64(magnitude)
}

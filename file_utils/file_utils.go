package file_utils

import (
	"bufio"
	"io"
	"os"
	"path/filepath"

	"gotils/arr_utils"
)

const chunkSize = 512

// CopyFile copies source path to dest path.
func CopyFile(source string, dest string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

// Returns whether the path contains an image file extension.
// Extension if checked against internal list:
// {".jpg", ".jpeg", ".png", ".tif", ".tiff", ".iff", ".tga", ".exr"}
func IsImageFile(filePath string) bool {
	var imageFileExtensions = []string{
		".jpg", ".jpeg", ".png", ".tif", ".tiff", ".iff", ".tga", ".exr",
	}
	ext := filepath.Ext(filePath)
	return arr_utils.StringSliceContains(imageFileExtensions, ext)
}

// IsBinaryFile determines if a file is binary by examining its content.
func IsBinaryFile(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, chunkSize)

	// Read the first chunk of the file
	n, err := reader.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return false, err
	}

	// Check each byte for non-ASCII characters
	for i := 0; i < n; i++ {
		if buffer[i] > 127 {
			return true, nil
		}
	}

	return false, nil
}

package file

import (
	"path/filepath"

	"github.com/chigopher/pathlib"

	"gotils/arr"
)

var imageFileExtensions = []string{".jpg", ".jpeg", ".png", ".tif", ".tiff", ".iff", ".tga", ".exr"}

// Returns whether the path contains an image file extension.
// Extension if checked against internal list:
// {".jpg", ".jpeg", ".png", ".tif", ".tiff", ".iff", ".tga", ".exr"}
func IsImageFile(filePath pathlib.Path) bool {
	ext := filepath.Ext(filePath.String())
	return arr.StringSliceContains(imageFileExtensions, ext)
}

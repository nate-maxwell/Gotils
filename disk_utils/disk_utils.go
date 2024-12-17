package disk_utils

import (
	"os"
	"path/filepath"

	"github.com/chigopher/pathlib"
	"golang.org/x/sys/windows"
)

// Gets the various space statistics for teh given path.
// Takes path(pathlib.Path).
// Returns freeBytesAvailable(int), totalBytes(int), totalFreeBytes(int), err(error).
func GetDriveFreeSpace(path pathlib.Path) (uint64, uint64, uint64, error) {
	var freeBytesAvailable, totalBytes, totalFreeBytes uint64

	pathPtr, err := windows.UTF16PtrFromString(path.String())
	if err != nil {
		return 0, 0, 0, err
	}

	err = windows.GetDiskFreeSpaceEx(pathPtr, &freeBytesAvailable, &totalBytes, &totalFreeBytes)
	if err != nil {
		return 0, 0, 0, err
	}

	return freeBytesAvailable, totalBytes, totalFreeBytes, nil
}

// Returns the byte size of the directory by recursively indexing its contents.
func GetDirSize(folderPath pathlib.Path) (int64, error) {
	var totalSize int64

	err := filepath.WalkDir(folderPath.String(), func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			totalSize += info.Size()
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return totalSize, nil
}

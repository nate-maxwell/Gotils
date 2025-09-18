package disk_utils

import (
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

// GetDriveFreeSpace gets the various space statistics for given path.
// Returns freeBytesAvailable(int), totalBytes(int), totalFreeBytes(int), err(error).
func GetDriveFreeSpace(path string) (uint64, uint64, uint64, error) {
	var freeBytesAvailable, totalBytes, totalFreeBytes uint64

	pathPtr, err := windows.UTF16PtrFromString(path)
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
func GetDirSize(folderPath string) (int64, error) {
	var totalSize int64

	err := filepath.WalkDir(folderPath, func(path string, d os.DirEntry, err error) error {
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

package disk_utils

import (
	"github.com/chigopher/pathlib"
	"golang.org/x/sys/windows"
)

// Gets the various space statistics for teh given path.
// Takes path(pathlib.Path).
// Returns freeBytesAvailable(int), totalBytes(int), totalFreeBytes(int), err(error).
func GetFreeSpace(path pathlib.Path) (uint64, uint64, uint64, error) {
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

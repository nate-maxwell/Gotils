package dir_utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gotils/file_utils"
	"gotils/time_utils"
)

// CountFilesByName returns the number of files within rootDir of targetName.
func CountFilesByName(rootDir string, targetName string) (int, error) {
	count := 0
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == targetName {
			count++
		}
		return nil
	})
	return count, err
}

// GetDirContents returns array of items in path or error.
func GetDirContents(path string) ([]string, error) {
	var contents []string

	items, err := os.ReadDir(path)
	if err != nil {
		return make([]string, 0), err
	}
	for _, item := range items {
		entry := item.Name()
		contents = append(contents, entry)
	}
	return contents, nil
}

// CreateDatedDirectory creates a directory with today's date as the name.
func CreateDatedDirectory(path string) error {
	datePath := filepath.Join(path, time_utils.GetDate())
	return os.MkdirAll(datePath, 0777)
}

// DeleteSafeItemsInDir deletes all items in a directory.
func DeleteSafeItemsInDir(folderPath string) error {
	files, err := GetDirContents(folderPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		err := os.Remove(f)
		if err != nil {
			return err
		}
	}
	return nil
}

// CopyFolderContents copies contents of source into dest.
func CopyFolderContents(source string, dest string) error {
	err := os.MkdirAll(dest, 0777)
	if err != nil {
		return err
	}

	curItems, err := GetDirContents(source)
	if err != nil {
		return err
	}

	for _, item := range curItems {
		curItemPath := filepath.Join(source, item)
		destPath := filepath.Join(dest, item)

		curInfo, err := os.Stat(curItemPath)
		if err != nil {
			return err
		}
		dir := curInfo.IsDir()

		if dir {
			err := CopyFolderContents(curItemPath, destPath)
			if err != nil {
				return err
			}
		} else {
			err := file_utils.CopyFile(curItemPath, destPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// DeleteFilesOlderThan deletes items older than num days in path dir.
func DeleteFilesOlderThan(dir string, days int) error {
	cutoff := time.Now().AddDate(0, 0, -days)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.ModTime().Before(cutoff) {
			fmt.Printf("Deleting %s\n", path)
			if err := os.Remove(path); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

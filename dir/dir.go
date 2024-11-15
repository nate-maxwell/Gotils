// * Directory Utilities
//
// A simple toolkit for folder and file handling that eliminates
// boilerplate or wraps commonly used functions in a consistent
// namespace for easy remembrance/importing.

package dir

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/chigopher/pathlib"

	"gotils/time_utils"
)

// Returns the number of files within a directory structure of the targetName.
//
// Args:
//
//	rootDir(pathlib.Path): The root of the structure to count the files in.
//	targetName9string): The name the counted files should have.
//
// Returns:
//
//	int: The number of files containing the targetName within the structure.
//	error: Any encountered error.
func CountFilesByName(rootDir pathlib.Path, targetName string) (int, error) {
	count := 0
	err := filepath.Walk(rootDir.String(), func(path string, info os.FileInfo, err error) error {
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

// Gets the contents' names, or full path for contents, of a directory.
//
// Args:
//
//	path(pathlib.Path): Directory path to list the contents of.
//	fullPath(bool): To return string names or full paths of directory contents.
//
// Returns:
//
//	[]string: String names or full paths of directory contents.
//	error: Any error created from attempting to read the directory, else nil.
func GetDirContents(path pathlib.Path, fullPath bool) ([]string, error) {
	var contents []string

	items, err := os.ReadDir(path.String())
	if err != nil {
		return make([]string, 0), err
	}
	for _, item := range items {
		var entry string
		if fullPath {
			entry = fmt.Sprintf("%s/%s", path.String(), item.Name())
		} else {
			entry = item.Name()
		}
		contents = append(contents, entry)
	}
	return contents, nil
}

// Creates a directory from the given path.
//
// Args:
//
//	path(pathlib.Path): The directory path to create.
//
// Returns:
//
//	error: Any error created while attempting to create the directory, else nil.
func CreateDirectory(path pathlib.Path) error {
	exists, err := path.Exists()
	if err != nil {
		return err
	}
	if exists {
		err = os.MkdirAll(path.String(), 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

// Creates a directory with today's date as the name.
//
// Args:
//
//	path(pathlib.Path): The path to create the new folder in.
//
// Returns:
//
//	error: Any error created while attempting to create the directory, else nil.
func CreateDatedDirectory(path pathlib.Path) error {
	datePath := path.Join(time_utils.GetDate())
	err := CreateDirectory(*datePath)
	if err != nil {
		return err
	}
	return nil
}

// Deletes a directory and its contents.
//
// Args:
//
//	folderPath(pathlib.Path): The folder path to delete.
//
// Returns:
//
//	error: the *PathError created from os.RemoveAll if one was created, else nil.
func DeleteDirectory(folderPath pathlib.Path) error {
	err := os.RemoveAll(folderPath.String())
	if err != nil {
		return err
	}
	return nil
}

// Removes specified file.
//
// Args:
//
//	filepath(pathlib.Path): The path to the file you wish to delete.
//
// Returns:
//
//	error: A custom error if the filepath was not within the safety path or a *PathError err from
//	os.Remove, else Nil.
func DeleteFile(filepath pathlib.Path) error {
	err := os.Remove(filepath.String())
	if err != nil {
		return err
	}
	return nil
}

// Delete all files in a directory.
//
// Args:
//
//	directory_path(pathlib.Path): The path to the directory.
//
// Returns:
//
//	any *PathError crated from DeleteSafeFile or errors from GetDirContents, else nil.
func DeleteSafeFilesInDirectory(folderPath pathlib.Path) error {
	files, err := GetDirContents(folderPath, true)
	if err != nil {
		return err
	}
	for _, file := range files {
		err := DeleteFile(*pathlib.NewPath(file))
		if err != nil {
			return err
		}
	}
	return nil
}

// Copy file into a separate destination folder.
//
// Args:
//
//	source(pathlib.Path): File path of the file to copy.
//	dest(pathlib.Path): File path to copy the file too, optionally can have different name.
//
// Returns:
//
//	error: *PathError crated from os module or possible other error from io module else nil.
func CopyFile(source pathlib.Path, dest pathlib.Path) error {
	sourceFile, err := os.Open(source.String())
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest.String())
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

// Copy contents of a folder to the given destination.
//
// Args:
//
//	sourcePath(pathlib.Path): Folder path to the folder that is to be copied.
//	destination(pathlib.Path): Folder path to copy the folder + contents to.
//
// Returns:
//
//	error: Any relevant errors created durring process, usually os *PathErrors else nil.
func CopyFolderContents(sourcePath pathlib.Path, destination pathlib.Path) error {
	err := CreateDirectory(destination)
	if err != nil {
		return err
	}

	curItems, err := GetDirContents(sourcePath, false)
	if err != nil {
		return err
	}

	for _, item := range curItems {
		curItemPath := sourcePath.Join(item)
		destPath := destination.Join(item)

		dir, err := curItemPath.IsDir()
		if err != nil {
			return err
		}
		if dir {
			err := CopyFolderContents(*curItemPath, *destPath)
			if err != nil {
				return err
			}
		} else {
			err := CopyFile(*curItemPath, *destPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Exports a string map to json file path.
//
// Args:
//
//	filepath(pathlib.Path): The file path to place the .json file.
//	data(map[string]interface{}): Any map with string keys and values that can be converted to strings.
//	overWrite(bool): To overwrite json file if it already exists in path.
//
// Returns:
//
//	error: Any relevant error from the json handling or file writing process.
func ExportMapToJson(filePath pathlib.Path, data map[string]interface{}, overWrite bool) error {
	exists, err := filePath.Exists()
	if err != nil {
		return err
	}

	if !exists || overWrite {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}

		file, err := os.Create(filePath.String())
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.Write(jsonData)
		if err != nil {
			return err
		}

		return nil
	}
	return nil
}

// DeleteFilesOlderThan deletes files older than the specified duration in the given directory.
//
// Args:
//
//	dir(pathlib.Path): The path to check the files for.
//	days(int): The age, in days, the file must meet to be deleted.
//
// Returns:
//
//	error: The relevant error to whichever step failed, else nil.
func DeleteFilesOlderThan(dir pathlib.Path, days int) error {
	cutoff := time.Now().AddDate(0, 0, -days)

	err := filepath.Walk(dir.String(), func(path string, info os.FileInfo, err error) error {
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

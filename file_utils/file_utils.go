package file_utils

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/chigopher/pathlib"

	"gotils/arr_utils"
)

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

// Returns whether the path contains an image file extension.
// Extension if checked against internal list:
// {".jpg", ".jpeg", ".png", ".tif", ".tiff", ".iff", ".tga", ".exr"}
func IsImageFile(filePath pathlib.Path) bool {
	var imageFileExtensions = []string{".jpg", ".jpeg", ".png", ".tif", ".tiff", ".iff", ".tga", ".exr"}
	ext := filepath.Ext(filePath.String())
	return arr_utils.StringSliceContains(imageFileExtensions, ext)
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// filterOut filters out files based on the provided criteria.
// Parameters:
//   - path: The file path to check.
//   - ext: The file extension to filter by. If empty, no extension filtering is applied.
//   - minSize: The minimum file size in bytes. Files smaller than this size are filtered out.
//   - info: The FileInfo structure containing file metadata.
//
// Returns:
//   - true if the file should be filtered out, false otherwise.
func filterOut(path, ext string, minSize int64, info os.FileInfo) bool {
	if info.IsDir() || info.Size() < minSize {
		return true
	}

	if ext != "" && filepath.Ext(path) != ext {
		return true
	}

	return false
}

// listFile writes the given file path to the provided writer.
// Parameters:
//   - path: The file path to list.
//   - out: The writer to output the file path.
//
// Returns:
//   - An error if writing to the writer fails, nil otherwise.
func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}


// delFile deletes the file at the specified path and logs the deletion.
//
// Parameters:
//   - path: The path to the file to be deleted.
//   - delLogger: A logger to log the deletion of the file.
//
// Returns:
//   - error: An error if the file could not be deleted, otherwise nil.
func delFile(path string, delLogger *log.Logger) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	delLogger.Println(path)
	return nil
}

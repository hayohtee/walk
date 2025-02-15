package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// filterOut checks if the given path has to be filtered out according
// to the following conditions:
//   - the path points to a directory
//   - the file is less than the minimum size
//   - the file extension doesn't match the provided extension.
func filterOut(path, ext string, minSize int64, info os.FileInfo) bool {
	if info.IsDir() || info.Size() < minSize {
		return true
	}

	if ext != "" && filepath.Ext(path) != ext {
		return true
	}

	return false
}

// listFile writes the path of the file to the STDOUT.
func listFile(path string, out io.Writer) error {
	_, err := fmt.Fprintln(out, path)
	return err
}

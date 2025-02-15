package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// config holds the configurations for the walk CLI.
type config struct {
	// ext represents the extension to filter out.
	ext string
	// size represents the minimum file size.
	size int64
	// list determine whether to list the files or not.
	list bool
}

func main() {
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// run holds the implementation for the CLI.
//
// It walks through the root directory, performing the
// requested operations on the files and sub-directories
// it encounter as specified in the cfg parameter.
func run(root string, out io.Writer, cfg config) error {
	return filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}

		// If list was explicitly set, don't do anything else.
		if cfg.list {
			return listFile(path, out)
		}

		// List is the default option if nothing else was set.
		return listFile(path, out)
	})
}

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

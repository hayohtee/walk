package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
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
	// del determine whether to delete the files or not.
	del bool
	// out represents the log destination writer.
	out io.Writer
	// archive is the name of the directory to store the archived file.
	archive string
}

func main() {
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	del := flag.Bool("del", false, "Delete files")
	logFile := flag.String("log", "", "Log deletes to this file")
	flag.Parse()

	var out io.Writer = os.Stdout

	if *logFile != "" {
		f, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
		out = f
	}

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
		out:  out,
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
	delLogger := log.New(cfg.out, "DELETED FILE: ", log.LstdFlags)

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

		if cfg.del {
			return delFile(path, delLogger)
		}

		// List is the default option if nothing else was set.
		return listFile(path, out)
	})
}

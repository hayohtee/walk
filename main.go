package main

import (
	"flag"
	"fmt"
	"os"
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

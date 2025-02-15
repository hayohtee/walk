package main

// config holds the configurations for the walk CLI.
type config struct {
	// ext represents the extension to filter out.
	ext string
	// size represents the minimum file size.
	size int64
	// list determine whether to list the files or not.
	list bool
}

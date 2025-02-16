package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		root     string
		cfg      config
		expected string
	}{
		{
			name:     "NoFilter",
			root:     "testdata",
			cfg:      config{ext: "", size: 0, list: true},
			expected: "testdata/dir.log\ntestdata/dir2/script.sh\n",
		},
		{
			name:     "FilterExtensionMatch",
			root:     "testdata",
			cfg:      config{ext: ".log", size: 0, list: true},
			expected: "testdata/dir.log\n",
		},
		{
			name:     "FilterExtensionSizeMatch",
			root:     "testdata",
			cfg:      config{ext: ".log", size: 10, list: true},
			expected: "testdata/dir.log\n",
		},
		{
			name:     "FilterExtensionSizeNoMatch",
			root:     "testdata",
			cfg:      config{ext: ".log", size: 20, list: true},
			expected: "",
		},
		{
			name:     "FilterExtensionNoMatch",
			root:     "testdata",
			cfg:      config{ext: ".gz", size: 0, list: true},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			if err := run(tc.root, &buf, tc.cfg); err != nil {
				t.Fatal(err)
			}

			res := buf.String()

			if tc.expected != res {
				t.Errorf("expected %q, got %q instead", tc.expected, res)
			}
		})
	}
}

func createTempDir(t *testing.T, files map[string]int) (string, func()) {
	t.Helper()

	tempDir, err := os.MkdirTemp("", "walktest")
	if err != nil {
		t.Fatal(err)
	}

	for k, n := range files {
		for j := 1; j <= n; j++ {
			fileName := fmt.Sprintf("file%d%s", j, k)
			filePath := filepath.Join(tempDir, fileName)
			if err := os.WriteFile(filePath, []byte("dummy"), 0644); err != nil {
				t.Fatal(err)
			}
		}
	}

	return tempDir, func() { os.RemoveAll(tempDir) }
}

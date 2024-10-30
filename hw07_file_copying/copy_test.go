package main

import (
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	tempDir := os.TempDir()
	defer os.RemoveAll(tempDir)

	srcFile := tempDir + "/source.txt"
	dstFile := tempDir + "/dest.txt"

	// Write some content to the source file
	err := os.WriteFile(srcFile, []byte("hello world"), 0o644)
	if err != nil {
		t.Fatalf("Could not write to source file: %v", err)
	}

	// Copy the file with no offsets or limits
	err = CopyFile(srcFile, dstFile, 0, 0)
	if err != nil {
		t.Fatalf("CopyFile failed: %v", err)
	}
}

package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "golang-test")
	if err != nil {
		t.Fatalf("Could not create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	srcFile := tempDir + "/source.txt"
	dstFile := tempDir + "/dest.txt"

	// Write some content to the source file
	err = os.WriteFile(srcFile, []byte("hello world"), 0644)
	if err != nil {
		t.Fatalf("Could not write to source file: %v", err)
	}

	// Copy the file with no offsets or limits
	err = CopyFile(srcFile, dstFile, 0, 0)
	if err != nil {
		t.Fatalf("CopyFile failed: %v", err)
	}
}

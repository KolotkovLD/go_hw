package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

//func Copy(fromPath, toPath string, offset, limit int64) error {
//	// Place your code here.
//	return nil
//}

func Copy(from, to string, offset, limit int64) error {
	src, err := os.Open(from)
	if err != nil {
		return err
	}
	defer src.Close()

	fi, err := src.Stat()
	if err != nil {
		return err
	}
	if offset > fi.Size() {
		return fmt.Errorf("offset is larger than filesize \n")
	}

	if limit == 0 || offset+limit > fi.Size() {
		limit = fi.Size() - offset
	}

	_, err = src.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	dst, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dst.Close()
	//bar := progressbar.DefaultBytes(limit, "Copying")
	_, err = io.Copy(io.MultiWriter(dst), io.LimitReader(src, limit)) // , bar
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"os"
)

type fita struct {
	file *os.File
}

func (f *fita) Write(p []byte) (n int, err error) {
	f.file.Truncate(0)
	f.file.Seek(0, 0)
	return f.file.Write(p)
}

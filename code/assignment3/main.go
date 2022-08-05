package main

import (
	"fmt"
	"io"
	"os"
)

type MyReader struct {
	file *os.File
}

type MyWriter struct{}

func (r MyReader) Read(p []byte) (n int, err error) {
	return r.file.Read(p)
}

func (MyWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))

	return len(p), nil
}

func main() {
	file, _ := getFile()

	r := MyReader{file}
	w := MyWriter{}

	io.Copy(w, r)
}

func getFile() (*os.File, error) {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("No file specified")
		os.Exit(1)
	}

	filename := args[1]

	return os.Open(filename)
}

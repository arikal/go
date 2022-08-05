package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type logWriter struct{}

func (logWriter) Write(p []byte) (int, error) {
	n := 0

	for _, b := range p {
		fmt.Println(strings.ToUpper(string(b)))
		n++
	}

	return n, nil
}

func main() {
	resp, err := http.Get("https://www.example.com/")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

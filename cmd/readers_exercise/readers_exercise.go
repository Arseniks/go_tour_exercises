package main

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct{}

type ReaderError int

func (e ReaderError) Error() string {
	return fmt.Sprintf("too short b capacity = %v", int(e))
}

func (r MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, ReaderError(cap(b))
	}
	count := 0
	for i := range b {
		b[i] = 'A'
		count++
	}
	return count, nil
}

func main() {
	fmt.Println('A' + 2)
	reader.Validate(MyReader{})
}

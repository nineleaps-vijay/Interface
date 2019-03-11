package main

import (
	"fmt"
	"io"
	"os"
	"bytes"
)

func main() {
	var w io.Writer
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	w.Write([]byte("Hello\n"))
	w = new(bytes.Buffer)
	w.Write([]byte("World"))
	fmt.Printf("%T\n", w)
	w = nil
	fmt.Printf("%T\n", w)

}
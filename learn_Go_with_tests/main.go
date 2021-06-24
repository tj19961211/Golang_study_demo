package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(write io.Writer, name string) {
	fmt.Fprintf(write, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}

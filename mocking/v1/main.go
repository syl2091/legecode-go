package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	fmt.Fprintln(out, "3")
}

func main() {
	Countdown(os.Stdout)
}

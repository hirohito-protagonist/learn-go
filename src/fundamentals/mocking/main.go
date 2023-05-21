package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, strconv.FormatInt(int64(i), 10))
	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintf(out, strconv.FormatInt(int64(i), 10))
		time.Sleep(1 * time.Second)
	}
	fmt.Fprintf(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}

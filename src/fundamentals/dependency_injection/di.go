package di

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// func GreetHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "world")
// }

// func main() {
// 	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(GreetHandler)))
// }

package main

import (
	"fmt"
	"os"
)

// 1.1  Modify to print the os.Args[0]
func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

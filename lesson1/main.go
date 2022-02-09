package main

import (
	"flag"
	"fmt"
	"lesson1/biroot"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		useFile := flag.String("file", "", "-file=[filename] - file using for open discriptor")
		flag.Parse()
		biroot.ReadArgs(*useFile)
	} else {
		fmt.Println("Too much arguments")
		os.Exit(1)
	}
}

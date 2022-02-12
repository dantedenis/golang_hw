package main

import (
	"homework3/executer"
	"homework3/parser"
	"os"
)

func main() {
	str, check := parser.Start()
	if !check {
		os.Exit(1)
	}
	executer.Exec(str)
	return
}

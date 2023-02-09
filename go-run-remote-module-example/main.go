package main

import "fmt"

//go:generate go run golang.org/x/tools/cmd/stringer@v0.1.10 -type=Level

type Level int

const (
	Info Level = iota
	Error
	Fatal
)

func main() {
	fmt.Printf("%s: Hello world!\n", Info)
}

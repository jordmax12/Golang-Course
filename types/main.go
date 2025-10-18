package main

import "fmt"

type str string

func (s str) log() {
	fmt.Println(s)
}

func main() {
	var s str = "Hello, World!"
	s.log()
}

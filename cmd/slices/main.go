package main

import (
	"fmt"
	"github.com/roman-baldaev/golang-learn/slices"
)

func main() {
	s := "        Hello  world    "
	b := []byte(s)
	//slices.SpaceMerge(b)
	//copy(b[5:9], b[5:])
	s = string(slices.SpaceMerge(b))
	fmt.Println(s, len(s)) // " Hello world", len = 12
}

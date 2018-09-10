package main

import (
	"fmt"
	"github.com/roman-baldaev/golang-learn/strings"
)

func main() {
	fmt.Println(strings.RecursiveSeparator("adnasdadaakd", "!", 2))
	fmt.Println(strings.BufferSeparator("adnasdadadakd", "!", 2))
}

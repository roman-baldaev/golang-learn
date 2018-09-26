package main

import (
	"os"
	"github.com/roman-baldaev/golang-learn/html"
	)

func main() {
	for _, url := range os.Args[1:] {
		html.outline(url)

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//f, err := ioutil.ReadFile("data/test_data/files/wordfreq.txt")
	f, err := os.Open("data/test_data/files/wordfreq.txt")
	if err != nil {
		panic(err)
	}
	//var words []string
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	//_, str, err := bufio.ScanWords(f, false)
	//fmt.Println(string(str))
}

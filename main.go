package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("At least one file needed as input")
		fmt.Println("Usage: golisp INPUT_FILE")
		os.Exit(1)
	}

	input, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bufR := bufio.NewReader(input)
	lNumber := 1
	for {
		line, isPrefix, err := bufR.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			break
		}

		fmt.Printf("%d: %s\n", lNumber, string(line))
		if !isPrefix {
			lNumber = lNumber + 1
		}
	}
}

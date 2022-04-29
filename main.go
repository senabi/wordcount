package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// get args
	args := os.Args[1:]
	file, err := os.Open("./generator/lorem.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	target := args[0]
	ix := int64(1)
	occurencies := int64(0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linestring := scanner.Text()
		for _, word := range strings.Split(linestring, " ") {
			for _, w := range strings.Split(word, ".") {
				if w == target {
					occurencies++
				}
			}
		}
		ix++
	}
	fmt.Println(target, ":", occurencies)
}

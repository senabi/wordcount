package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./generator/lorem.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	wordsMap := make(map[string]int)

	scanner := bufio.NewScanner(file)
	ix := int64(1)
	for scanner.Scan() {
		linestring := scanner.Text()
		for _, word := range strings.Split(linestring, " ") {
			for _, w := range strings.Split(word, ".") {
				wordsMap[w]++
			}
		}
		ix++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("lines:", ix)
	fmt.Println("unique words:", len(wordsMap))

	file, err = os.Create("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for k, v := range wordsMap {
		fmt.Fprintln(file, k, v)
	}
}

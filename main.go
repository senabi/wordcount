package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	// get args
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("no input file")
	}
	inputFile := args[0]
	// file size
	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()
	fmt.Println("file size:", fileSize)

	wordsMap := make(map[string]int)
	wg := sync.WaitGroup{}
	signal := make(chan bool, 1)
	wordsChan := make(chan string)

	start := time.Now()
	// count words
	go func() {
		for word := range wordsChan {
			wordsMap[word]++
		}
		signal <- true
	}()

	currOffset := int64(0)
	limit := int64(100 * 1024 * 1024) // 100MB per thread

	for i := int64(0); i < fileSize/limit+2; i++ {
		wg.Add(1)
		fmt.Println("sending", i)
		go func() {
			sender(inputFile, wordsChan, currOffset, limit, fileSize)
			fmt.Println("thread", i, "done")
			wg.Done()
		}()
		currOffset += limit + 1
	}

	wg.Wait()
	close(wordsChan)

	<-signal
	close(signal)

	fmt.Println("read", time.Since(start))
	fmt.Println("unique words:", len(wordsMap))

	start = time.Now()
	file, err := os.Create("./words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for k, v := range wordsMap {
		fmt.Fprintln(file, k, v)
	}
	fmt.Println("write", time.Since(start))
}

func sender(fileName string, wordsChan chan (string), offset int64, limit int64, fileSize int64) {
	if offset > fileSize {
		fmt.Println("out of bounds offset")
		return
	}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Seek(offset, 0)
	reader := bufio.NewReader(file)
	bytesRead := int64(0)

	for {
		if bytesRead > limit {
			break
		}
		bytes, err := reader.ReadBytes(' ')
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		bytesRead += int64(len(bytes))
		str := strings.TrimSpace(string(bytes))
		for _, v := range strings.Split(str, "\n") {
			wordsChan <- v
		}
	}
}

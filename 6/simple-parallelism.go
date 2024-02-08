package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	inputChannel := make(chan string)
	quitChan := make(chan string)
	wg.Add(1)
	go printEveryLetter(inputChannel, quitChan)

	inputChannel <- "Ufuk Baris Tunca"
	inputChannel <- "Golang"
	inputChannel <- "Concurrency"
	inputChannel <- "Parallelism"
	quitChan <- "quit"
	wg.Wait()
}

func printEveryLetter(inputChan, outputChan chan string) {
	defer wg.Done()
	for {
		select {
		case chanInput := <-inputChan:
			fmt.Println(chanInput)
		case <-outputChan:
			return
		}
	}
}

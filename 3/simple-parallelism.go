package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	startTime := time.Now()
	urls := []string{"https://www.google.com.tr", "https://pkg.go.dev/net/http", "https://www.modanisa.com"}
	for _, url := range urls {
		wg.Add(1)
		go getWebpage(url)
	}
	wg.Wait()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Total execution time:", elapsedTime)
}

func getWebpage(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(url, string(body))
	wg.Done()
}

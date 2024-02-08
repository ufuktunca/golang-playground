package main

import (
	"fmt"
)

var calculationChan = make(chan string)

func main() {
	numberList := []int{1, 2, 3, 4, 5}

	for _, number := range numberList {
		go calculateASquare(number)
	}

	msg := <-calculationChan
	fmt.Println(msg)
	msg = <-calculationChan
	fmt.Println(msg)
	msg = <-calculationChan
	fmt.Println(msg)
	msg = <-calculationChan
	fmt.Println(msg)
	msg = <-calculationChan
	fmt.Println(msg)
}

func calculateASquare(number int) {
	square := number * number
	squareString := fmt.Sprintf("%d", square)
	calculationChan <- squareString
}

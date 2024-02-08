package main

import (
	"fmt"
)

var calculationChan = make(chan int)

func main() {
	numberList := []int{5}

	for _, number := range numberList {
		go calculateASquare(number)
		go calculateCube(number)
	}

	msg := <-calculationChan
	msg2 := <-calculationChan
	fmt.Println(msg + msg2)

}

func calculateASquare(number int) {
	square := number * number
	calculationChan <- square
}

func calculateCube(number int) {
	cube := number * number * number
	calculationChan <- cube
}

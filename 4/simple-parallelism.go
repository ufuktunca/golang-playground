package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numberChannel := generateNumbers(numbers)
	factorialChannel := calculateFactorialOfNumber(numberChannel)
	printResult(factorialChannel)

	fmt.Scanf("%d")

}

func generateNumbers(numbers []int) chan int {
	numberChannel := make(chan int)

	go func() {
		for _, value := range numbers {
			numberChannel <- value
		}
		close(numberChannel)
	}()
	return numberChannel
}

func calculateFactorialOfNumber(inboundChannel chan int) chan int {
	outboundChannel := make(chan int)

	go func() {
		for value := range inboundChannel {
			totalValue := 1
			for i := 1; i <= value; i++ {
				totalValue = totalValue * i
			}
			outboundChannel <- totalValue
		}
		close(outboundChannel)
	}()

	return outboundChannel
}

func printResult(inboundChannel chan int) {
	for value := range inboundChannel {
		fmt.Println(value)
	}
}

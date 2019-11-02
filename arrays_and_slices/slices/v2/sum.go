package main

import "fmt"

func SumAll(numberstoSum ...[]int) []int {

	var sums []int

	for _, number := range numberstoSum {

		sums = append(sums, Sum(number))
	}
	return sums
}

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	fmt.Println(SumAll([]int{1, 2}, []int{0, 9}))
}

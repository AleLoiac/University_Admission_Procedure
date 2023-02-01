package main

import (
	"fmt"
)

var exam1, exam2, exam3 int

func meanValue(e1, e2, e3 int) float64 {
	sum := float64(e1 + e2 + e3)
	mean := sum / 3
	return mean
}

func main() {

	_, err := fmt.Scan(&exam1, &exam2, &exam3)
	if err != nil {
		return
	}
	fmt.Println(meanValue(exam1, exam2, exam3))

	fmt.Println("Congratulations, you are accepted!")
}

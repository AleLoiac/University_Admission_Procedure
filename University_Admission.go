package main

import (
	"fmt"
)

var exam1, exam2, exam3 int
var score float64

func isAccepted(meanScore float64) {
	if meanScore > 60.0 {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}

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

	score = meanValue(exam1, exam2, exam3)

	fmt.Println(score)

	isAccepted(score)

}

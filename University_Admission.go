package main

import (
	"fmt"
)

var studentsNumber int
var acceptedNumber int

type Student struct {
	fullName string
	GPA      float64
}

func isAccepted(meanScore float64) {
	if meanScore >= 60.0 {
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

}

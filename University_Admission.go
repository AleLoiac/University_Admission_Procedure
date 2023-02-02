package main

import (
	"fmt"
	"sort"
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

func newStudent(name string, gpa float64) Student {
	s := Student{fullName: name}
	s.GPA = gpa
	return s
}

func main() {
	fmt.Scan(&studentsNumber)
	fmt.Scan(&acceptedNumber)

	Students := make([]Student, studentsNumber)

	for i := 0; i < studentsNumber; i++ {
		var name string
		var gpa float64
		fmt.Scan(&name, &gpa)
		Students = append(Students, newStudent(name, gpa))
	}

	sort.Slice(Students, func(i, j int) bool {
		if Students[i].GPA != Students[j].GPA {
			return Students[i].GPA < Students[j].GPA
		}
		return Students[i].fullName < Students[j].fullName
	})
	fmt.Println(Students)

}

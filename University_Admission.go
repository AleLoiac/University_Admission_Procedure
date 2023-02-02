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

func sortStudents(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		if students[i].GPA != students[j].GPA {
			return students[i].GPA > students[j].GPA
		}
		return students[i].fullName > students[j].fullName
	})
}

func main() {
	_, err := fmt.Scan(&studentsNumber, &acceptedNumber)
	if err != nil {
		return
	}

	Students := make([]Student, 0)

	for i := 0; i < studentsNumber; i++ {
		var name string
		var gpa float64
		_, err2 := fmt.Scan(&name, &gpa)
		if err2 != nil {
			return
		}
		Students = append(Students, newStudent(name, gpa))
	}

	sortStudents(Students)

	fmt.Println("Successful applicants:")
	for i := 0; i < acceptedNumber; i++ {
		fmt.Print(Students[i].fullName, " ")
	}

}

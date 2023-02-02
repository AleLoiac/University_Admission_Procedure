package main

import (
	"fmt"
	"sort"
)

var studentsNumber int
var acceptedNumber int

type Student struct {
	firstName string
	lastname  string
	GPA       float64
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

func newStudent(firstname string, lastname string, gpa float64) Student {
	s := Student{firstName: firstname}
	s.lastname = lastname
	s.GPA = gpa
	return s
}

func sortStudents(students []Student) {
	sort.Slice(students, func(i, j int) bool {
		if students[i].GPA != students[j].GPA {
			return students[i].GPA > students[j].GPA
		} else if students[i].firstName != students[j].firstName {
			return students[i].firstName < students[j].firstName
		}
		return students[i].lastname < students[j].lastname
	})
}

func main() {
	_, err := fmt.Scan(&studentsNumber, &acceptedNumber)
	if err != nil {
		return
	}

	Students := make([]Student, 0)

	for i := 0; i < studentsNumber; i++ {
		var firstname, lastname string
		var gpa float64
		_, err2 := fmt.Scan(&firstname, &lastname, &gpa)
		if err2 != nil {
			return
		}
		Students = append(Students, newStudent(firstname, lastname, gpa))
	}

	sortStudents(Students)

	fmt.Println("Successful applicants:")
	for i := 0; i < acceptedNumber; i++ {
		fmt.Println(Students[i].firstName, Students[i].lastname)
	}

}

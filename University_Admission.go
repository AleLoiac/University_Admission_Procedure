package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var acceptedNumber int // the number of accepted students in each department

var Students []Student

var Mathematics []Student
var Physics []Student
var Biotech []Student
var Chemistry []Student
var Engineering []Student

type Student struct {
	firstName    string
	lastname     string
	GPA          float64
	firstChoice  string
	secondChoice string
	thirdChoice  string
}

func fileToSlice(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		field := strings.Split(scanner.Text(), " ")

		var s Student

		var err error

		s.firstName = field[0]
		s.lastname = field[1]
		s.GPA, err = strconv.ParseFloat(field[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		s.firstChoice = field[3]
		s.secondChoice = field[4]
		s.thirdChoice = field[5]

		Students = append(Students, s)
	}
}

func populateDep(department []Student, name string) {
	count := 0
	for i, v := range Students {
		if count == acceptedNumber {
			break
		} else if v.firstChoice == name {
			department = append(department, v)
			Students = append(Students[:i], Students[i+1:]...) //removes the student from the Students slice
			count++
		}
	}
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

	_, err2 := fmt.Scan(&acceptedNumber)
	if err2 != nil {
		return
	}

	file, err := os.Open("applicant_list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileToSlice(file)

	sortStudents(Students)

	Mathematics = make([]Student, 0)

	//count := 0
	//for i, v := range Students {
	//	if count == acceptedNumber {
	//		break
	//	} else if v.firstChoice == "Mathematics" {
	//		Mathematics = append(Mathematics, v)
	//		Students = append(Students[:i], Students[i+1:]...) //removes the student from the Students slice
	//		count++
	//	}
	//}

	populateDep(Mathematics, "Mathematics")
	//Physics = make([]Student, 0)
	//populateDep(Physics)
	//Biotech = make([]Student, 0)
	//populateDep(Biotech)
	//Chemistry = make([]Student, 0)
	//populateDep(Chemistry)
	//Engineering = make([]Student, 0)
	//populateDep(Engineering)

	//fmt.Println("Mat:\n", Mathematics, "Phy:\n", Physics, "Bio:\n", Biotech, "Che:\n", Chemistry, "Eng:\n", Engineering)
	fmt.Println(Mathematics)
	fmt.Println()
	fmt.Println(Students)
}

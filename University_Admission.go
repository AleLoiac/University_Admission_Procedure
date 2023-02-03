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
	assigned     bool
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

func populateDep(department *[]Student, name string) {
	count := 0
	for i, v := range Students {
		if count == acceptedNumber {
			break
		} else if v.firstChoice == name && v.assigned != true {
			v.assigned = true
			Students[i].assigned = true
			*department = append(*department, v)
			count++
		}
	}
}

func secondRound(department *[]Student, name string) {
	count := len(*department)
	for i, v := range Students {
		if count == acceptedNumber {
			break
		} else if v.secondChoice == name && v.assigned != true {
			v.assigned = true
			Students[i].assigned = true
			*department = append(*department, v)
			count++
		}
	}
}

func thirdRound(department *[]Student, name string) {
	count := len(*department)
	for i, v := range Students {
		if count == acceptedNumber {
			break
		} else if v.thirdChoice == name && v.assigned != true {
			v.assigned = true
			Students[i].assigned = true
			*department = append(*department, v)
			count++
		}
	}
}

func printDep(dep []Student) {
	for _, v := range dep {
		fmt.Print(v.firstName, " ", v.lastname, " ", v.GPA, "\n")
	}
	fmt.Println()
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

	//first round
	Mathematics = make([]Student, 0)
	populateDep(&Mathematics, "Mathematics")
	Physics = make([]Student, 0)
	populateDep(&Physics, "Physics")
	Biotech = make([]Student, 0)
	populateDep(&Biotech, "Biotech")
	Chemistry = make([]Student, 0)
	populateDep(&Chemistry, "Chemistry")
	Engineering = make([]Student, 0)
	populateDep(&Engineering, "Engineering")

	secondRound(&Mathematics, "Mathematics")
	secondRound(&Physics, "Physics")
	secondRound(&Biotech, "Biotech")
	secondRound(&Chemistry, "Chemistry")
	secondRound(&Engineering, "Engineering")

	thirdRound(&Mathematics, "Mathematics")
	thirdRound(&Physics, "Physics")
	thirdRound(&Biotech, "Biotech")
	thirdRound(&Chemistry, "Chemistry")
	thirdRound(&Engineering, "Engineering")

	fmt.Println("Biotech")
	printDep(Biotech)
	fmt.Println("Chemistry")
	printDep(Chemistry)
	fmt.Println("Engineering")
	printDep(Engineering)
	fmt.Println("Mathematics")
	printDep(Mathematics)
	fmt.Println("Physics")
	printDep(Physics)
}

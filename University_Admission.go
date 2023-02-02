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

var acceptedNumber int

var Students []Student

var Mathematics []Student

type Student struct {
	firstName    string
	lastname     string
	GPA          float64
	firstChoice  string
	secondChoice string
	thirdChoice  string
}

func newStudent(firstname string, lastname string, gpa float64, first string, second string, third string) Student {
	s := Student{firstName: firstname}
	s.lastname = lastname
	s.GPA = gpa
	s.firstChoice = first
	s.secondChoice = second
	s.thirdChoice = third
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

	_, err2 := fmt.Scan(&acceptedNumber)
	if err2 != nil {
		return
	}

	file, err := os.Open("applicant_list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		field := strings.Split(scanner.Text(), " ")

		var s Student

		s.firstName = field[0]
		s.lastname = field[1]
		s.GPA, err = strconv.ParseFloat(field[2], 64)
		s.firstChoice = field[3]
		s.secondChoice = field[4]
		s.thirdChoice = field[5]

		//fmt.Scan(fields[1], &s.lastname)
		//fmt.Scan(fields[2], &s.GPA)
		Students = append(Students, s)
	}

	sortStudents(Students)

	Mathematics = make([]Student, 0)

	count := 0
	for i, v := range Students {
		if count == acceptedNumber {
			break
		} else if v.firstChoice == "Mathematics" {
			Mathematics = append(Mathematics, v)
			Students = append(Students[:i], Students[i+1:]...) //removes the student from the Students slice
			count++
		}
	}
	fmt.Println(Mathematics)
	fmt.Println()
	fmt.Println(Students)
}

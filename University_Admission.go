package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var acceptedNumber int // the number of accepted students in each department

var Students []Student

type Student struct {
	firstName    string
	lastname     string
	physics      string
	chemistry    string
	math         string
	compScience  string
	firstChoice  string
	secondChoice string
	thirdChoice  string
	assigned     bool
}

type Department struct {
	name     string
	students []Student
}

func fileToSlice(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		field := strings.Split(scanner.Text(), " ")

		var s Student

		s.firstName = field[0]
		s.lastname = field[1]
		s.physics = field[2]
		s.chemistry = field[3]
		s.math = field[4]
		s.compScience = field[5]
		s.firstChoice = field[6]
		s.secondChoice = field[7]
		s.thirdChoice = field[8]

		Students = append(Students, s)
	}
}

//func sortStudents(students []Student) []Student {
//	sort.Slice(students, func(i, j int) bool {
//		if students[i].GPA != students[j].GPA {
//			return students[i].GPA > students[j].GPA
//		} else if students[i].firstName != students[j].firstName {
//			return students[i].firstName < students[j].firstName
//		}
//		return students[i].lastname < students[j].lastname
//	})
//	return students
//}

//func secondRound(department *[]Student, name string) {
//	count := len(*department)
//	for i, v := range Students {
//		if count == acceptedNumber {
//			break
//		} else if v.secondChoice == name && v.assigned != true {
//			v.assigned = true
//			Students[i].assigned = true
//			*department = append(*department, v)
//			count++
//		}
//	}
//}

//func thirdRound(department *[]Student, name string) {
//	count := len(*department)
//	for i, v := range Students {
//		if count == acceptedNumber {
//			break
//		} else if v.thirdChoice == name && v.assigned != true {
//			v.assigned = true
//			Students[i].assigned = true
//			*department = append(*department, v)
//			count++
//		}
//	}
//}

//func printDep(dep []Student) {
//	for _, v := range dep {
//		fmt.Printf("%v %v %.2f \n", v.firstName, v.lastname, v.GPA)
//	}
//	fmt.Println()
//}

func sortForDep(dep Department, stud []Student) []Student {
	sort.Slice(stud, func(i, j int) bool {
		var x, y string
		if dep.name == "Physics" {
			x = stud[i].physics
			y = stud[j].physics
		} else if dep.name == "Biotech" {
			x = stud[i].chemistry
			y = stud[j].chemistry
		} else if dep.name == "Chemistry" {
			x = stud[i].chemistry
			y = stud[j].chemistry
		} else if dep.name == "Mathematics" {
			x = stud[i].math
			y = stud[j].math
		} else if dep.name == "Engineering" {
			x = stud[i].compScience
			y = stud[j].compScience
		}
		if x != y {
			return x > y
		} else if stud[i].firstName != stud[j].firstName {
			return stud[i].firstName < stud[j].firstName
		}
		return stud[i].lastname < stud[j].lastname
	})
	return stud
}

func firstRound(department *[]Student, name string) {
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

	Biotech := Department{
		name:     "Biotech",
		students: make([]Student, 0),
	}
	Chemistry := Department{
		name:     "Chemistry",
		students: make([]Student, 0),
	}
	Engineering := Department{
		name:     "Engineering",
		students: make([]Student, 0),
	}
	Mathematics := Department{
		name:     "Mathematics",
		students: make([]Student, 0),
	}
	Physics := Department{
		name:     "Physics",
		students: make([]Student, 0),
	}

	sortForDep(Biotech, Students)
	firstRound(&Biotech.students, "Biotech")
	sortForDep(Chemistry, Students)
	firstRound(&Chemistry.students, "Chemistry")
	sortForDep(Engineering, Students)
	firstRound(&Engineering.students, "Engineering")
	sortForDep(Mathematics, Students)
	firstRound(&Mathematics.students, "Mathematics")
	sortForDep(Physics, Students)
	firstRound(&Physics.students, "Physics")

	fmt.Print(Biotech)
	fmt.Println()
	fmt.Print(Chemistry)
	fmt.Println()
	fmt.Print(Engineering)
	fmt.Println()
	fmt.Print(Mathematics)
	fmt.Println()
	fmt.Print(Physics)
	fmt.Println()
	fmt.Println(Students)
}

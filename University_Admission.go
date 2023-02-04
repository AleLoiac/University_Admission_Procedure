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

type Student struct {
	firstName       string
	lastname        string
	physics         string
	chemistry       string
	math            string
	compScience     string
	exam            string
	firstChoice     string
	secondChoice    string
	thirdChoice     string
	assigned        bool
	meanPhysics     string
	meanEngineering string
	meanBiotech     string
}

type Department struct {
	name     string
	students []Student
}

func meanPhy(student Student) string {
	x, _ := strconv.ParseFloat(student.physics, 64)
	y, _ := strconv.ParseFloat(student.math, 64)
	meanValue := strconv.FormatFloat((x+y)/2, 'f', 1, 64)
	return meanValue
}

func meanEng(student Student) string {
	x, _ := strconv.ParseFloat(student.compScience, 64)
	y, _ := strconv.ParseFloat(student.math, 64)
	meanValue := strconv.FormatFloat((x+y)/2, 'f', 1, 64)
	return meanValue
}

func meanBio(student Student) string {
	x, _ := strconv.ParseFloat(student.chemistry, 64)
	y, _ := strconv.ParseFloat(student.physics, 64)
	meanValue := strconv.FormatFloat((x+y)/2, 'f', 1, 64)
	return meanValue
}

func fileToSlice(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		field := strings.Split(scanner.Text(), " ")

		var s Student

		s.firstName = field[0]
		s.lastname = field[1]
		//phyx, _ := strconv.ParseInt(field[2], 10, 64)
		//phy := strconv.FormatFloat(float64(phyx), 'f', 1, 64)
		s.physics = field[2] + ".0"
		//chex, _ := strconv.ParseInt(field[3], 10, 64)
		//che := strconv.FormatFloat(float64(chex), 'f', 1, 64)
		s.chemistry = field[3] + ".0"
		//matx, _ := strconv.ParseInt(field[4], 10, 64)
		//mat := strconv.FormatFloat(float64(matx), 'f', 1, 64)
		s.math = field[4] + ".0"
		//comx, _ := strconv.ParseInt(field[5], 10, 64)
		//com := strconv.FormatFloat(float64(comx), 'f', 1, 64)
		s.compScience = field[5] + ".0"
		//exax, _ := strconv.ParseInt(field[6], 10, 64)
		//exa := strconv.FormatFloat(float64(exax), 'f', 1, 64)
		s.exam = field[6] + ".0"
		s.firstChoice = field[7]
		s.secondChoice = field[8]
		s.thirdChoice = field[9]
		s.meanPhysics = meanPhy(s)
		s.meanEngineering = meanEng(s)
		s.meanBiotech = meanBio(s)

		Students = append(Students, s)
	}

	for i, v := range Students {
		if v.physics < v.exam {
			Students[i].physics = Students[i].exam
		}
		if v.chemistry < v.exam {
			Students[i].chemistry = Students[i].exam
		}
		if v.math < v.exam {
			Students[i].math = Students[i].exam
		}
		if v.compScience < v.exam {
			Students[i].physics = Students[i].exam
		}
		if v.meanPhysics < v.exam {
			Students[i].meanPhysics = Students[i].exam
		}
		if v.meanEngineering < v.exam {
			Students[i].meanEngineering = Students[i].exam
		}
		if v.meanBiotech < v.exam {
			Students[i].meanBiotech = Students[i].exam
		}
	}
}

func sortForDep(dep Department, stud []Student) []Student {
	sort.Slice(stud, func(i, j int) bool {
		var x, y string
		if dep.name == "Physics" {
			x = stud[i].meanPhysics
			y = stud[j].meanPhysics
		} else if dep.name == "Biotech" {
			x = stud[i].meanBiotech
			y = stud[j].meanBiotech
		} else if dep.name == "Chemistry" {
			x = stud[i].chemistry
			y = stud[j].chemistry
		} else if dep.name == "Mathematics" {
			x = stud[i].math
			y = stud[j].math
		} else if dep.name == "Engineering" {
			x = stud[i].meanEngineering
			y = stud[j].meanEngineering
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

func printDep(stud []Student, dep Department) {
	var x float64
	for _, v := range stud {
		switch {
		case dep.name == "Biotech":
			x, _ = strconv.ParseFloat(v.meanBiotech, 64)
		case dep.name == "Physics":
			x, _ = strconv.ParseFloat(v.meanPhysics, 64)
		case dep.name == "Chemistry":
			x, _ = strconv.ParseFloat(v.chemistry, 64)
		case dep.name == "Mathematics":
			x, _ = strconv.ParseFloat(v.math, 64)
		case dep.name == "Engineering":
			x, _ = strconv.ParseFloat(v.meanEngineering, 64)
		}
		fmt.Printf("%v %v %.1f \n", v.firstName, v.lastname, x)
	}
	fmt.Println()
}

func main() {

	_, err := fmt.Scan(&acceptedNumber)
	if err != nil {
		return
	}

	file, err2 := os.Open("applicant_list.txt")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer func(file *os.File) {
		err3 := file.Close()
		if err3 != nil {

		}
	}(file)

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

	fmt.Println(Students)

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

	sortForDep(Biotech, Students)
	secondRound(&Biotech.students, "Biotech")
	sortForDep(Chemistry, Students)
	secondRound(&Chemistry.students, "Chemistry")
	sortForDep(Engineering, Students)
	secondRound(&Engineering.students, "Engineering")
	sortForDep(Mathematics, Students)
	secondRound(&Mathematics.students, "Mathematics")
	sortForDep(Physics, Students)
	secondRound(&Physics.students, "Physics")

	sortForDep(Biotech, Students)
	thirdRound(&Biotech.students, "Biotech")
	sortForDep(Chemistry, Students)
	thirdRound(&Chemistry.students, "Chemistry")
	sortForDep(Engineering, Students)
	thirdRound(&Engineering.students, "Engineering")
	sortForDep(Mathematics, Students)
	thirdRound(&Mathematics.students, "Mathematics")
	sortForDep(Physics, Students)
	thirdRound(&Physics.students, "Physics")

	//fmt.Println("Biotech")
	sortForDep(Biotech, Biotech.students)
	//printDep(Biotech.students, Biotech)
	//fmt.Println("Chemistry")
	sortForDep(Chemistry, Chemistry.students)
	//printDep(Chemistry.students, Chemistry)
	//fmt.Println("Engineering")
	sortForDep(Engineering, Engineering.students)
	//printDep(Engineering.students, Engineering)
	//fmt.Println("Mathematics")
	sortForDep(Mathematics, Mathematics.students)
	//printDep(Mathematics.students, Mathematics)
	//fmt.Println("Physics")
	sortForDep(Physics, Physics.students)
	//printDep(Physics.students, Physics)

	fileBiotech, err4 := os.Create("biotech.txt")
	if err4 != nil {
		log.Fatal(err4)
	}
	defer fileBiotech.Close()

	for _, line := range Biotech.students {
		_, err4 = fmt.Fprintln(fileBiotech, line.firstName, line.lastname, line.meanBiotech) // writes each line of the 'data' slice of strings
		if err4 != nil {
			log.Fatal(err4)
		}
	}

	fileChemistry, err5 := os.Create("chemistry.txt")
	if err5 != nil {
		log.Fatal(err5)
	}
	defer fileChemistry.Close()

	for _, line := range Chemistry.students {
		_, err5 = fmt.Fprintln(fileChemistry, line.firstName, line.lastname, line.chemistry) // writes each line of the 'data' slice of strings
		if err5 != nil {
			log.Fatal(err5)
		}
	}

	fileEngineering, err6 := os.Create("engineering.txt")
	if err6 != nil {
		log.Fatal(err6)
	}
	defer fileEngineering.Close()

	for _, line := range Engineering.students {
		_, err6 = fmt.Fprintln(fileEngineering, line.firstName, line.lastname, line.meanEngineering) // writes each line of the 'data' slice of strings
		if err6 != nil {
			log.Fatal(err6)
		}
	}

	fileMathematics, err7 := os.Create("mathematics.txt")
	if err7 != nil {
		log.Fatal(err7)
	}
	defer fileMathematics.Close()

	for _, line := range Mathematics.students {
		_, err7 = fmt.Fprintln(fileMathematics, line.firstName, line.lastname, line.math) // writes each line of the 'data' slice of strings
		if err7 != nil {
			log.Fatal(err7)
		}
	}

	filePhysics, err8 := os.Create("physics.txt")
	if err8 != nil {
		log.Fatal(err8)
	}
	defer filePhysics.Close()

	for _, line := range Physics.students {
		_, err8 = fmt.Fprintln(filePhysics, line.firstName, line.lastname, line.meanPhysics) // writes each line of the 'data' slice of strings
		if err8 != nil {
			log.Fatal(err8)
		}
	}
}

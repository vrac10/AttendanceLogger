package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

func getStudentInfo() (name, age, course string) {
	fmt.Println("Enter the student name:")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the student age:")

	inputReader = bufio.NewReader(os.Stdin)
	age, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the course:")

	inputReader = bufio.NewReader(os.Stdin)
	course, _ = inputReader.ReadString('\n')

	return name, age, course
}

func main() {
	name, age, course := getStudentInfo()

	file, err := os.Create("attendance.txt")
	if err != nil {
		errorHandler(err)
		log.Fatalf("%s", err)
	}

	defer file.Close()

	_, err2 := file.WriteString(name)
	_, err3 := file.WriteString(age)
	_, err4 := file.WriteString(course)

	if err2 != nil {
		errorHandler(err2)
		log.Fatalf("%s", err2)
	}

	if err3 != nil {
		errorHandler(err3)
		log.Fatalf("%s", err3)
	}

	if err4 != nil {
		errorHandler(err4)
		log.Fatalf("%s", err4)
	}

}

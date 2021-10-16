package main

//Import packages
import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//Funtion for logging errors
func errorHandler(err error) {
	println("Ops, something went wrong:", err)
}

//Function to get attendance


	

func getStudentInfo() (normtime, epochtime, name, roll, course string) {
	now:= time.Now()
	fmt.Println("Time: ", now.Local(), "\n")
	epoch:=now.Unix()
	norm:=now.Local()
	epochtime=fmt.Sprint(epoch)
	normtime=fmt.Sprint(norm)
	fmt.Println("Enter the student name:")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the student roll number:")

	inputReader = bufio.NewReader(os.Stdin)
	roll, _ = inputReader.ReadString('\n')

	fmt.Println("Enter the course:")

	inputReader = bufio.NewReader(os.Stdin)
	course, _ = inputReader.ReadString('\n')

	return strings.TrimSpace(normtime), strings.TrimSpace(epochtime), strings.TrimSpace(name), strings.TrimSpace(roll), strings.TrimSpace(course)
}

//Main
func main() {
	 normtime, epochtime, name, roll, course := getStudentInfo()
	record := []string{normtime, epochtime, name, roll, course}

	file, err := os.Create("attendance.txt")
	if err != nil {
		errorHandler(err)
		log.Fatalf("%s", err)
	}

	w := csv.NewWriter(file)

	defer file.Close()

	w.Write(record)
	w.Flush()
	err = w.Error()

	if err != nil {
		errorHandler(err)
		log.Fatalf("%s", err)
	}
}

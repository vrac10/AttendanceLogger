package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Enter the student name:")

	inputReader := bufio.NewReader(os.Stdin)
	name, _ := inputReader.ReadString('\n')

	file, err := os.Create("attendance.txt")
	if err != nil {
		//TODO Add error message here
		log.Fatalf("%s", err)
	}

	defer file.Close()

	_, err2 := file.WriteString(name)

	if err2 != nil {
		//TODO Add error message here
		log.Fatalf("%s", err)
	}

}

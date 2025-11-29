package project3

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"strings"
)

type students3 struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

var student []students3

func saveStudents(filename string) error {
	data, err := json.MarshalIndent(student, "", " ")
	if err != nil {
		return nil
	}
	return os.WriteFile(filename, data, 0644)
}
func loadstudents(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &student)
}
func Addstudents(name string, age int, grade string) {
	students3 := students3{Name: name, Age: age, Grade: grade}
	student = append(student, students3)
}
func Liststudents() {
	for _, s := range student {
		fmt.Printf("Name :%s , Age :%d , Grade :%s\n", s.Name, s.Age, s.Grade)
	}
}

func Projects3() {
	loadstudents("students.json")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("__students database")
		fmt.Println("1.Add students")
		fmt.Println("2.List students")
		fmt.Println("3.save and exit")

		fmt.Println("Enter choice:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Printf("Enter name:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Printf("Enter age:")
			ages, _ := reader.ReadString('\n')
			ages = strings.TrimSpace(ages)
			age, _ := strconv.Atoi(ages)

			fmt.Printf("Enter grade:")
			grade, _ := reader.ReadString('\n')
			grade = strings.TrimSpace(grade)

			Addstudents(name, age, grade)
			fmt.Println("Students added succesfully.")

		case "2":
			fmt.Println("All students")
			Liststudents()

		case "3":
			saveStudents("students.json")
			fmt.Println("saved succesfully")
			return
		default:
			fmt.Println("Invalid choice please enter again.")

		}
	}

}

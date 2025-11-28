package project

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

var students []student

func saveStudents(filename string) error {
	data, err := json.MarshalIndent(students, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
func loadstudents(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &students)
}
func Addstudents(name string, age int, grade string) {
	student := student{Name: name, Age: age, Grade: grade}
	students = append(students, student)
}
func Liststudents() {
	for _, s := range students {
		fmt.Printf("Name :%s,Age :%d,Grade :%s\n", s.Name, s.Age, s.Grade)
	}
}
func StudentProject() {
	loadstudents("students.json")

	Addstudents("Akash", 20, "A")
	Addstudents("Bunty", 20, "B")

	fmt.Println("All students:")
	Liststudents()

	saveStudents("students.json")
}

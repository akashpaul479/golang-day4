package day4

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Branch string `json:"branch"`
}

func Json1() {
	file, err := os.Open("students.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	value, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var students []student
	err = json.Unmarshal(value, &students)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("All students")
	for _, e := range students {
		fmt.Printf("Name: %s | Age: %d | Branch: %s\n", e.Name, e.Age, e.Branch)
	}

}

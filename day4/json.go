package day4

import (
	"encoding/json"
	"fmt"
)

type students struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Branch string `json:"branch"`
}

func Json() {
	students := []students{
		{Name: "akash", Age: 20, Branch: "cse"},
		{Name: "abhi", Age: 20, Branch: "cse"},
		{Name: "bunty", Age: 20, Branch: "mpcs"},
		{Name: "kushal", Age: 21, Branch: "mpcs"},
		{Name: "tanmay", Age: 20, Branch: "mpcs"},
	}
	jsondata, err := json.MarshalIndent(students, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("person data is :", string(jsondata))
}

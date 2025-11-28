package day4

import (
	"encoding/json"
	"fmt"
	"os"
)

type students1 struct {
	Name   string
	Age    int
	Branch string
}

func Json2() {
	s := []students1{
		{Name: "Akash", Age: 20, Branch: "cse"},
		{Name: "abhi", Age: 20, Branch: "cse"},
		{Name: "bunty", Age: 20, Branch: "mpcs"},
		{Name: "kushal", Age: 21, Branch: "mpcs"},
		{Name: "tanmay", Age: 20, Branch: "mpcs"},
	}
	data, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = os.WriteFile("students.json", data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("students json created succesfully.")
}

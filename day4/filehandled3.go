package day4

import (
	"fmt"
	"os"
)

func Filehandled3() {
	file, err := os.OpenFile("demo.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	name := "\nAKASH"
	_, errors := file.WriteString(name)
	if errors != nil {
		fmt.Println("Error:", errors)
		return
	}

	file.WriteString("name added")
}

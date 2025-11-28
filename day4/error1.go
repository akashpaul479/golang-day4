package day4

import (
	"fmt"
	"os"
)

func Error1() {
	_, err := os.Open("Fake file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("File opened sucessfully")
}

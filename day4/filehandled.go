package day4

import (
	"fmt"
	"io"
	"os"
)

func Filehandled() {
	file, err := os.Create("demo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	content := "hey!, my name is akash\n" +
		"iam from telangana\n" +
		"iam a sports person\n" +
		"my fav sport is football\n"

	_, errors := io.WriteString(file, content)
	if errors != nil {
		fmt.Println("Error:", errors)
		return
	}
	fmt.Println("Succesfully created.")
}

package day4

import (
	"fmt"
	"io"
	"os"
)

func Filehandled1() {

	file, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}
}

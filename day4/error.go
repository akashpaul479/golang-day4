package day4

import (
	"fmt"
	"strconv"
)

func Error() {
	value, err := strconv.Atoi("hello")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Converted:", value)
}

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Input your file: ")
	var input string
	fmt.Scanln(&input)

	f, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully")
}

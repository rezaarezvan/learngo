package main

import "fmt"

func main() {
	m := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	for i := 1; i <= 100; i++ {
		output := ""

		for k, v := range m {
			if i%k == 0 {
				output += v
			}
		}

		if output != "" {
			fmt.Println(output)
		}
	}
}

package main

import "fmt"

func main() {
	addM := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}

	addFive := addM(5)
	result := addFive(6)

	fmt.Println(result)
}

package main

import "fmt"

func main() {
	result := calc(1, 2)
	fmt.Println(result)
}

func calc(a, b float64) float64 {
	return a + b
}

package main

import "fmt"

func Sum(a, b int) (result int) {
	result = a + b
	return result
}

func main() {
	result := Sum(5, 5)
	fmt.Println(result)
}

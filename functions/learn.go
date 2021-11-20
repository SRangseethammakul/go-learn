package functions

import (
	"fmt"
	"strings"
)

// var FullName string
func add(x, y int) int {
	return x + y
}
func convert(title, email string) (string, string) {
	s1 := strings.ToUpper(title)
	return s1, email
}
func sum(numbers ...int) { // same destructuring in js
	// fmt.Println(numbers) // Array
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
}
func Learn() {
	fmt.Println(add(5, 15))
	fmt.Println(convert("Hello", "World@gmail.com"))
	sum(10, 20, 30, 40, 50)
}

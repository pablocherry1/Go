package main

import (
	"flag"
	"fmt"
)

func main() {

	num1 := flag.Float64("num1", 0, "The first number")

	num2 := flag.Float64("num2", 0, "The second number")

	op := flag.String("op", "+", "The operator to use (e.g., +, -, *, /)")

	flag.Parse()

var result float64

switch *op {
case "+":
	result = *num1 + *num2
case "-":
	result = *num1 - *num2
case "*":
	result = *num1 * *num2
case "/":
	if *num2 == 0 {
		fmt.Println("Error: Cannot divide by zero!")
		return
	}
	result = *num1 / *num2
default:
	fmt.Printf("Error: Invalid operator '%s' \n", *op)
	return
}

fmt.Printf("The result is: %.2f \n", result)

}
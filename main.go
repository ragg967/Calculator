package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ragg967/GoCalculator/calculate" // Updated import path
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <num1> <operator> <num2>")
		return
	}

	num1Str := os.Args[1]
	operator := os.Args[2]
	num2Str := os.Args[3]

	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("Error parsing num1:", err)
		return
	}

	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("Error parsing num2:", err)
		return
	}

	var result float64

	switch operator {
	case "+":
		result = calculate.Add(num1, num2) // Updated package name
	case "-":
		result = calculate.Subtract(num1, num2) // Updated package name
	case "x":
		result = calculate.Multiply(num1, num2) // Updated package name
	case "/":
		result, err = calculate.Divide(num1, num2) // Updated package name
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	default:
		fmt.Println("Unknown operator:", operator)
		return
	}

	fmt.Printf("Result: %f\n", result)
}

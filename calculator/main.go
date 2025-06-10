package main

import (
    "fmt"        // For input/output
    "strconv"    // For string to number conversion
    "bufio"      // For reading user input
    "os"         // For os.Stdin
    "strings"    // For string manipulation
)

func main() {
	fmt.Println("Simple Go Calculator")
	fmt.Println("Enter expression (e.g., 5 + 3, 10 / 2). Type 'exit' to quit.")

	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ") // Prompt for input
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator. Goodbye!")
			break
		}

		num1, operator, num2, err := parseExpression(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result, err := calculate(num1, operator, num2)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("Result: %.2f\n", result) // Format to 2 decimal places
	}
}

// parseExpression takes a string expression and returns the two numbers and the operator.
func parseExpression(expr string) (float64, string, float64, error) {
	parts := strings.Fields(expr)
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid expression format. Expected: number operator number")
	}

	num1Str := parts[0]
	operator := parts[1]
	num2Str := parts[2]

	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid first number: %w", err)
	}

	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid second number: %w", err)
	}

	validOperators := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	if !validOperators[operator] {
		return 0, "", 0, fmt.Errorf("invalid operator '%s'. Supported operators: +, -, *, /", operator)
	}

	return num1, operator, num2, nil
}

// calculate performs the arithmetic operation based on the operator.
func calculate(num1 float64, operator string, num2 float64) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, fmt.Errorf("division by zero is not allowed")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", operator)
	}
}

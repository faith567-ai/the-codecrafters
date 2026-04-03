package main

import (
	"fmt"
	"strconv"
)

func main() {
	const (
		red    = "\033[0;31m"
		green  = "\033[0;32m"
		yellow = "\033[0;33m"
		blue   = "\033[0;34m"
		reset  = "\033[0m"
		purple = "\033[0;35m"
	)

start:
	fmt.Print(red + "FIRST VALUE: " + reset)
	var num1 string
	fmt.Scanln(&num1)

	switch num1 {
	case "quit", "QUIT":
		fmt.Println(blue + "GOODBYE !!" + reset)
		return

	case "help", "HELP":
		fmt.Println(green +"<num1> + <num2>   → addition" + reset)
		fmt.Println("<num1> - <num2>   → addition")
		fmt.Println(green +"<num1> * <num2>   → addition")
		fmt.Println("<num1> / <num2>   → addition")
		fmt.Println(" ")
		goto start
	}

	val1, err := strconv.Atoi(num1)
	if err != nil {
		fmt.Println("you cannot put alphabet in a calculator")
		fmt.Println("try again")
		goto start
	}

operate:

	fmt.Println("Available operators: (+, -, *, /)")
	fmt.Print("OPERATOR: ")
	var operator string
	fmt.Scanln(&operator)

	switch operator {
	case "quit", "QUIT":
		fmt.Println("GOODBYE !!")
		return

	case "help", "HELP":
		fmt.Println("<num1> + <num2>   → addition")
		fmt.Println("<num1> - <num2>   → addition")
		fmt.Println("<num1> * <num2>   → addition")
		fmt.Println("<num1> / <num2>   → addition")
		fmt.Println(" ")
		goto operate
	}

second:
	fmt.Print("SECOND VALUE: ")
	var num2 string
	fmt.Scanln(&num2)

	switch num2 {
	case "quit", "QUIT":
		fmt.Println("GOODBYE !!")
		return

	case "help", "HELP":
		fmt.Println("<num1> + <num2>   → addition")
		fmt.Println("<num1> - <num2>   → addition")
		fmt.Println("<num1> * <num2>   → addition")
		fmt.Println("<num1> / <num2>   → addition")
		fmt.Println(" ")
		goto second
	}

	val2, err := strconv.Atoi(num2)
	if err != nil {
		fmt.Println("possibly missing comma")
		fmt.Println("try again")
		goto second
	}

	switch operator {
	case "+":
		fmt.Printf("RESULT: %v %s %v = %v\n", val1, operator, val2, val1+val2)

	case "-":
		fmt.Printf("RESULT: %v %s %v = %v\n", val1, operator, val2, val1-val2)

	case "*":
		fmt.Printf("RESULT: %v %s %v = %v\n", val1, operator, val2, val1*val2)

	case "/":
		if val2 == 0 {
			fmt.Println("Can't divide by zero")
			goto start
		}
		fmt.Printf("RESULT: %v %s %v = %v\n", val1, operator, val2, val1/val2)

	default:
		fmt.Println("invalid operator")
		fmt.Println("try again")
		goto start
	}

	goto start
}

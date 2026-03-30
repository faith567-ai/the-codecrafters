package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	start:
	fmt.Print("INPUT:")
	var num string
	fmt.Scanln(&num)

	switch num {
	case "quit","QUIT":
		fmt.Println("BYE!!")
		return
	}
	fmt.Println("Available conversion: <hex>, <bin>, <dec>")
	fmt.Println("Which convert do you want to apply:")
	var converter string
	fmt.Scanln(&converter)

	switch converter {
	case "hex", "HEX":
		val, err := strconv.ParseInt(num,16,64)

		if err != nil {
			fmt.Println("value is not hexadecimal num")
			fmt.Println("try again")
			fmt.Println(" ")
			goto start
		}

		fmt.Printf("decimal : %v\n", val)

	case "bin", "BIN":
		val, err := strconv.ParseInt(num, 2, 64)

		if err != nil {
			fmt.Println("value is not a binary num")
			fmt.Println("try again")
			fmt.Println(" ")
			goto start
			

		}

		fmt.Printf("binary : %v\n", val)

	case "dec", "DEC":
		val, err := strconv.ParseInt(num, 10, 64)

		if err != nil{
			fmt.Println("value is not a decimal num")
			fmt.Println("try again")
			fmt.Println(" ")
			goto start
			
		}
		binary := strconv.FormatInt(val, 2)
		hex := strings.ToUpper(strconv.FormatInt(val, 16))

		fmt.Printf("binary : %v\n", binary)
		fmt.Printf("hex : %v\n", hex)

	default:
		fmt.Println("invalid operator")
		fmt.Println("try again")
		goto start
	}
	goto start
}
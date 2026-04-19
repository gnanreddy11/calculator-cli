package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func Addition() {
	
}

func main() {
	var opt string
	var n1 string
	var n2 string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("              **** calculator-cli ****              ")
		fmt.Println()
		fmt.Println("please pick any 4 arithmetic operations given below:")
		fmt.Println()
		fmt.Println("1. addition (+)                       2. subtraction (-)")
		fmt.Println()
		fmt.Println("3. multiplication (x)                 4. division (/)")
		fmt.Println()
		fmt.Print("your option (enter q|Q for exit): ")
		scanner.Scan()
		opt = scanner.Text()
		switch opt {
			case "1":
				fmt.Println()
				fmt.Print("enter 1st number: ")
				scanner.Scan()
				n1 = scanner.Text()
				fmt.Print("enter 2nd number: ")
				scanner.Scan()
				n2 = scanner.Text()
				floatN1, err := strconv.ParseFloat(n1, 64)
				if err != nil {
					fmt.Println("invalid number. please try again!")
				}
				floatN2, err := strconv.ParseFloat(n2, 64)
				if err != nil {
					fmt.Println("invalid number. please try again!")
				}
				sum := floatN1 + floatN2
				fmt.Println()
				fmt.Printf("result: %.4f\n\n", sum)
			case "2":
				fmt.Println()
				fmt.Print("enter 1st number: ")
				scanner.Scan()
				n1 = scanner.Text()
				fmt.Print("enter 2nd number: ")
				scanner.Scan()
				n2 = scanner.Text()
				floatN1, err := strconv.ParseFloat(n1, 64)
				if err != nil {
					fmt.Println("invalid number. please try again!")
				}
				floatN2, err := strconv.ParseFloat(n2, 64)
				if err != nil {
					fmt.Println("invalid number. please try again!")
				}
				sub := floatN1 - floatN2
				fmt.Println()
				fmt.Printf("result: %.4f\n\n", sub)
			case "3":
				fmt.Println()
				fmt.Print("enter 1st number: ")
				scanner.Scan()
				n1 = scanner.Text()
				fmt.Print("enter 2nd number: ")
				scanner.Scan()
				n2 = scanner.Text()
				floatN1, err := strconv.ParseFloat(n1, 64)
				if err != nil {
					fmt.Println("invalid number. please try again!")
				}
				floatN2, err := strconv.ParseFloat(n2, 64)
				if err != nil {
					fmt.Println("invalid number. please try again!")
				}
				mul := floatN1 * floatN2
				fmt.Println()
				fmt.Printf("result: %.4f\n\n", mul)
			case "4":
				fmt.Println()
				fmt.Print("enter 1st number: ")
				scanner.Scan()
				n1 = scanner.Text()
				fmt.Print("enter 2nd number: ")
				scanner.Scan()
				n2 = scanner.Text()
				floatN1, err := strconv.ParseFloat(n1, 64)
				if err != nil {
					fmt.Println("invalid number. please try again!")
				}
				floatN2, err := strconv.ParseFloat(n2, 64)
				if err != nil {
					fmt.Println("invalid number. please try again!")
				}
				div := floatN1 / floatN2
				fmt.Println()
				fmt.Printf("result: %.4f\n\n", div)
			case "q", "Q":
				return
			default :
				fmt.Println()
				fmt.Println("invalid option. please try again!")
				fmt.Println()
		}
	}
}
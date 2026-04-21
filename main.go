package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getOptions(scanner *bufio.Scanner) (float64, float64, error) {
	fmt.Println()
	fmt.Print("enter 1st number: ")
	scanner.Scan()
	n1, err1 := strconv.ParseFloat(scanner.Text(), 64)
	fmt.Print("enter 2nd number: ")
	scanner.Scan()
	n2, err2 := strconv.ParseFloat(scanner.Text(), 64)
	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("invalid number. please try again!")
	}
	return n1, n2, nil
}

func main() {
	var opt string
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
				num1, num2, err := getOptions(scanner)
				sum := num1 + num2
				if err != nil {
					fmt.Println("error:", err)
					fmt.Println()
				} else {
					fmt.Println()
					fmt.Printf("result: %.4f\n\n", sum)
				}
			case "2":
				num1, num2, err := getOptions(scanner)
				diff := num1 - num2
				if err != nil {
					fmt.Println("error:", err)
					fmt.Println()
				} else {
					fmt.Println()
					fmt.Printf("result: %.4f\n\n", diff)
				}
			case "3":
				num1, num2, err := getOptions(scanner)
				mul := num1 * num2
				if err != nil {
					fmt.Println("error:", err)
					fmt.Println()
				} else {
					fmt.Println()
					fmt.Printf("result: %.4f\n\n", mul)
				}
			case "4":
				num1, num2, err := getOptions(scanner)
				if num2 == float64(0) {
					fmt.Println("denominator cannot be zero number. please try again!")
					fmt.Println()
				} else {
					div := num1 / num2
					if err != nil {
						fmt.Println("error:", err)
						fmt.Println()
					} else {
						fmt.Println()
						fmt.Printf("result: %.4f\n\n", div)
					}
				}
			case "q", "Q":
				return
			default :
				fmt.Println()
				fmt.Println("invalid option. please try again!")
				fmt.Println()
		}
	}
}
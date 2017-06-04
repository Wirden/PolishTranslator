package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var WRONG_SYNTHAX_ERROR = " Malformed expression"
var VALID_MESSAGE = " is equals to "
var ARGUMENT_NULL_ERROR = " An equation cannot be empty"
var EMPTY_LINE = " This line is empty"

func main() {

	arg := os.Args[1]

	if arg == "" {
		fmt.Print(arg, EMPTY_LINE)
		os.Exit(0)
	}

	var stack []float64
	for _, tok := range strings.Fields(arg) {
		switch tok {
		case "+":
			if len(stack) < 2 {
				error1(arg)
			} else {
				stack[len(stack)-2] += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "-":
			if len(stack) < 2 {
				error1(arg)
			} else {
				stack[len(stack)-2] -= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "*":
			if len(stack) < 2 {
				error1(arg)
			} else {
				stack[len(stack)-2] *= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "/":
			if len(stack) < 2 {
				error1(arg)
			} else {
				stack[len(stack)-2] /= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		default:
			if _, err := strconv.Atoi(tok); err == nil {
				f, _ := strconv.ParseFloat(tok, 64)
				stack = append(stack, f)
			} else {
				fmt.Print(arg, err)
				os.Exit(0)
			}
		}
	}

	if len(stack) > 1 {
		fmt.Print(arg, WRONG_SYNTHAX_ERROR)
		os.Exit(0)
	}

	fmt.Print(arg, VALID_MESSAGE, stack[0])
}

func error1(arg string) {
	fmt.Print(arg, WRONG_SYNTHAX_ERROR)
	os.Exit(0)
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var WRONG_SYNTHAX_ERROR = " Malformed expression"
var VALID_MESSAGE = " is equals to"
var ARGUMENT_NULL_ERROR = " An equation cannot be empty"
var EMPTY_LINE = " This line is empty"

func main() {
	start := time.Now()
	arg := os.Args[1]

	if arg == "" {
		elapsed := time.Since(start)
		fmt.Println(arg, EMPTY_LINE, ", ", elapsed)
		os.Exit(0)
	}

	var stack []float64
	for _, tok := range strings.Fields(arg) {
		switch tok {
		case "+":
			if len(stack) < 2 {
				error1(arg, start)
			} else {
				stack[len(stack)-2] += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "-":
			if len(stack) < 2 {
				error1(arg, start)
			} else {
				stack[len(stack)-2] -= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "*":
			if len(stack) < 2 {
				error1(arg, start)
			} else {
				stack[len(stack)-2] *= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "/":
			if len(stack) < 2 {
				error1(arg, start)
			} else {
				stack[len(stack)-2] /= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		default:
			if _, err := strconv.Atoi(tok); err == nil {
				f, _ := strconv.ParseFloat(tok, 64)
				stack = append(stack, f)
			} else {
				elapsed := time.Since(start)
				fmt.Println(arg, err, ", ", elapsed)
				os.Exit(0)
			}
		}
	}

	if len(stack) > 1 {
		elapsed := time.Since(start)
		fmt.Println(arg, WRONG_SYNTHAX_ERROR, ", ", elapsed)
		os.Exit(0)
	}

	elapsed := time.Since(start)
	fmt.Println(arg, VALID_MESSAGE, stack[0], ", ", elapsed)
}

func error1(arg string, start time.Time) {
	elapsed := time.Since(start)
	fmt.Println(arg, WRONG_SYNTHAX_ERROR, ", ", elapsed)
	os.Exit(0)
}

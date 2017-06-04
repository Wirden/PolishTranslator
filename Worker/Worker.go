package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var WRONG_SYNTHAX_ERROR = "Malformed notation ( Only Integers and the 4 basics operators can be used )"
var VALID_MESSAGE = " is equals to "
var ARGUMENT_NULL_ERROR = "An equation cannot be empty"

func main() {

	l := log.New(os.Stderr, "", 0)

	if len(os.Args) != 2 {
		l.Fatal(ARGUMENT_NULL_ERROR)
	}

	arg := os.Args[1]

	var stack []float64
	for _, tok := range strings.Fields(arg) {
		switch tok {
		case "+":
			if len(stack) < 2 {
				error1()
			} else {
				stack[len(stack)-2] += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "-":
			if len(stack) < 2 {
				error1()
			} else {
				stack[len(stack)-2] -= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "*":
			if len(stack) < 2 {
				error1()
			} else {
				stack[len(stack)-2] *= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		case "/":
			if len(stack) < 2 {
				error1()
			} else {
				stack[len(stack)-2] /= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		default:
			if _, err := strconv.Atoi(tok); err == nil {
				f, _ := strconv.ParseFloat(tok, 64)
				stack = append(stack, f)
			} else {
				l.Fatal(err)
			}
		}
	}

	fmt.Println(arg, VALID_MESSAGE, stack[0])
}

func error1() {
	l := log.New(os.Stderr, "", 0)
	err := errors.New(WRONG_SYNTHAX_ERROR)
	l.Fatal(err)
}

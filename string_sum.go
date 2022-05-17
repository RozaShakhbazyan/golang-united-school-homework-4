package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var a int
	var b int

	input = strings.TrimSpace(input)

	if input == "" {
		return "", errorEmptyInput
	}

	for i := 0; i < len(input); i++ {
		if i == 0 && input[0] == '-' {

			continue

		}

		if isOperand(rune(input[i])) {
			if isOperand(rune(input[(i+1)%len(input)])) {
				return "", fmt.Errorf("%s", errorNotTwoOperands)

			}

			a, err = strconv.Atoi(strings.TrimSpace(input[:i]))
			if err != nil {
				return "", fmt.Errorf("%s", errorNotTwoOperands)
			}
			b, err = strconv.Atoi(strings.TrimSpace(input[i+1:]))
			if err != nil {
				return "", fmt.Errorf("%s", errorNotTwoOperands)
			}

			if input[i] == '+' {
				return strconv.Itoa(a + b), nil
			} else {
				return strconv.Itoa(a - b), nil
			}
		}
	}
	return "", fmt.Errorf("%s", errorNotTwoOperands)
}
func isOperand(input rune) bool {
	return input == '+' || input == '-'
}

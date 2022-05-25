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
	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		err = fmt.Errorf("%w", errorEmptyInput)
		return
	}

	number := ""

	if input[0] == '-' || input[0] == '+' {
		number += string(input[0])
		input = input[1:]
	}

	rightBound := strings.Index(input, "+")
	if rightBound == -1 {
		rightBound = strings.Index(input, "-")
	}

	if rightBound != -1 {
		number += input[:rightBound]
		input = input[rightBound:]
	} else {
		err = fmt.Errorf("%w", errorNotTwoOperands)
		return
	}

	signsCount := strings.Count(input, "+") + strings.Count(input, "-")

	if signsCount != 1 {
		err = fmt.Errorf("%w", errorNotTwoOperands)
		return
	}

	a, err := strconv.Atoi(number)

	if err != nil {
		err = &strconv.NumError{Func: "Atoi", Num: number, Err: err}
		return
	}

	b, err := strconv.Atoi(input)

	if err != nil {
		err = &strconv.NumError{Func: "Atoi", Num: input, Err: err}
		return
	}

	return strconv.Itoa(a + b), nil
}

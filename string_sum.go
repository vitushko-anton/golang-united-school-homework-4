package string_sum

import (
	"errors"
	"fmt"
	"regexp"
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
	wraErr := func(err error) (string, error) {
		return "", fmt.Errorf("an error occurred:%w", err)
	}
	input = strings.TrimSpace(input)
	if len(input) == 0 {

		return wraErr(errorEmptyInput)
	}

	x := strings.TrimLeft(input, "-+")
	x = strings.TrimRight(x, "-+")
	stringReg := regexp.MustCompile("[+-]{1}").Split(x, -1)
	if len(stringReg) == 1 || len(stringReg) > 2 {

		return wraErr(errorNotTwoOperands)
	}

	k := strings.LastIndexAny(input, "-+")
	n, err := strconv.Atoi(input[0:k])
	if err != nil {
		return wraErr(err)
	}
	m, err := strconv.Atoi(input[k:len(input)])
	if err != nil {
		return wraErr(err)
	}
	output = strconv.Itoa(n + m)
	return output, nil
}

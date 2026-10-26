package main

import (
	"errors"
	"strconv"
)

var ErrZero = errors.New("zero is not allowed")

func fizzBuzz(n int) (string, error) {
	if n == 0 {
		return "", ErrZero
	}

	isThreeMultiple := n%3 == 0
	isFiveMultiple := n%5 == 0

	switch {
	case isThreeMultiple && isFiveMultiple:
		return "FizzBuzz", nil
	case isThreeMultiple:
		return "Fizz", nil
	case isFiveMultiple:
		return "Buzz", nil
	default:
		return strconv.Itoa(n), nil
	}
}

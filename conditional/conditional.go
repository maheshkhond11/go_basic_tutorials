package conditional

import (
	"errors"
	"fmt"
)

func ConditionalBlocks() {
	var printValue = "Hello, World!"
	printMe(printValue)
	var numerator, denominator = 11, 2
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the ineteger division is %v", result)
	} else {
		fmt.Printf("the result of integere division is %v with remainder %v\n", result, remainder)
	}
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the ineteger division is %v", result)
	default:
		fmt.Printf("the result of integere division is %v with remainder %v", result, remainder)
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New(("cannot Divide by Zero"))
		return 0, 0, err
	}
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}

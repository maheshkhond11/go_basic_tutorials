package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func main() {
	//conditionalBlocks()
	//collections()
	//memoryAllocationSpeedTest()
	//stringsInGo()
	structs()
}

func structs() {

}

func stringsInGo() {
	var myString = []rune("resume")
	var indexed = myString[1]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\nthe length of 'myString' is %v", len(myString))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "s", "c"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)

	//string builder
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)
}

func memoryAllocationSpeedTest() {
	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func collections() {
	//arrays
	//var intArr [3]int32 = [3]int32{1, 2, 3}
	//intArr := [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	//slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	//maps
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])
	var age, ok = myMap2["Jason"]
	//delete(myMap2, "Adam")
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v Age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}
	var i int = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func conditionalBlocks() {
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

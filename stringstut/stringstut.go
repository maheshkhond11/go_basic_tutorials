package stringstut

import (
	"fmt"
	"strings"
)

func StringsInGo() {
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

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)
}

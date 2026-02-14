package pointers

import "fmt"

func Pointers() {
	var p *int32 = new(int32)
	var i int32
	*p = 10

	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	p = &i
	*p = 1
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
}

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
	//ownerInfo owner
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

// type owner struct {
// 	name string
// }

func main() {
	conditionalBlocks()
	collections()
	memoryAllocationSpeedTest()
	stringsInGo()
	structs()
	pointers()
	routines()
	channels()
	generics()
}

func generics() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(isEmpty(intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))
	fmt.Println(isEmpty(float32Slice))
}

func isEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func channels() {
	// var c = make(chan int)
	// go process(c)
	// for i := range c {
	// 	fmt.Println(i)
	// }
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nFound a deal on chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\nFound a deal on tofu at %v", website)
	}
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
}

// var m = sync.Mutex{}
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func routines() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		//go dbCall(i)
		go count()
	}
	wg.Wait()
	fmt.Printf("\ntotal execution time: %v", time.Since(t0))
	//fmt.Printf("\nThe results are %v", results)
}

func count() {
	var res int
	for i := 0; i < 100000000; i++ {
		res++
	}
	wg.Done()
}

func dbCall(i int) {

	//var delay float32 = rand.Float32() * 2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	//fmt.Println("\nThe result from the database is:", dbData[i])
	//save(dbData[i])
	//log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("\nthe current results are: %v", results)
	m.RUnlock()
}

func pointers() {
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

func structs() {
	//var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}}
	var myEngine gasEngine = gasEngine{25, 15}
	//myEngine.mpg = 20
	canMakeIt(myEngine, 50)
	fmt.Printf("Total miles left in tank: %v", myEngine.milesLeft())
	fmt.Println(myEngine.mpg, myEngine.gallons)
	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)
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

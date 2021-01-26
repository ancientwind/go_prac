package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const Name = "Main"

var Pi float64

func init() {
	Pi = 4 * math.Atan(1)
}

func main() {
	var goos = runtime.GOOS // You can omit the variable type from the declaration, when you are assigning a value to a variable at the time of declaration.
	fmt.Printf("The operating system is: %s \n", goos)
	path := os.Getenv("PATH")
	fmt.Println("Path is ", path)
	var a uint32 = 100 // you can change the variable type of the assigning value by specifying the type at declaration
	b := -200          // := is 赋值操作符，编辑器会自动推断类型
	fmt.Printf("a is: %d \t b is: %d", a, b)

	i, _, k := 29, 989, "hello" // _ 空白标识符， 抛弃一个值，比如只想活得某个函数的部分返回值
	fmt.Printf("i: %d, _, k: %s", i, k)

	fmt.Printf("Pi is %f", Pi)

	rawStr := `this is a raw string \n \t \\ \u`
	fmt.Println(rawStr, len(rawStr))

	// strings
	fmt.Println(strings.HasPrefix(rawStr, "th"),
		strings.Contains(rawStr, "raw"),
		strings.LastIndex(rawStr, "is"),
		strings.Replace(rawStr, "is", "IS", 1),
		len(strings.Fields(rawStr)),
		len(strings.Split(rawStr, "\\")))

	// strconv
	data, err := strconv.Atoi("2324")
	if err != nil {
		fmt.Println(data, "is not a valid number string")
		//fmt.Printf("Program stoped with error: %v", err)
		//os.Exit(1)
		return
	}
	fmt.Println(strconv.Itoa(32), data)

	// time
	t := time.Now()
	fmt.Println(t, t.Day(), t.Weekday())

	/**
	pointer
	1. 指针运算，比如*ptr++，在Go中是不被允许的
	2. Go的指针更像是一种引用，可以减少内存和提高效率
	*/
	myNum := 64
	fmt.Printf("myNum is %d and stress in mem is %p . \n", myNum, &myNum)
	var numPtr *int
	fmt.Printf("numPtr before set value: %p ; and is numPtr empty? %t \n", numPtr, numPtr == nil) // In Go, nil is the zero value for pointers, interfaces, maps, slices, channels and function types, representing an uninitialized value.
	numPtr = &myNum
	fmt.Printf("numPrt stress after set value: %p, and its referenced data is: %d \n", numPtr, *numPtr)

	hello := "hello"
	var hiPtr = &hello
	*hiPtr = "world"
	fmt.Println(*hiPtr)
	fmt.Println(hello)
	//var namePtr = &Name   // error, cannot get address of a literal or const
	//var tenPtr = &10

	/**
	control statement
	*/
	var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."
	if len(prompt) != 0 {
		fmt.Println(prompt)
	}

	/**
	switch
	*/
	var ab [2]int
	ab[0] = 23
	ab[1] = 86
	// define variables in expression
	switch a2, b2 := ab[0], ab[1]; {
	case a2 > b2:
		fmt.Printf("%v > %v", a2, b2)
	case a2 == b2:
		fmt.Printf("%v == %v", a2, b2)
	case a2 < b2:
		fmt.Printf("%v < %v", a2, b2)
	default:
		fmt.Println("N/A")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // use commas to separate multiple expr
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	//t := time.Now()
	switch { // no expr; as if/else
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	case t.Hour() > 12 && t.Hour() < 18:
		fmt.Println("it's afternoon")
	default:
		fmt.Println("it's evening")

	}

	/**
	for iteration
	 */
	for i, j := 0, 10; i < j; i, j = i + 1, j - 2 {
		fmt.Println(i, j)
	}

	gol := "Go is a beautiful language!"
	for ix := 0; ix < len(gol); ix++ {
		fmt.Printf("Character on position %d is %c \n", ix, gol[ix])
	}

	for i:=0; i<3; i++ {
		for j:=0; j<10; j++ {
			if j>5 {
				break // break will only jump out the nearest control
			}
			fmt.Print(j)
		}
		fmt.Print("  ")
	}

	/**
	LABEL and goto --- not recommend to use
	 */
    i = 0
	HERE:
		fmt.Print(i)
        i++
        if i == 5 {
			return}
		goto HERE
}

package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

const Name = "Main"

var Pi float64

func init()  {
	Pi = 4 * math.Atan(1)
}

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s \n", goos)
	path := os.Getenv("PATH")
	fmt.Println("Path is ", path)
	var a uint32= 100
	b := -200 // := is 赋值操作符，编辑器会自动推断类型
	fmt.Printf("a is: %d \t b is: %d", a, b)

	i, _, k := 29, 989, "hello" // _ 空白标识符， 抛弃一个值，比如只想活得某个函数的部分返回值
	fmt.Printf("i j k", i, k)

	fmt.Printf("Pi is %f", Pi)

}

package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var i int
var first, second int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func combine() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func main() {
	combine()
	fmt.Println(add(first, second))

	fmt.Println(i, c, python, java)
	fmt.Println(split(17))

	a, b := swap("Hello", "World")
	fmt.Println(a, b)

	fmt.Println("Welcome to the Playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favourite number is", rand.Intn(100))
	fmt.Println("Pi is ", math.Pi)
	fmt.Println("Sum for 5 and 6 is", add(5, 6))
}

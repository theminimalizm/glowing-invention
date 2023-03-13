package main

import (
	"fmt"
	"math"
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

func main() {
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

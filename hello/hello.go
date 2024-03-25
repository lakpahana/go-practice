package main

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(2, 3))
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(quote.Go())

	message := greetings.Hello("Gladys")
	fmt.Println(message)

}

package main

import "fmt"

func double(x int) int {
	return x + x
}

func total(num1, num2 int) int {
	return num1 + num2
}

func sayHello() {
	fmt.Println("Hello everyone!")
}

func main() {
	sayHello() // -> Hello everyone!

	sum := double(5)
	fmt.Println("sum is:", sum) // -> 10
	
	count := total(sum, 3)
	fmt.Println(count) // -> 13

	nestingFunction := total(double(2), 8)
	fmt.Println(nestingFunction) // -> 12
}

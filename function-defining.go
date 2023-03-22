// Function definition method
package main

func average(num1, num2 int) int {
	return (num1 + num2) / 2
}

func multiReturn() (int, int, int) {
	return 1, 2, 3
}

a, b, _ := multiReturn() // -> a: 1 , b: 2, number 3 is ignored
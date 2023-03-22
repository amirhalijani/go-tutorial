package main

import "fmt"

func main() {
	var userName = "Kurt Cobain"
	fmt.Println("username is", userName)

	var firstName string = "Joey"
	fmt.Println("first name:", firstName)

	name := "Mark"
	fmt.Println("name:", name)

	var avg int
	fmt.Println("average is", avg)

	crs1, crs2 := 18, 16
	fmt.Println("course 1 is", crs1, "course 2 is", crs2)

	avg = (crs1 + crs2) / 2
	fmt.Println("average is", avg)
}

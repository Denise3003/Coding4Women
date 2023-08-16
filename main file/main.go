package main

import (
	"fmt"
	"mymod/mypackage"
)

func add(a, b int) int {
	
	return a + b
}

func main() {
	var newNumber int
	newNumber = add(2, 3)

	fmt.Println(newNumber)
	mypackage.Testfunc()

}

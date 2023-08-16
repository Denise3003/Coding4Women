package main

import "fmt"

func Account(num int) bool {
	if num > 0 {
		fmt.Println("You have money")
		return true

	}
	fmt.Println("Get a job.")
	return false
}

// func main() {
// 	fmt.Println(Account(1))
// }

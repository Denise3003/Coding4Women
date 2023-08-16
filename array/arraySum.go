package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arrSum := 0

	for i := 0; i < len(arr1); i++ {
		arrSum = arrSum + arr1[i]
	}
	fmt.Println("The Sum is=", arrSum)
}

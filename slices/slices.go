package main

import "fmt"

func main() {

	var s []string
	fmt.Println("Uninit:", s, s == nil, len(s) == (0))

}

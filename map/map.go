package main

import "fmt"

func Mymap() {
	myMap := map[string]int{"Danni": 56, "Diana": 48}
	myMap["Olly"] = 65
	myMap["Danni"] = 3
	fmt.Println(myMap["Olly"])
}

func main() {
	Mymap()
}

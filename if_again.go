package main

import "fmt"

func main() {

	testScore := (49)
	testScoreGrade5 := 100
	testScoreGrade4 := 70
	testScoreGrade3 := 50

	if testScore >= testScoreGrade5 {
		fmt.Println("Top Mark")
	} else if testScore >= testScoreGrade4 {
		fmt.Println("Pass with Distiction")
	} else if testScore >= testScoreGrade3 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Failed")
	}
}

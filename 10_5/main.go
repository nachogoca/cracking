package main

import "fmt"

// Given a sorted array of strings that is interspersed with empty strings,
// write a method to find the location of a given string

func main() {

	words := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	w := "at"

	fmt.Println(find(words, w))

}

func find(ss []string, s string) (int, error) {

	if len(ss) == 0 {
		return 0, fmt.Errorf("word not found")
	}

	leftPivot := 0
	rightPivot := len(ss) - 1
	middlePivot := (rightPivot - leftPivot) / 2
	moveLeft := true
	middleMovement := 1

	for middlePivot >= leftPivot && middlePivot <= rightPivot {
		fmt.Println(leftPivot, middlePivot, rightPivot)

		if ss[middlePivot] != "" {
			// found a non-empty word
			// check for string
			if ss[middlePivot] == s {
				return middlePivot, nil
			}

			// move the pivots
			if ss[middlePivot] < s {
				// move left pivot
				leftPivot = middlePivot
				middlePivot = (rightPivot - leftPivot) / 2
			} else {
				// move right pivot
				rightPivot = middlePivot
				middlePivot = (rightPivot - leftPivot) / 2
			}
			middleMovement = 1
		} else {
			// found an empty string, we have to move the middle pivot
			if moveLeft {
				middlePivot -= middleMovement
			} else {
				middlePivot += middleMovement
			}
			moveLeft = !moveLeft
			middleMovement += 1

		}

	}

	return 0, fmt.Errorf("word not found")
}

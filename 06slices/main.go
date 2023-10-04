package main

import (
	"fmt"
	"sort"
)

func main() {
	//Creaing slices

	// 1st Method

	var fruitList = []string{"banana", "peach", "apple"}

	fmt.Println("1st Slice : ", fruitList)

	//Adding to slice using append,  Any changes done to slice will change original slice also
	// append reallocates memory

	fruitList = append(fruitList, "watermelon")

	fmt.Println("2nd slice : ", fruitList)

	// Manipularing slices

	fruitList = append(fruitList[1:3]) // it will slice fruitList from index value 1 and upto 3 (remember - value at index 3 is not included)

	fmt.Println("3rd slice : ", fruitList)

	// 2nd Method

	// Using the make keyword

	highScore := make([]int, 4) // make(1st para, 2nd para) 1st para:- define slice with [] and give type to it , 2nd para:- give number of items to be stored in the slice

	// Adding values to slice

	highScore[0] = 111
	highScore[1] = 333
	highScore[2] = 666
	highScore[3] = 555

	fmt.Println("highscore slice :- ", highScore)

	//if you dont add value to slice it will keep empty space at that index

	// Adding values to slice using append

	highScore = append(highScore, 444, 777, 999, 888)

	fmt.Println("Updated highScore slice :- ", highScore)

	// Sorting slice

	sort.Ints(highScore)

	fmt.Println("Sorted slice in increasing order :- ", highScore)

	// Checking slice is sorted or not

	fmt.Println("Is highScore sorted ? ", sort.IntsAreSorted(highScore))

	/* Removing value from slice based on index */

	var courses = [4]string{"React", "Go", "Angular", "C++"}

	fmt.Println("Course slice :-", courses)

	var index int = 2

	var newCourseList = append(courses[:index], courses[index+1:]...)

	fmt.Println("Updated course list :- ", newCourseList)

}

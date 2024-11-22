package main

import (
	"fmt"
)

func main() {

	var emptySlice []int
	//emptySlice = make([]int, 5) //5개의 원소 생성
	if len(emptySlice) == 0 {
		emptySlice = append(emptySlice, 77)
	}
	fmt.Printf("%#v\n", emptySlice)

	var gpa [5]float64 = [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} //make array
	gpa_slice := gpa[1:4]                                     //4.1, 4.5, 3.9
	gpa[2] = 2.71
	gpa_slice = append(gpa_slice, 4.3)
	fmt.Println(len(gpa_slice), gpa_slice, gpa)

}

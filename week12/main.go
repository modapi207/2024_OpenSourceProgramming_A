package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {

	var gpa [3]float64

	for i := 0; i < len(gpa); i++ {
		fmt.Print("Input float number :") //go get github.com/headfirstgo/keyboard
		gpa[i], _ = keyboard.GetFloat()
	}

	for index, value := range gpa {
		fmt.Printf("%f\n", index, value)
	}
}

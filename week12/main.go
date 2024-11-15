package main

import (
	"fmt"

	"github.com/headfirstgo/keyboard"
)

func main() {
	//go get github.com/headfirstgo/keyboard
	var gpa [3]float64

	for i := 0; i < len(gpa); i++ {
		fmt.Print("Input folat number")
		gpa[i], _ = keyboard.GetFloat()
	}

	for index, value := range gpa {
		fmt.Printf("%d %f\n", index, value)
	}
}

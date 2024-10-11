package main

import (
	"fmt"
	"strings"
)

func main() {
	letter := "공부 ?이팅!"
	fixletter := strings.NewReplacer("?", "파")
	fmt.Println(fixletter.Replace(letter))
}

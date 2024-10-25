package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //42 같은 고정 값을 넣으면 안 돼
	target := rand.Intn(6) + 1   //지금은 1~5
	fmt.Printf("%d\n", target)
}

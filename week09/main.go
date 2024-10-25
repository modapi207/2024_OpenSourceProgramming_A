package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) //42 같은 고정 값을 넣으면 안 돼
	answer := rand.Intn(6) + 1   //지금은 1~5
	fmt.Printf("%d\n", answer)

	fmt.Print("숫자 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i) // python strip과 유사
	guess, err := strconv.Atoi(i)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)

	if answer == guess {
		fmt.Println("정답입니다!")

	} else if answer > guess {
		fmt.Println("입력하신 수는 정답보다 작은 수입니다. LOW")
	} else {
		fmt.Println("입력하신 수는 정답보다 큰 수입니다. HIGH")
	}
}

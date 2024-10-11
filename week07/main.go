package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("정수 입력 : ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                  // python strip과 유사
	score, err := strconv.ParseInt(i, 10, 32) // 문자열 변수 i 값을 정수형(32비트)으로 변환, 입력 받은 값은 16진수로 처리.
	var aOrNot string
	if score >= 90 {
		aOrNot = "A"

	} else {
		aOrNot = "BCDF"

	}
	fmt.Printf("%d점은 %s등급입니다\n", score, aOrNot)
}

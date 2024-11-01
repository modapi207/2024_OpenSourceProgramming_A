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

	i = strings.TrimSpace(i)

	n, err := strconv.Atoi(i)
	counts := 0
	//몇 번 나누어 떨어졌는지 기록하는 변수

	for j := 1; j <= n; j++ { // 1부터 입력된 수까지 반복
		if n%j == 0 { // 약수면
			counts++ // 나누어 떨어지는 횟수
		}
	}

	if counts == 2 {
		fmt.Printf("%d는(은) 소수입니다", n)
	} else {
		fmt.Printf("%d는(은) 소수가 아닙니다", n)
	}
}

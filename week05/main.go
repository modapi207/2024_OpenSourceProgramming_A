package main

import "fmt"

func main() {
	var i int = 55 //타입이 뒤에 표시, 세미클론 X
	fmt.Printf("value i : %d\n", i)
	fmt.Print("value i : ", i, "\n") //줄바꿈 필요할 경우 역슬래시 n 붙임
	fmt.Println("value i :", i)      //줄바꿈,뒤의 띄어쓰기 포함되어있음

}

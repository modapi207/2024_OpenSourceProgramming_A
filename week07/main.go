package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	//var month int = int(now.Month()) //앞에 int 붙여줘서 캐스팅 해줌. 원래 자료형이 int가 아니었음.
	fmt.Printf("오늘은 %d년 %d월 %d일 입니다.\n", now.Year(), int(now.Month()), now.Day())
	fmt.Printf("지금은 %d시 %d분 %d초 입니다.\n", now.Hour(), int(now.Minute()), now.Second())
	fmt.Println(now.Month())
}

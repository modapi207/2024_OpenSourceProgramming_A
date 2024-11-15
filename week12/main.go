package main

import (
	"fmt"
	"time"
)

func main() {
	dates := [3]time.Time{time.Unix(0, 0),
		time.Unix(1447920000, 0),
		time.Unix(1447920001, 0), //쉼표가 반드시 필요
	}
	//fmt.Println(dates[0], dates[1], dates[2])
	//fmt.Println(dates) //배열 그 자체를 짝음
	//fmt.Printf("%#v\n", dates) //어레이 리터럴 찍어봄.

	for i := 0; i < len(dates); i++ {
		fmt.Println(i, dates[i])
	}
}

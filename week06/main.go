package main

import (
	"fmt"
) //필요한 모듈 자동 추가, 필요 없으면 자동 삭제
func main() {
	var f float64
	var i int
	var b bool
	var s string

	fmt.Println(f, i, b, s)
	fmt.Printf("%f%t%s%d\n", f, b, s, i) //zero value
	f = 2.7                              //이미 선언된 것에 단축 연산자 사용 X
	i = 3
	fmt.Print("\n\n", f > float64(i), "\n") //comparision (true/false)
}

package main

import (
	"fmt"
	"math"
	"reflect"
	"strings" //패키지 (비슷 =stdio.h)
) //필요한 모듈 자동 추가, 필요 없으면 자동 삭제
func main() {
	//var i int = 55                   // 타입이 뒤에 표시, 세미클론 X

	// var i int
	// i = 55

	var f float32 = 4.30 //명시적/ f :=4.30 단, 무슨 값인지 한 번에 알아보기 힘듦
	//float32=float  float64==c언어의 double, GO언어의 경우 필요 없는 변수가 있으면 에러가 남.

	//f := 4.30 //float64
	i := 55 //단축 연산자 (넘어오는 값을 추론해 실수면 float 등, 자동으로)' 추측 가능해야 함. 오른쪽이 숫자니까 float겠구나 등

	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Printf("%s\n", strings.Title("jung inha")) //타이틀은 대문자로
	fmt.Println(math.Ceil(3.99))                   //ceil은 올림

	fmt.Printf("value i : %d\n", i) //줄바꿈X
	fmt.Println("value i :", i)
	fmt.Print("value i :", i, "\n") //줄바꿈X 필요할 경우 역슬래시 n 붙임
	fmt.Println("value i :", i)     //줄바꿈O,뒤의 띄어쓰기 포함되어있음

}

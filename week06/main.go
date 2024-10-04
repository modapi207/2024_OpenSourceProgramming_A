package main

import (
	"fmt"
	"reflect"
) //필요한 모듈 자동 추가, 필요 없으면 자동 삭제
func main() {
	i := 13 //단축 연산자 (넘어오는 값을 추론해 실수면 float 등, 자동으로)'
	f := 12.9
	fmt.Printf("value i : %d\n, value f : %f\n", i, f) //줄바꿈X invalid operation: i*f (mismatched typed int and float64) 라는 메세지가 뜸
	fmt.Printf("%d * %f = %f\n", i, f, float64(i)*f)   //i 타입은 그대로임. 일시적으로 바뀐 것. conversion (int로 형태 변환시 소수점 아래는 삭제)
	fmt.Println(reflect.TypeOf(i))
}

package main

import "fmt"

type visitor struct {
	age  int
	cost int
}

func calculateCost(visitors []visitor) int {
	totalcost := 0
	for _, v := range visitors {
		totalcost = totalcost + v.cost
	}
	return totalcost
} //공개 안 할 거니까 앞에 소문자 사용. visit이라는 구조체 타입.

func main() {
	var numVisitors int
	fmt.Println("How many visitors? ")
	fmt.Scanln(&numVisitors)

	var vs []visitor                  //구조체 슬라이스
	vs = make([]visitor, numVisitors) //숫자를 고정시키지 않으려 함.

	for i := 0; i < numVisitors; i++ {
		var age int
		fmt.Print("Input age : ")
		fmt.Scanln(&age)

		switch {
		case age < 12:
			vs[i] = visitor{age: age, cost: 5000}

		case age >= 12 && age < 65:
			vs[i] = visitor{age: age, cost: 10000}

		default: //65세 이상
			vs[i] = visitor{age: age, cost: 7000}
		} //if로 바꿔보기

	} //슬라이스니까 index 활용
	fmt.Printf("Total price is %dwon.", calculateCost(vs)) //구조체 받아옴
}

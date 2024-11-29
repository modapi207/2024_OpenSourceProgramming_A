package main

import "fmt"

func main() {
	ages := make(map[string]int)

	var name string
	var age int

	for {
		fmt.Print("이름 입력 (종료시 'q' 입력): ")
		fmt.Scanln(&name) //지역변수 name에 저장 (주소)
		if name == "q" {
			break
		}
		fmt.Print("나이 입력: ")
		fmt.Scanln(&age)

		ages[name] = age
	}

	for name, age := range ages {
		fmt.Printf("%s(은)는 %d살입니다. \n", name, age)
	}
}

package main

import "fmt"

func main() {
	//b := make([]interface{}, 3) // make는 배열을 만드는 듯 함, interface{}를 사용하면 아무 타입이나 올 수 있다.
	s := make([]string, 3) // make는 배열을 만드는 듯 함, stirng으로만 이루어진
	fmt.Println("emp : ", s)
	//s[0] = 1 go는 타입을 강한 수준으로 검사함으로 안됨
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s) //왜? c가 stirng 배열이여서, len(s)를 통하여 배열의 길이를 얼마로 할지
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sli1: ", l)

}

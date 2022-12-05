package main

import "fmt"

func main() {
	m := make(map[string]int) //key는 스트링 value는 int
	//a := make(map[interface{}]interface{}) //이러면 아무타입이나 올 수 있음
	//b := make(map[string]interface{})

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2") //해당 방식을 통해 key, value를 제거할 수 있다.
	fmt.Println("map:", m)

	_, prs := m["k2"] //_,가 뭐지 이러면 k2만들어가지 않나 근데 k2는 값이 없으니까 기본값이 들어가겠지
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2} //배열을 선언하는 방법 중 하나
	fmt.Println("map:", n)

}

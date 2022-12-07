package main

import "fmt"

type person1 struct {
	name interface{}
	age  interface{}
}

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20}) //이렇게 key값을 선언하지 않아도

	fmt.Println(person{name: "Alice", age: 30}) //key값을 선언해도

	fmt.Println(person{name: "Fred"}) //하나만 선언하면 기본값이!

	fmt.Println(person1{name: "Son"}) //하나만 선언하면 기본값이들어가는데 interface{} 의 기본값은 뭐야 nil~

	fmt.Println(&person{name: "Ann", age: 40}) //&넣으면 구조체 포인터 생성

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) //.name, .age로 변수에 대한 접근 가능

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51 //포인터로 접근하였기에 변수 변경 가능
	fmt.Println(sp.age)

}

//구조체를 응용할 수 있는 방식, 플젝은?

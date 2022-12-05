package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())   // 10 * 5
	fmt.Println("perim: ", r.perim()) // 10 * 2 + 5 * 2

	rp := &r //포인터로 가져오기 값을 바꾸는 로직이 없어 똑같이 나올 것
	fmt.Println("area : ", rp.area())
	fmt.Println("perim : ", rp.perim())

}

package main

import "fmt"

func plus2(a int, b int) int {
	return a + b
}

func plusPlus2(plus, c int) int {
	return plus + c //이렇게 func(?) 구조체를 상속해도 된다.
}

func main() {
	res := plus2(1, 2)
	fmt.Println("1 + 2 = ", res)

	res = plusPlus2(res, 3)
	fmt.Println("1 + 2 + 3 = ", res)
}

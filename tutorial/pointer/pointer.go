package main

import "fmt"

func zeroVal(iVal int) {
	iVal = 0 //iVal에 복사하는 역할만 수행한다
}

func zeroPtr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal: ", i)

	zeroPtr(&i) //메모리 주소를 참조하기에 값을 바꿀 수 있다
	fmt.Println("zeroPtr:", i)

	fmt.Println("pointer", &i) //메모리 주소를 참조하기에 값을 바꿀 수 있다

}

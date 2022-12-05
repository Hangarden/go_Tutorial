package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a) //"emp : 5"

	a[4] = 100
	fmt.Println("set:", a)    // set 0 0 0 0 100
	fmt.Println("get:", a[4]) // get : 100

	fmt.Println("len: ", len(a)) // len : 5

	b := [5]int{1, 2, 3, 4, 5} //배열과 동시에 값 선언
	fmt.Println("dcl : ", b)   //dcl : [1 2 3 4 5]

	var twoD [2][3]int       //배열 선언 기본값 들어감
	for i := 0; i < 2; i++ { //2
		for j := 0; j < 3; j++ { //3 2*3 생성
			twoD[i][j] = i + j //0,1,2, 1,2,3
		}
	}

	fmt.Println("2d: ", twoD) //2d : [[0,1,2] [1,2,3]]
}

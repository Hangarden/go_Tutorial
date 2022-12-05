package main

import "fmt"

func main() {

	i := 1
	for i <= 3 { // 조건
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ { //선언, 조건, 후행
		fmt.Println(j)
	}
	for { //무한루프
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ { //continue시에는 후행처리 ++함.
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)

	}

}


package main

import "fmt"

func sum(nums ...int) { //int의 갯수를 한정하지 않을 떄ㅐ
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums { //_, 사용하는 경우 정확히 파악
		total += num
	}
	fmt.Println(total)

}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4} //배열 선언 시 사용
	sum(nums...)

}

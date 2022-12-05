package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)           // 6e+11
	fmt.Println(int64(d))    // (3의 20제곱 / 500000000)  64비트까지
	fmt.Println(math.Sin(n)) // -0.28470407323754404 싸인 값 구하기
}

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write", i, "as") //"Write", 2, "as"
	switch i {
	case 1:
		fmt.Println("one")
	case 2: //2임으로 two출력
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { //현재시간 출력, time 패키지
	case time.Saturday, time.Sunday: // 토일이면 주말
		fmt.Println("It's the weekend")
	default: //아닐시 weekday if문과 비슷한 구조
		fmt.Println("It's a weekday")
	}

	t := time.Now() //현재시간
	switch {
	case t.Hour() < 12: //현재시간이 12시 전이면 before noon
		fmt.Println("It's before noon")
	default: //이후면 after noon
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) { //whatAmI 선언 interface{} 가 뭐지 -> Go의 모든 타입을 나타내기 위해 사용
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}

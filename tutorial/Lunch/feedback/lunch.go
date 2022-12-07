package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var slice = []string{"Reacher", "Juda", "Liam", "Zora", "Aiden", "Cindy", "Tommy", "Lina", "Garden"} //배열의 개수를 적지 않고 선언, slice에 직접 입력해도 될 것같고
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	for i := 0; i < 3; i++ {
		fmt.Printf("%v \n", slice[3*i:3*(i+1)])
	}
}

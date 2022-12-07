package main

import (
	"fmt"
	"math"
	"math/rand"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	var slice = []string{"Reacher", "Juda", "Liam", "Zora", "Aiden", "Cindy", "Tommy", "Lina", "Garden"} //배열의 개수를 적지 않고 선언, slice에 직접 입력해도 될 것같고
	var a = len(slice)

	const team = 3
	var humanPerTeam = int(math.Round((float64(a) / float64(team)))) //math.Ceil을 통해서 소수점이 애매하게 떨어질 때는 올려주어야 한다.
	//const human = humanPerTeam                                       //math.random(max=9(slice길이), max.timenow, pop. slice_tmp for

	aaa := make([]string, team)
	for i := 0; i < team; i++ {
		aaa[i] = make([]string)

	}

	var lunch []string = make([]string, int(humanPerTeam))

	for i := 0; i < humanPerTeam; i++ {
		for j := 0; j < 1; j++ {
			var A = rand.Intn(len(slice))
			lunch[i] = slice[A]
			slice = remove(slice, A)
			fmt.Println(lunch)
		}
		// make(type, len, cap)
	}

	//fmt.Println(lunch)
	//var teamA = make([]string, human, human)                // make(type, len, cap)
	//var teamB = make([]string, human, human)                // make(type, len, cap)
	//var teamC = make([]string, human, human)                // make(type, len, cap)
	//
	//rand.Seed(time.Now().UnixNano()) //seed를 다르게 줘야 값이 다르게 나오기 때문
	//
	//fmt.Println(slice)
	//for i := 0; i < 3; i++ {
	//	for j := 0; j < 1; j++ {
	//		var A = rand.Intn(len(slice))
	//		teamA[i] = slice[A]
	//		slice = remove(slice, A)
	//		var B = rand.Intn(len(slice))
	//		teamB[i] = slice[B]
	//		slice = remove(slice, B)
	//		var C = rand.Intn(len(slice))
	//		teamC[i] = slice[C]
	//		slice = remove(slice, C)
	//	}
	//}
	//
	//fmt.Printf("teamA : %v \nteamB : %v \nteamC : %v ", teamA, teamB, teamC)

}

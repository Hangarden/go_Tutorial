package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	var members = []string{"Reacher", "Juda", "Liam", "Zora", "Aiden", "Cindy", "Tommy", "Lina", "Garden"} //배열의 개수를 적지 않고 선언, slice에 직접 입력해도 될 것같고
	var a = len(members)

	var team = 2
	var humanPerTeam float64 = (float64(a) / float64(team)) //math.Ceil을 통해서 소수점이 애매하게 떨어질 때는 올려주어야 한다.
	var human int = int(math.Round(humanPerTeam))           //math.random(max=9(slice길이), max.timenow, pop. slice_tmp for

	var teamA = make([]string, human, human) // make(type, len, cap)
	var teamB = make([]string, human, human) // make(type, len, cap)
	var teamC = make([]string, human, human) // make(type, len, cap)

	rand.Seed(time.Now().UnixNano()) //seed를 다르게 줘야 값이 다르게 나오기 때문

	fmt.Println(members)
	for i := 0; i < 3; i++ {
		for j := 0; j < 1; j++ {
			var A = rand.Intn(len(members))
			teamA[i] = members[A]
			members = remove(members, A)
			var B = rand.Intn(len(members))
			teamB[i] = members[B]
			members = remove(members, B)
			var C = rand.Intn(len(members))
			teamC[i] = members[C]
			members = remove(members, C)
		}
	}

	fmt.Printf("teamA : %v \nteamB : %v \nteamC : %v ", teamA, teamB, teamC)

}

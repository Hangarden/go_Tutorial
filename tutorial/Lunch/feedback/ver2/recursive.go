package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"time"
)

func judge(temp1 float64) int {
	var temp2 = math.Floor(temp1)
	var temp3 = temp1 - temp2

	if temp3 <= 0.5 {
		math.Ceil(temp1)
	} else {
		math.Floor(temp1)
	}
	return int(temp1)
}

func lunchTeam(slice []string, totalTeamPeoPle int, teamCount int, a [][]string) {
	//a := make([]string, len(slice), len(slice))
	var teamPerPeople = float64(totalTeamPeoPle) / float64(teamCount)
	if teamCount > totalTeamPeoPle {
		fmt.Println("각자도생")
	} else if totalTeamPeoPle%teamCount == 0 {
		for i := 0; i < teamCount; i++ {
			slices := slice[int(teamPerPeople)*i : int(teamPerPeople)*(i+1)] //[]string
			a[i] = make([]string, len(slices))
			copy(a[i], slices)
			fmt.Println(a[i])
			//copy(a[i], slices)
			//copy(arr, slice)
			//fmt.Printf("%v \n", slice[int(teamPerPeople)*i:int(teamPerPeople)*(i+1)])
		}
	} else {
		var temp1 = judge(teamPerPeople)
		slices := slice[:int(temp1)] //[]string
		a[1] = make([]string, len(slices))
		//fmt.Printf("%v \n", slice[:int(temp1)])
		lunchTeam(slice[int(temp1):], totalTeamPeoPle-int(temp1), teamCount-1, a[1:])
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var slice = []string{"Reacher", "Juda", "Liam", "Zora", "Aiden", "Cindy",
		"Tommy", "Lina", "Garden", "Alex", "Bob", "Camille", "Zoey", "Andy", "Pleasure"} //배열의 개수를 적지 않고 선언, slice에 직접 입력해도 될 것같고
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	var team int
	fmt.Print("점심 팀 수 : ")
	fmt.Scanln(&team)
	a := make([][]string, team)
	lunchTeam(slice, len(slice), team)

	for i := 0; i < team; i++ {
		fmt.Println("type:", reflect.ValueOf(a[i]).Type())
	}

}

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var slice = []string{"Reacher", "Juda", "Liam", "Zora", "Aiden", "Cindy", "Tommy", "Lina", "Garden"} //배열의 개수를 적지 않고 선언, slice에 직접 입력해도 될 것같고
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	var team int
	fmt.Print("점심 팀 수 : ")
	fmt.Scanln(&team)
	//var teamPerPeople float64 = float64(len(slice)) / float64(team)
	//var temp = math.Floor(teamPerPeople)
	//var temp1 = teamPerPeople - temp

	// 9명 4팀
	for i := 0; i < team; i++ {
		var temp1 = len(slice) % team                   // 9명 / 4팀
		var temp2 = float64(len(slice)) / float64(team) // 9명 / 4팀
		if temp1 == 0 {                                 // 0으로 딱 나누어 떨어질 때
			fmt.Printf("%v \n", slice[int(temp2)*i:int(temp2)*(i+1)])
		} else { //아닐 때
			var temp3 = math.Floor(temp2) // 2.25 내림 -> 2
			var temp4 = temp2 - temp3     // 2.25 - 2 = 0.25
			if temp3 <= 0.5 {             // 0.25 <= 0.5
				var temp4 = int(math.Ceil(temp1)) //2.25 올림 3
				fmt.Printf("%v \n", slice[:temp4])
				var temp5 = int(math.Round(temp1))
				for i = 0; i < team-1; i++ {
					fmt.Printf("%v \n", slice[temp4+*i:temp1*(i+1)])
				}

			} else {
				var temp4 = math.Floor(tem1)
				fmt.Printf("%v \n", slice[:int(temp4)])
			}
		}
	}
}

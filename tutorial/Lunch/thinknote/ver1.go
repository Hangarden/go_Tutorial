package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Slice []string

// IsEmpty - 스택이 비어있는지 확인하는 함수
func (s *Slice) IsEmpty() bool {
	return len(*s) == 0
}

// Pop - 스택에 값을 제거하고 top위치에 값을 반환하는 함수.
func (s *Slice) Pop() interface{} {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return nil
	} else {
		top := len(*s) - 1
		data := (*s)[top] // top 위치에 있는 값을 가져 옴
		*s = (*s)[:top]   // 스택에 마지막 데이터 제거함
		return data
	}
}

func remove(slice []string, s int) []string {
	//slice = append(slice[:s], slice[s+1:]...)
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	var slice = []string{"Reacher", "Juda", "Liam", "Zora", "Aiden", "Cindy", "Tommy", "Lina", "Garden"} //배열의 개수를 적지 않고 선언, slice에 직접 입력해도 될 것같고
	var a = len(slice)

	var team = 3
	var humanPerTeam float64 = (float64(a) / float64(team)) //math.Ceil을 통해서 소수점이 애매하게 떨어질 때는 올려주어야 한다.
	//var humanPerTeam float64 = float64(a / team) //int로 먼저 계산해서 소수점이 버려진다. 이렇게 사용하면 안됨
	var human int = int(math.Round(humanPerTeam)) //math.random(max=9(slice길이), max.timenow, pop. slice_tmp for
	var teamA = make([]string, human, human)      // make(type, len, cap)
	var teamB = make([]string, human, human)      // make(type, len, cap)
	var teamC = make([]string, human, human)      // make(type, len, cap)
	//var B = slice.Pop()
	rand.Seed(time.Now().UnixNano()) //random 숫자 -> 0,1,2,3,4,5 slice[0]
	//fmt.Println(slice[rand.Intn(9)])

	//var A = rand.Intn(len(slice))
	//fmt.Println(A)
	//fmt.Println(B)
	//fmt.Println(len(slice)) // seed가 똑같으니까 이렇게 나오겠지 seed
	//var A = rand.Intn(len(slice))
	//teamA[0] = slice[A]
	//fmt.Println(teamA)
	//slice = remove(slice, A)
	fmt.Println(slice)
	for i := 0; i < 3; i++ {
		for j := 0; j < 1; j++ {
			var A = rand.Intn(len(slice))
			teamA[i] = slice[A]
			slice = remove(slice, A)
			var B = rand.Intn(len(slice))
			teamB[i] = slice[B]
			slice = remove(slice, B)
			var C = rand.Intn(len(slice))
			teamC[i] = slice[C]
			slice = remove(slice, C)
		}
	}

	fmt.Printf("teamA : %v \nteamB : %v \nteamC : %v ", teamA, teamB, teamC)

	//배열은 선언해서 for문으로 입력해도 될 것 같고 range를
	//fmt.Println("우리 팀원은 명단은 : ", slice, ", slice의 타입은", reflect.TypeOf(slice))
	//fmt.Println("우리팀원 명수는 : ", a, ", a의 타입은", reflect.TypeOf(a))
	//fmt.Println("점심을 나누어먹을 팀 수는 : ", team, ", team의 타입은", reflect.TypeOf(team))
	//fmt.Println("올림시에는전에는? : ", humanPerTeam, ", humanPerTeam의 타입은", reflect.TypeOf(humanPerTeam))
	//fmt.Println("한팀당 들어갈 명수는 : ", human, ", human의 타입은", reflect.TypeOf(human))
	//fmt.Printf("Slice A = %v \n", teamA)
	//fmt.Printf("Slice B = %v \n", teamB)
	//fmt.Printf("Slice C = %v \n", teamC)
	//fmt.Println("teamA[0] : ", teamA[0])

}

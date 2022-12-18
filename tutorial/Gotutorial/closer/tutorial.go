package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Balance = 1000

	EarnMoney = 500

	LooseMoney = 100

	GameOverMoney = 0
)

func InputNumber() (int, error) {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}

var stdin = bufio.NewReader(os.Stdin)

func main() {
	rand.Seed(time.Now().UnixNano())
	balance := Balance
	for {
		fmt.Print("1~5사이 값을 입력하세요 : ")
		n, err := InputNumber()
		if err != nil {
			fmt.Println("숫자만 입력하세요")
		} else if n < 1 || n > 5 {
			fmt.Println(fmt.Println("1에서 5사이 값만 입력하세요. "))
		} else {
			r := rand.Intn(5) + 1
			if n == r {
				balance += 500
				fmt.Println("축하합니다. 맞추셨습니다 잔고 : ", balance)
				if balance >= 5000 {
					fmt.Println("게임 승리")
					break
				}
			} else {
				balance -= 100
				fmt.Println("아쉽습니다. 틀렸습니다 잔고 : ", balance)
				if balance <= 0 {
					fmt.Println("게임 오버")
					break
				}
			}
		}
	}
}

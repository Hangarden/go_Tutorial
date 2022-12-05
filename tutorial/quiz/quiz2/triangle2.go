package main

import "fmt"

func main() {
	for i := 5; i >= 1; i-- { //5부터 1보다 크거나 작을 때까지 1씩 --
		for j := 0; j < i; j++ {
			if j < i {
				fmt.Print("*")
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

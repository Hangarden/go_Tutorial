package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 4; j >= 0; j-- {
			if i < j {
				fmt.Print("   ")
			} else {
				fmt.Print(" * ")
			}
		}
		fmt.Println("")
	}
}

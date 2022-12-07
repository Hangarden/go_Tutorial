package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i <= j {
				fmt.Print(" * ")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println("")
	}
}

//for i := 0; i < 6; i++ {
//	for j := 5; j > 0; j-- {
//		if i < j {
//			fmt.Print("   ")
//		} else {
//			fmt.Print(" * ")
//		}
//	}
//	fmt.Println("")
//}

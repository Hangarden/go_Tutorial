package main

import "fmt"

func values() (int, int) {
	return 3, 7
}
func main() {
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	_, c := values() //일부값만 반환하려고 할 때
	fmt.Println(c)
}

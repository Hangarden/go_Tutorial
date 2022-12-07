package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	words := strings.Fields("ink runs from the corners of my mouth")
	fmt.Println(len(words))                   // 8
	rand.Shuffle(len(words), func(i, j int) { // length -> 8
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)

}

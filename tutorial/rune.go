package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const s = "สวัสดี" // hello in thai

	fmt.Println("Len:", len(s)) //길이 측정

	for i := 0; i < len(s); i++ { //s의 길이만큼 반복해서
		fmt.Printf("%x ", s[i]) //각각의 유니코드를 출력
	}
	fmt.Println() //줄바꿈

	fmt.Println("Rune count:", utf8.RuneCountInString(s)) //Rune카운트는 몇개? String에서 Rune카운트

	//rune이 32바이트니까 192바이트라는 소리인가? 바이트를 측정해보자

	for idx, runeValue := range s { //runeValue는 글자 수 출력. idx 0,1,2,3,4,5
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

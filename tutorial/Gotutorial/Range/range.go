package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	//key또는 인덱스만 필요한 경우 num := range nums라 잡아도 됩니다.
	//value만 사용해야 하는 경우 첫번째 항목을 빈식별자 사용 _,
	for _, num := range nums { //인덱스가 필요없기에 빈식별자로 무시한다 해당 구문이 근데 조건문인가? 이렇게 가도 괜찮나? range덕분인가?
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	//for num := range nums {
	//	if num.expired() { //expired?
	//		delete(nums, num)
	//	}
	//}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v) //kv값을 대입 range를 하면 자ㅏ동으로 들어가나 봄
	} //for문 돌릴 때 용이한 range

	for i, c := range "go" {
		fmt.Println(i, c) //유니코드 들어감
	}

}

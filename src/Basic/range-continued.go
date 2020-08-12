package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// indexのみ
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}

	// 値のみ
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// index, value
	for index, value := range pow {
		fmt.Println(index, value)
}
}

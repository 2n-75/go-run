package main

import "fmt"

func main() {
	// チャンネルを初期化 第２引数にバッファの長さを入れる
	ch := make(chan int, 2)
	// 長さが足りないとdeadlockが起きる
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

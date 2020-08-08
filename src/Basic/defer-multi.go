package main

import "fmt"

func main() {
	fmt.Println("counting")
	/*
	遅延関数呼び出しはスタックにプッシュされます。
	関数が戻ると、その据え置き呼び出しは
	後入れ先出しの順序で実行されます。
	*/
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

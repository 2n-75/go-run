package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 関数の中でvarを省略した書き方で変数宣言ができる
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

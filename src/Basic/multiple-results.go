package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("uma", "ushi")
	fmt.Println(a, b)
}

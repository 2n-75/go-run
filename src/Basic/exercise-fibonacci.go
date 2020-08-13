package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n2 := 0 // 2つ前の項
	n1 := 1 // 1つ前の項
	index := 0
	return func() int {
		if index == 0 {
			index++
			return 0
		}
		if index == 1 {
			index++
			return 1
		}
		n := n2 + n1
		n2 = n1
		n1 = n
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

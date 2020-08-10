package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	// b[0]を変更するとa[1]とname[1]の値も変更される
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

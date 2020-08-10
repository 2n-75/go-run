package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	// 出力する値,その名前
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
	fmt.Println(a.Name) // Arthur Dent
}

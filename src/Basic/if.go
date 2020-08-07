package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	/* こっちでも動く
	if( x < 0 ){
		return sqrt(-x) + "i"
	}
	*/
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

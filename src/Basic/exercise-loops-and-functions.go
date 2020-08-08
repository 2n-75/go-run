package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	// 10回繰り返す
	/*for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}*/

	// 十分差が小さくなったら終わり
	d := 0.0
	for i := 0; i < 10; i++ {
		d = (z*z - x) / (2 * z)
		z -= d
		if math.Abs(d) < 1e-10 {
			break
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

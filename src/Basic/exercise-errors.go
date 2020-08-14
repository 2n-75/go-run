package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

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
	// error型は成功した時nilを返す
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

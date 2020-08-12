package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	newPic := make([][]uint8, dy)
	for y := range newPic {
		newPic[y] = make([]uint8, dx)
	}
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			newPic[x][y] = uint8((x+y)/2)
		}
	}

	return newPic
}

func main() {
	pic.Show(Pic)
}

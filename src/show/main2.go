package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (a [][]uint8) {

	//a := [][]uint8{}
	for i := 0; i < dx; i ++ {
		b := []uint8{}
		for j := 0; j < dy; j ++ {

			b = append(b, uint8(i*j))
		} 
		a = append(a, b)
	}
	return
}

func main() {

	pic.Show(Pic)
}

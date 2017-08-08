package main
import "fmt"

func main() {
//		var ars *[]int = new([4]int) 
//		for i, j := range(ars) {
//
//				fmt.Println(i, j)
//		}
		ars1 := [...]int{1,2,3,4}
		for i, j := range(ars1) {

				fmt.Println(i, j)
		}
		var ars2 []int = make([]int, 4)
		for i, j := range(ars2) {

				fmt.Println(i, j)
		}
		var str string = "qweqwredsdf"
		for i, j := range(str) {

				fmt.Println(i, j)
		}
}

package main
import "fmt"

type Vector []float64


const numCPU = 4

func DoAll() {
		c := make(chan int, numCPU) 
		DoSome := func(iter int, n int, c chan int) {
				for i := 0; i < iter + 1; i ++ {
						fmt.Printf("[%d]round:%d\n", iter, i)
				}
				c <- 1
		}
		for i := 0; i < numCPU; i ++ {
				go DoSome(i * 10, numCPU, c)
		}

		for i := 0; i < numCPU; i ++ {
				<-c
		}
}

func main () {
		fmt.Println("Begin:")
		DoAll()
}

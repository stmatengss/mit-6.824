package main

import (
		"fmt"
)

func send(ch chan string) {
		ch <- "Fuck"
		ch <- "Fuck1"
		ch <- "Fuck2"
		ch <- "Fuck3"
		ch <- "exit"

}

func get(ch chan string, ch2 chan int) {
		var output string

		for {
				output = <-ch
				if output == "exit" {
						break
				}
				fmt.Println(output)
		}
		ch2 <- 1
}

func main() {
		ch := make(chan string)
		ch2 := make(chan int)

		go send(ch)
		go get(ch, ch2)

		<-ch2
}

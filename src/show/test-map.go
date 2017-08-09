package main
import "fmt"

func main() {
		var mapLit map[string]int = map[string]int{"a":1, "b":2}
		
		mapLit["fuck"] = 1

		fmt.Printf("%d", mapLit["c"])
		fmt.Printf("%d", mapLit["a"])

		mapLit2 := make(map[string]int) //Always use make to create map data structure
		mapLit2 = mapLit
}

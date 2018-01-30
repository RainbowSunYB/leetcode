package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution("c"))
}

func Solution(s string) int {
	start, max, length := 0, 0, 0
	index := map[int32]int{}
	for loc, data := range s {
		oriLoc, ok := index[data]
		if ok && oriLoc >= start {
			length = loc - start
			start = oriLoc + 1
		} else {
			length++
		}
		if length > max {
			max = length
		}
		index[data] = loc
	}
	return max
}

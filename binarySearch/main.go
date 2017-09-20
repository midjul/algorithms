package main

import "fmt"

func bSearch(n int, hs []int) bool {
	var low int
	l := len(hs)
	high := l - 1

	for low <= high {
		median := (low + high) / 2
		if hs[median] < n {
			low = median + 1
		} else {
			high = median - 1
		}

		if low == l || hs[low] != n {
			return false
		}
	}
	return true
}
func main() {
	data := []int{3, 5, 6, 7, 9, 12, 14, 16, 19, 23}
	nd := 54
	fmt.Println(bSearch(nd, data))
}

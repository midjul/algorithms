package main

import "fmt"

func lSearch(data []int, key int) bool {
	for _, val := range data {
		if val == key {
			return true
		}
	}
	return false
}
func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	key := 23

	fmt.Println(lSearch(data, key))
}

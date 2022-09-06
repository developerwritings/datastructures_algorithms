//  Program to write an linear Search
package main

import "fmt"

func main() {
	arr := []int{62, 63, 64, 65, 66}
	fmt.Println(linearSearch(arr, 63))
}

// Standard Searching
func linearSearch(arr []int, searchItem int) bool {
	for _, val := range arr {
		if val == searchItem {
			return true
		}
	}
	return false
}

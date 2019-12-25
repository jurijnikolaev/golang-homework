package main

import "fmt"

func main() {
	arr := [9]int{1, 2, 3, 4, 1, 2, 2, 3, 2}
	fmt.Println(unique_count(&arr))
}

func unique_count(array *[9]int) int {
	count := 1
	for i := 0; i < len(*array); i++ {
		for j := 0; j < len(*array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	for i := 0; i < len(*array)-1; i++ {
		if array[i] != array[i+1] {
			count++
		}
	}
	return count
}

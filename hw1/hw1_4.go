package main

import "fmt"

func main() {
	arr := [8]float64{44.7, 23.66, 1.09, 9.01, 13.4, 22.2001, 4.01, 16.00006}
	bubble_sort(&arr)
}

func bubble_sort(array *[8]float64) {
	for i := 0; i < len(*array); i++ {
		for j := 0; j < len(*array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	fmt.Println(*array)
}

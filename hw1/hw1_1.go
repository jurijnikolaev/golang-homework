package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	fmt.Println("Введите 1-й множитель:")
	fmt.Fscan(os.Stdin, &a)
	fmt.Println("Введите 2-й множитель:")
	fmt.Fscan(os.Stdin, &b)
	fmt.Println("Результат:")
	fmt.Print(a, " * ", b, " = ", multiply(a, b))
}

func multiply(x, y int) int {
	if x < 0 {
		x = -x
		y = -y
	}
	z := make([]int, x)
	result := 0
	for i := 0; i < x; i++ {
		z[i] = y
		result += z[i]
	}
	return result
}

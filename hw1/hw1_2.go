package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Println("Введите n:")
	fmt.Fscan(os.Stdin, &n)
	fmt.Print(n, "-е число Фибоначчи = ", fibonacci1(n))
}

func fibonacci1(x int) int {
	f := make([]int, x+1)
	f[0] = 1
	f[1] = 1
	for i := 2; i <= x; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[x-1]
}

package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Println("Введите n:")
	fmt.Fscan(os.Stdin, &n)
	fmt.Print(n, "-е число Фибоначчи = ", fibonacci2(n))
}

func fibonacci2(x int) int {
	if x == 1 || x == 2 {
		return 1
	} else {
		return fibonacci2(x-1) + fibonacci2(x-2)
	}
}

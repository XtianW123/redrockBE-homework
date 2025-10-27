package main

import "fmt"

func main() {
	a := 1
	i := 1
	x := a * i
	for i = 1; i < 10; i++ {
		for a <= i {
			x = a * i
			fmt.Printf("%d*%d=%d\n", a, i, x)
			a++
		}
		fmt.Printf("\n")
		a = 1
	}

}

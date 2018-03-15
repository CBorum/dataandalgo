package main

import (
	"fmt"
)

func main() {
	x := 3
	precision := .0001
	fmt.Println(calcRoot1(x, precision))
	fmt.Println(calcRoot2(x, precision))
}

func calcRoot1(x int, precision float64) float64 {
	root := 0.0
	for int(root) < x/2+1 {
		if int(root*root) == x {
			break
		}
		// fmt.Println(root)
		root += precision
	}
	return root
}

func calcRoot2(x int, precision float64) float64 {
	root := 0.0
	min := .0
	max := float64(x)

	for {
		root = (min + max) / 2
		fmt.Println(root)
		fmt.Println(root * root)
		if root * root == float64(x) {
			fmt.Println("equal")
			break
		} else if int(root*root) < x {
			min = root
			fmt.Println("less")
		} else if int(root*root) > x {
			max = root
			fmt.Println("more")
		}
	}

	return root
}

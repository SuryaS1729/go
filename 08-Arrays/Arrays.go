package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 120
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) //gets length of array

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD3by2 [3][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			twoD3by2[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD3by2) // okey nice

	var twoDtwoby3 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDtwoby3[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoDtwoby3)

	twoDtwoby3 = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("new2d: ", twoDtwoby3)
}

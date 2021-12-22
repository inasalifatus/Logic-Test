package main

import "fmt"

func main() {
	arr3d := [3][3]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, 12},
	}
	fmt.Println(arr3d)

	A := arr3d[0][0] + arr3d[1][1] + arr3d[2][2]
	B := arr3d[0][2] + arr3d[1][1] + arr3d[2][0]
	result := A + B
	fmt.Println(result)
}

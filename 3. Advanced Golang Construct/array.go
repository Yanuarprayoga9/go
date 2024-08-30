package main

import "fmt"

func main() {
	// p := []int{6, 4, 5, 6, 6}
	p2 := [][][]int{
		{
			{2},
		}, {
			{3},
		},
	}

	// for i := 0; i < len(p2); i++ {
	// 	for j := 0; j < len(p2[i]); j++ {
	// 		for k := 0; k < len(p2[i][j]); k++ {
	
	// 			fmt.Println(p2[i][j][k])
	// 		}
	// 	}
	// }
	// Loop through the 3-dimensional slice
	for i := 0; i < len(p2); i++ {
		for j := 0; j < len(p2[i]); j++ {
			for k := 0; k < len(p2[i][j]); k++ {
				fmt.Println(p2[i][j][k])
			}
		}
	}
	// fmt.Println(p2[0][1])
}

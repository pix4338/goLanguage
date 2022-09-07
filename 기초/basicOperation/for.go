package main

import "fmt"

func main() {
	var i int
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i = 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println("최종 i 값은", i)
	// 0 1
	// 1 3
	// 2 5
	// 3 7
	// 4 5
	// 5 3
	// 6 1

	k := 1
	q := 3

	for i = 0; i < 7; i++ {
		// 공백 부분

		for p := 0; p < q; p++ {
			fmt.Print(" ")
		}

		if i < 3 {
			q--
		} else {
			q++
		}
		// * 부분

		for j := 0; j < (2*k - 1); j++ {
			fmt.Print("*")
		}

		if i < 3 {
			k++
		} else {
			k--
		}
		fmt.Print("\n")
	}
}

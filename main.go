package main

import "fmt"

func main() {
	var ss int
	var cc int
	fmt.Println("Введи высоту")
	fmt.Scan(&ss)
	fmt.Println("Введи ширину")
	fmt.Scan(&cc)
	for i := 0; i < ss; i++ {
		for j := 0; j < cc; j++ {

			if (i+j)%2 == 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

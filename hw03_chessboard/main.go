package main

import "fmt"

func main() {
	// Place your code here.
	var h, v int
	fmt.Print("Введите строки столбцы > ")
	_, err := fmt.Scan(&h, &v)
	if err != nil {
		panic(err)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < v; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

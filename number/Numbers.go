package main

import "fmt"

func main() {
	var X int
	for {
		_, err := fmt.Scanln(&X)
		if err != nil {
			fmt.Println("input error!")
		}
		if X == 0 {
			break
		}
		for i := 1; i <= X; i++ {
			fmt.Printf("%d ", i)
		}
		fmt.Println()
	}
}

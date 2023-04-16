package main

import "fmt"

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&arr[i])
		if arr[i] <= 0 {
			arr[i] = 1
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("X[%d] = %d\n", i, arr[i])
	}
}

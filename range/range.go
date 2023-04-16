package main

import "fmt"

func main() {
	var M, N int
	for {
		_, err := fmt.Scan(&M, &N)
		if err != nil || M <= 0 || N <= 0 {
			break
		}
		if M > N {
			M, N = N, M
		}
		sum := 0
		for i := M; i <= N; i++ {
			fmt.Printf("%d ", i)
			sum += i
		}
		fmt.Printf("Sum=%d\n", sum)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	chars := strings.Split(s, "")
	newStr := strings.Join(chars, " ")
	fmt.Println(newStr)
}

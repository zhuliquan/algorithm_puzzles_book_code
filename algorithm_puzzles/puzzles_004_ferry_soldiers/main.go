package main

import "fmt"

func main() {
	var N int
	fmt.Println("input how many soldier:")
	fmt.Scanf("%d", &N)
	fmt.Println(4*(N-1) + 1)
}

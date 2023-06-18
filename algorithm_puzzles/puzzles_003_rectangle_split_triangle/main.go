package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("input a number: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		panic(err.Error())
	}
	k := n / 2

	for i := 0; i < k; i++ {
		if n == 2*k+1 && i == k-1 {
			fmt.Print(`|------|
|    / |
|   /  |
|  /\  |
| /  \ | 
|/    \|`)

		} else {
			fmt.Print(`|------|
|     /|
|    / |
|   /  |
|  /   |
| /    | 
|/     |`)
		}

		if i == k-1 {
			fmt.Println("\n|------|")
		} else {
			fmt.Println()
		}
	}

	fmt.Print()
}

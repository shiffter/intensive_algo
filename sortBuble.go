package main

import "fmt"

func main() {
	n := []int{5, 3, 4, 1}
	sortBuble(n)
	fmt.Println(n)
}

func sortBuble(n []int) {
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-1; j++ {
			if n[j] > n[i] {
				n[j], n[i] = n[i], n[j]
			}
		}
	}
}

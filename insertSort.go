package main

import "fmt"

func main() {
	n := []int{3, 5, 3, 1}
	insertSort(n)
	fmt.Println(n)
}

func insertSort(n []int) {
	for i := 1; i < len(n); i++ {
		for j := i; j > 0 && n[j-1] > n[j]; j-- {
			n[j-1], n[j] = n[j], n[j-1]
		}
	}
}

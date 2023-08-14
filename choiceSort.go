package main

import "fmt"

func main() {
	n := []int{-100, 5, 2, 1, 1}
	choiceSort(n)
	fmt.Println(n)
}

func choiceSort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		index := i
		for j := i + 1; j < len(n); j++ {
			if n[j] < n[index] {
				index = j
			}
		}
		n[i], n[index] = n[index], n[i]
	}
}

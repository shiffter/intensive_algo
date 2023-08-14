package main

import "fmt"

func main() {
	n := []int{3, 5, 2, 1}
	choiceSort(n)
	fmt.Println(n)
}

func choiceSort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		min := n[i]
		index := i
		for j := i + 1; j < len(n); j++ {
			if n[j] < min {
				min = n[j]
				index = j
			}
		}
		n[i], n[index] = n[index], n[i]
	}
}

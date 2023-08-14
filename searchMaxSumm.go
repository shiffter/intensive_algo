package main

import "fmt"

func main() {
	mass := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(searchMaxSum(mass))
}

func searchMaxSum(mass []int) int {
	if len(mass) == 0 {
		return 0
	}
	summ := mass[0]
	currMax := mass[0]

	for i := 1; i < len(mass); i++ {
		summ += mass[i]
		if summ > currMax {
			currMax = summ
		}
		if summ < 0 {
			summ = 0
		}
	}
	return currMax
}

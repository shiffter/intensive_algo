package main

import "fmt"

type tree struct {
	Val   int
	Right *tree
	Left  *tree
}

func main() {
	tree := &tree{Val: 5,
		Right: &tree{Val: 3, Right: nil},
		Left:  &tree{Val: 10, Left: &tree{Val: -3}}}

	summ := 0
	SummTreeNode(tree, &summ)
	fmt.Println(summ)
}

func SummTreeNode(node *tree, summ *int) {

	if node == nil {
		return
	}

	*summ = *summ + node.Val

	if node.Left != nil {
		SummTreeNode(node.Left, summ)
	}

	if node.Right != nil {
		SummTreeNode(node.Right, summ)
	}
}

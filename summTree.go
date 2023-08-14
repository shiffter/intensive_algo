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
	summ = SummTreeNode(tree, summ)
	fmt.Println(summ)
}

func SummTreeNode(node *tree, summ int) int {

	if node == nil {
		return 0
	}

	return SummTreeNode(node.Left, node.Val) + SummTreeNode(node.Right, node.Val) + node.Val

}

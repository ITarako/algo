package linkedListDoubly

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

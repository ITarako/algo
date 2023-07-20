package linkedListSingly

type Node struct {
	Next *Node
	Val  int
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

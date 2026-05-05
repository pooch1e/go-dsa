package linkedlist

// RandomNode is a singly linked list in which a node may be connected to a random node
// in addition to the typical Next node.
type RandomNode struct {
	// Val is the Value of the node
	Val int
	// Next is the Next node in the link list
	Next *RandomNode
	// Random is sometime used to connect to a random node
	Random *RandomNode
}

// CopyLinkedListWithRandomPointer solves the problem in O(n) time and O(n) space.
func CopyLinkedListWithRandomPointer(head *RandomNode) *RandomNode {
	panic("not implemented")

}

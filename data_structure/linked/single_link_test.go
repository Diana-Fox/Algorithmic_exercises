package linked

import (
	"fmt"
	"testing"
)

func TestNewSingleLink(t *testing.T) {
	link := NewSingleLink[int]()
	node1 := NewSingleLinkNode[int](1)
	node2 := NewSingleLinkNode[int](2)
	node3 := NewSingleLinkNode[int](3)
	node4 := NewSingleLinkNode[int](4)
	link.InsertNodeFront(node1)
	link.InsertNodeFront(node2)
	link.InsertNodeFront(node3)
	link.InsertNodeFront(node4)
	node5 := NewSingleLinkNode[int](5)
	node6 := NewSingleLinkNode[int](6)
	node7 := NewSingleLinkNode[int](7)
	node8 := NewSingleLinkNode[int](8)
	link.InsertNodeBack(node5)
	link.InsertNodeBack(node6)
	link.InsertNodeBack(node7)
	link.InsertNodeBack(node8)
	fmt.Println(link.String())
	link.DeleteNode(node2)
	fmt.Println(link.String())
	link.DeleteIndex(3)
	fmt.Println(link.String())
	fmt.Println(link.GetNodeAtIndex(2))
}

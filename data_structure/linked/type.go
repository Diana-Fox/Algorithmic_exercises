package linked

// LinkedInterface 链表的接口
type LinkedInterface[T any] interface {
	GetFirstNode() *SingleLinkNode[T]                                                //头节点
	InsertNodeFront(node *SingleLinkNode[T])                                         //头插
	InsertNodeBack(node *SingleLinkNode[T])                                          //尾插
	InsertNodeValueFront(dest T, node *SingleLinkNode[T], f func(v1, v2 T) int) bool //头插
	InsertNodeValueBack(dest T, node *SingleLinkNode[T], f func(a T, b T) int) bool  //尾插
	GetNodeAtIndex(index int) *SingleLinkNode[T]                                     //根据索引查询
	DeleteNode(dest *SingleLinkNode[T]) bool                                         //删除一个节点
	DeleteIndex(index int) bool                                                      //删除指定位置的节点
	String() string                                                                  //字符串
}

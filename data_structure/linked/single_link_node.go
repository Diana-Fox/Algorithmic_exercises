package linked

type SingleLinkNode[T any] struct {
	value T //当前节点的值
	//prev  *SingleLinkNode[T] //前驱节点
	next *SingleLinkNode[T] //后继节点
}

// NewSingleLinkNode 创建新节点
func NewSingleLinkNode[T any](value T) *SingleLinkNode[T] {
	return &SingleLinkNode[T]{value: value, next: nil}
}

// Value 返回数据
func (s *SingleLinkNode[T]) Value() T {
	return s.value
}
func (s *SingleLinkNode[T]) Next() *SingleLinkNode[T] {
	return s.next
}
func (s *SingleLinkNode[T]) Set(value T) *SingleLinkNode[T] {
	s.value = value
	return s
}

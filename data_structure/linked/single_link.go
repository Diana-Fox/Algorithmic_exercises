package linked

import (
	"fmt"
	"strings"
)

type SingleLink[T any] struct {
	head   *SingleLinkNode[T] //链表头指针
	length int                //链表长度
}

// 新建
func NewSingleLink[T any]() LinkedInterface[T] {
	var h T //泛型不让用nil
	head := NewSingleLinkNode[T](h)
	return &SingleLink[T]{head: head, length: 0}
}

func (list *SingleLink[T]) GetFirstNode() *SingleLinkNode[T] {
	return list.head.next
}

func (list *SingleLink[T]) InsertNodeFront(node *SingleLinkNode[T]) {
	if list.head == nil {
		list.head = node
		node.next = nil
	} else {
		bak := list.head     //备份头结点
		node.next = bak.next //头节点的下一个节点是当前节点的下一个节点
		bak.next = node      //头节点的第一个节点就是当前节点
	}
	list.length++
}

func (list *SingleLink[T]) InsertNodeBack(node *SingleLinkNode[T]) {
	p := list.head
	for p.next != nil {
		p = p.next
	}
	p.next = node
	list.length++
}

func (list *SingleLink[T]) InsertNodeValueFront(dest T, node *SingleLinkNode[T], f func(v1, v2 T) int) bool {
	p := list.head
	isFind := false
	for p.next != nil { //因为是泛型必须借助比较方法
		if f(p.next.value, dest) != 0 {
			isFind = true
			if f(p.next.value, dest) == 0 {
				isFind = true
			}
			break
		}
		p = p.next
	}
	if isFind {
		node.next = p.next
		p.next = node
		list.length++
	}
	return isFind
}

func (list *SingleLink[T]) InsertNodeValueBack(dest T, node *SingleLinkNode[T], f func(v1, v2 T) int) bool {
	p := list.head
	isFind := false
	for p.next != nil { //因为是泛型必须借助比较方法
		if f(p.next.value, dest) != 0 {
			isFind = true
			if f(p.value, dest) == 0 {
				isFind = true
			}
			break
		}
		p = p.next
	}
	if isFind {
		node.next = p.next
		p.next = node
		list.length++
	}
	return isFind
}

func (list *SingleLink[T]) GetNodeAtIndex(index int) *SingleLinkNode[T] {
	if index < 0 || index >= list.length {
		return nil
	}
	p := list.head
	for index > -1 {
		p = p.next
		index--
	}
	return p
}

func (list *SingleLink[T]) DeleteNode(dest *SingleLinkNode[T]) bool {
	if dest == nil {
		return false
	}
	p := list.head
	for p.next != nil && p.next != dest {
		p = p.next
	}
	if p.next == dest { //符合
		p.next = p.next.next
		list.length--
		return true
	}
	return false
}

func (list *SingleLink[T]) DeleteIndex(index int) bool {
	if index < 0 || index >= list.length {
		return false
	} else {
		p := list.head
		for index > 0 {
			p = p.next
			index--
		}
		p.next = p.next.next
		return true
	}
}

func (list *SingleLink[T]) String() string {
	var sb strings.Builder
	p := list.head.next
	for p.next != nil {
		sb.WriteString(fmt.Sprintf("%v,", p.value))
		p = p.next
	}
	sb.WriteString(fmt.Sprintf("%v", p.value))
	return sb.String() //打印字符串
}

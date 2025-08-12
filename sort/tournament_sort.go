package sort

import "math"

// 叶节点
type Node[T any] struct {
	val  T    //叶节点的数据
	isOk bool //叶子状态是不是无穷大
	rank int  //叶子的排序
}

// TournamentSort ======================
// 锦标赛排序
type TournamentSort[T any] struct {
}

func NewTournamentSort[T any]() *TournamentSort[T] {
	return &TournamentSort[T]{}
}
func (t *TournamentSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	if len(array) < 2 {
		return array
	}
	var level int //树的层数，默认
	for t.pow(2, level) < len(array) {
		level++ //求出可以覆盖所有元素的层数
	}
	var leaf = t.pow(2, level)             //叶节点的数量
	totalNodes := leaf*2 - 1               //总节点数
	var tree = make([]Node[T], totalNodes) //构造对应的树
	for i := 0; i < len(array); i++ {      //填充树的全部叶子节点
		tree[leaf-1+i] = Node[T]{ //叶节点的数量比非叶节点数量的和多1，所以是从leaf-1开始填充
			val:  array[i],
			isOk: true, //有值的情况下，非∞
			rank: i,    //在原数组中的位置
		}
	}
	for i := 0; i < level; i++ {
		nodeCount := t.pow(2, level-i) //每次处理降低一个层次
		for j := 0; j < nodeCount/2; j++ {
			leftNode := nodeCount - 1 + j*2
			rightNode := leftNode + 1
			middleNode := (leftNode - 1) / 2
			tree[middleNode] = t.compareNodes(tree[leftNode], tree[rightNode], f)
		}
	} //构造树
	result := make([]T, 0, leaf)
	result = append(result, tree[0].val)
	for i := 0; i < len(array)-1; i++ {
		winNode := tree[0].rank + leaf - 1 //记录赢的节点
		tree[winNode].isOk = false
		for j := 0; j < level; j++ { //最多处理这么多层
			leftNode := winNode
			if winNode%2 == 0 {
				leftNode = winNode - 1
			}
			rightNode := leftNode + 1
			middleNode := (leftNode - 1) / 2
			tree[middleNode] = t.compareNodes(tree[leftNode], tree[rightNode], f)
			winNode = (leftNode - 1) / 2
		}
		result = append(result, tree[winNode].val)
	}
	return result
}

// x的y次方，x和y均为整数
func (t *TournamentSort[T]) pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

// 比较两个节点的大小
func (t *TournamentSort[T]) compareNodes(l Node[T], r Node[T], f func(a T, b T) int) Node[T] {
	switch {
	case !l.isOk && !r.isOk: //无效叶子节点，随便返回一个就行
		return l
	case !l.isOk: //左节点无效，返回右节点
		return r
	case !r.isOk:
		return l
	default:
		if f(l.val, r.val) <= 0 {
			return l //左边小
		}
		return r
	}
}

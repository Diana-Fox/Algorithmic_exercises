package sort

// 堆排序
type HeapSort[T any] struct {
}

func NewHeapSort[T any]() *HeapSort[T] {
	return &HeapSort[T]{}
}

func (h *HeapSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	if len(array) < 2 {
		return array
	}
	length := len(array)
	for i := 0; i < length; i++ {
		lastMessLen := length - i //每次针对这个范围排序
		h.HeapSortMax(array, lastMessLen, f)
		if i < length { //把最大值置顶
			array[0], array[lastMessLen-1] = array[lastMessLen-1], array[0]
		}
	}
	return array
}

func (h *HeapSort[T]) HeapSortMax(array []T, len int, f func(a T, b T) int) []T {
	if len < 2 {
		return array
	}
	depth := len/2 - 1 //深度
	for i := depth; i >= 0; i-- {
		topMaxId := i                                                  //假设当前最大
		lChildId := 2*i + 1                                            //左孩子索引
		rChildId := 2*i + 2                                            //右孩子索引
		if lChildId < len && f(array[lChildId], array[topMaxId]) > 0 { //左边更大
			topMaxId = lChildId //当前最大是左边
		}
		if rChildId < len && f(array[rChildId], array[topMaxId]) > 0 {
			topMaxId = rChildId //右边更大
		}
		if topMaxId != i {
			array[i], array[topMaxId] = array[topMaxId], array[i]
		}
	}
	return array
}

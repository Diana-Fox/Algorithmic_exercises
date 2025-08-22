package search

type BinarySearch[T any] struct {
}

func NewBinarySearch[T any]() Search[T] {
	return &BinarySearch[T]{}
}

func (b *BinarySearch[T]) Search(array []T, value T, f func(a T, b T) int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if f(array[mid], value) == 0 {
			//找到了
			return mid
		} else if f(array[mid], value) > 0 { //中间值比查询值大，往左移动
			high = mid - 1
		} else {
			low = mid + 1 //中间值比查询值小，往右边移动
		}
	}
	return -1
}

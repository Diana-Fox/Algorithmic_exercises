package search

type BinaryPlusSearch[T any] struct {
}

func NewBinaryPlusSearch[T any]() Search[T] {
	return &BinaryPlusSearch[T]{}
}

func (b *BinaryPlusSearch[T]) Search(array []T, value T, f func(a T, b T) int) int {
	low := 0
	high := len(array) - 1
	index := -1
	for low <= high {
		mid := (low + high) / 2
		if f(array[mid], value) > 0 { //mid>vale
			high = mid - 1
		} else if f(array[mid], value) < 0 { //mid<value
			low = mid + 1
		} else {
			if mid == 0 || f(array[mid-1], value) != 0 { //如果是最前面的value值
				index = mid
				break
			} else { //不是最前面的value值
				high = mid - 1
			}
		}
	}
	return index
}

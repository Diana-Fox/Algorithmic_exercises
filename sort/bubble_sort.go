package sort

// BubbleSort 冒泡排序
type BubbleSort[T any] struct {
	//SelectSort(array []T) []T
}

func NewBubbleSort[T any]() *BubbleSort[T] {
	return &BubbleSort[T]{}
}
func (s *BubbleSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	if len(array) < 2 {
		return array
	}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if f(array[i], array[j]) > 0 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

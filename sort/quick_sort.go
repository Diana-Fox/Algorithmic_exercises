package sort

type QuickSort[T any] struct {
}

// NewQuickSort 快速排序
func NewQuickSort[T any]() Sort[T] {
	return &QuickSort[T]{}
}
func (q *QuickSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	if len(array) < 2 {
		return array
	}
	//处理
	splitData := array[0]     //选第一个数字做轴
	low := make([]T, 0, 0)    //比轴小的
	high := make([]T, 0, 0)   //比轴大的
	middle := make([]T, 0, 0) //和轴一样大的
	middle = append(middle, splitData)
	for i := 1; i < len(array); i++ {
		if f(splitData, array[i]) > 0 {
			low = append(low, array[i])
		} else if f(splitData, array[i]) < 0 {
			high = append(high, array[i])
		} else {
			middle = append(middle, array[i])
		}
	}
	low, high = q.Sort(low, f), q.Sort(high, f) //分别对两段排序
	result := append(append(low, middle...), high...)
	return result
}

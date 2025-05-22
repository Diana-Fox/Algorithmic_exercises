package sort

// 选择排序
type SelectSort[T any] struct {
	//SelectSort(array []T) []T
}

// 通过泛型的方式来进行选择排序
func (s *SelectSort[T]) Sort(array []T, f func(a T, b T) bool) []T {
	if len(array) < 2 {
		return array
	}
	for i := 0; i < len(array); i++ { //对第i位进行排序
		minx := i //假设每次最小值都是剩余未排序的第一个值
		for j := i + 1; j < len(array); j++ {
			if f(array[minx], array[j]) {
				minx = j
			}
		}
		if minx != i { //选出的最小值不是当前位置的值
			array[minx], array[i] = array[i], array[minx] //将当前位置的值与选中的最小值进行交换
		}
	}
	return array
}

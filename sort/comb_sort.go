package sort

type CombSort[T any] struct {
}

func NewCombSort[T any]() *CombSort[T] {
	return &CombSort[T]{}
}

// Sort 希尔排序的改良
func (c *CombSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	length := len(array) //数组长度
	gap := length
	for gap > 1 {
		gap = gap * 10 / 13
		for i := 0; i+gap < length; i++ { //不断收缩
			if f(array[i], array[i+gap]) > 0 {
				array[i], array[i+gap] = array[i+gap], array[i]
			}
		}
	}
	return array
}

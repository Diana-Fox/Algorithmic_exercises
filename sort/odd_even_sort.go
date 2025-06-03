package sort

// OddEvenSort 奇偶排序
type OddEvenSort[T any] struct {
}

func NewOddEvenSort[T any]() *OddEvenSort[T] {
	return &OddEvenSort[T]{}
}

// Sort 先确认是奇数还是偶数，然后再针对奇偶进行排序
func (o *OddEvenSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	if len(array) < 2 {
		return array
	}
	isSorted := false //是否有序 奇数位，偶数位都不需要排序的时候
	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1; i = i + 2 { //偶数位 所有的偶数位和身后的奇数位比较交换
			if f(array[i], array[i+1]) > 0 {
				array[i], array[i+1] = array[i+1], array[i] //互换
				isSorted = false
			}
		}
		for i := 1; i < len(array)-1; i = i + 2 { //奇数位 所有的奇数位和身后的偶数位比较交换
			if f(array[i], array[i+1]) > 0 {
				array[i], array[i+1] = array[i+1], array[i] //互换
				isSorted = false
			}
		}
	}
	return array
}

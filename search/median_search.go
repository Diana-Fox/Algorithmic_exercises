package search

// MedianSearch 中值查找
type MedianSearch[T any] struct {
}

func NewMedianSearch[T any]() Search[T] {
	return new(MedianSearch[T])
}

func (m *MedianSearch[T]) Search(array []T, value T, f func(a T, b T) int) int {
	low := 0
	high := len(array) - 1
	for low <= high {
		//计算一下两个值的差值（不算很严谨，因为非数字类型的大小比较可能非差值，标注一下）
		leftV := float64(f(value, array[low]))      //要查询的值和当前最小值的差距
		allV := float64(f(array[high], array[low])) //要查询值和当前最大值的差距
		diff := float64(high - low)                 //整个数组的长度
		mid := int(float64(low) + leftV/allV*diff)  //计算中间值比例,其实就是每次的比例发生变更
		if mid < 0 || mid >= len(array)-1 {         //越界了
			return -1
		}
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

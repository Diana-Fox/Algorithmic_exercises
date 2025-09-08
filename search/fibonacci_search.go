package search

// FibonacciSearch 斐波拉契查找
type FibonacciSearch[T any] struct {
}

func NewFibonacciSearch[T any]() Search[T] {
	return new(FibonacciSearch[T])
}

// Search 斐波拉契查找
func (fs *FibonacciSearch[T]) Search(array []T, value T, f func(a T, b T) int) int {
	length := len(array)
	fibArr := fs.makeFib(array)         //去得到斐波拉契数组
	fillLength := fibArr[len(fibArr)-1] //要填充的长度
	fillArr := make([]T, fillLength)
	for i, v := range array { //填充数组数据
		fillArr[i] = v
	}
	lastData := array[length-1] //原数组最后一位
	for i := length; i < fillLength; i++ {
		fillArr[i] = lastData //就超出部分都适用最大值即可
	}
	left, mid, right := 0, 0, length //重新开始二分查找
	kIndex := len(fibArr) - 1
	for left <= right {
		mid = left + fibArr[kIndex-1] - 1 //斐波拉契切割
		if f(value, fillArr[mid]) < 0 {   //小于中间值
			right = mid - 1
			kIndex--
		} else if f(value, fillArr[mid]) > 0 {
			left = mid + 1
			kIndex -= 2 //挪两个
		} else {
			if mid > right { //越界了
				return right
			}
			return mid
		}
	}

	return -1
}

// 数组可能不构成斐波拉契，所以要修饰处理一下
func (fs *FibonacciSearch[T]) makeFib(array []T) []int {
	length := len(array) //数组长度
	flbLen := 2
	first, second, third := 1, 2, 3
	for third < length { //构造出一个最近的斐波拉契数组
		third, first, second = first+second, second, third //叠加计算斐波拉契
		flbLen++
	}
	fb := make([]int, flbLen)
	fb[0] = 1
	fb[1] = 1
	for i := 2; i < flbLen; i++ {
		fb[i] = fb[i-1] + fb[i-2] //叠加计算斐波拉契值
	}
	return fb
}

package sort

import "math/rand"

// QuickSortPlus 快速排序改进版
type QuickSortPlus[T any] struct {
}

func NewQuickSortPlus[T any]() Sort[T] {
	return new(QuickSortPlus[T])
}

// Sort 快速排序改进版
func (q *QuickSortPlus[T]) Sort(array []T, f func(a T, b T) int) []T {
	sortArr := q.sort(array, f, 0, len(array)-1)
	return sortArr
}
func (q *QuickSortPlus[T]) sort(array []T, f func(a T, b T) int, left int, right int) []T {
	if right-left < 2 {
		//q.sortForMerge(array, f, left, right) //插入排序就行
		return array
	} else {
		//随机找一个数字
		j := rand.Int()%(right-left+1) + left
		array[left], array[j] = array[j], array[left] //选一个随机数，放到首位
		vData := array[left]                          //坐标数组，比我小，左边，比我大右边
		lt := left                                    //arr[left+1,lt]<vData
		gt := right + 1                               //arr[gt...right]>vData
		i := left + 1                                 //arr[lt+1,...i]==vData
		for i < gt {
			if f(array[i], vData) < 0 { //当前值小于坐标
				array[i], array[lt+1] = array[lt+1], array[i] //移动到小于的位置
				lt++
				i++
			} else if f(array[i], vData) > 0 { //当前值大于坐标
				array[i], array[gt-1] = array[gt-1], array[i]
				gt--
			} else {
				i++
			}
		}
		array[lt], array[left] = array[left], array[lt] //交换一下轴的位置
		q.sort(array, f, left, lt-1)                    //递归处理前半段
		q.sort(array, f, gt, right)                     //递归处理后半段
	}
	return array
}

func (q *QuickSortPlus[T]) sortForMerge(array []T, f func(a T, b T) int, left int, right int) {
	for i := left; i < right+1; i++ {
		temp := array[i] //备份一下
		var j int
		for j = i; j > left && f(array[j-1], temp) > 0; j-- {
			array[j] = array[j-1]
		}
		array[j] = temp
	}
}

package int_sort

type RadixSort struct {
}

// NewRadixSort 桶排序
func NewRadixSort() *RadixSort {
	return new(RadixSort)
}

// Sort 基数排序-桶排序,只适合数字
func (r *RadixSort) Sort(array []int) []int {
	maxValue := r.getMax(array) //根据最大值确定范围
	for bit := 1; maxValue/bit > 0; bit *= 10 {
		r.bitSort(array, bit) //每次处理一个级别的排序,bit是1的时候，处理个位数，10的时候处理10位数，以此类推
	}
	return array
}

// 找到最大值
func (r *RadixSort) getMax(array []int) int {
	if len(array) < 2 {
		return array[0]
	}
	maxValue := array[0]
	for i := 1; i < len(array); i++ {
		if maxValue < array[i] {
			maxValue = array[i]
		}
	}
	return maxValue
}

func (r *RadixSort) bitSort(array []int, bit int) []int {
	length := len(array)
	bitCounts := make([]int, 10) //统计长度，也就是桶数量
	for i := 0; i < length; i++ {
		num := (array[i] / bit) % 10
		bitCounts[num]++ //统计余数相等的个数
	}
	for i := 1; i < 10; i++ {
		bitCounts[i] += bitCounts[i-1] //叠加，计算位置，进行一个赋值，这样，第i位减去i-1位，就可知道当前位置有几个值
	}
	tmp := make([]int, 10)
	for i := length - 1; i >= 0; i-- { //恢复一下
		num := (array[i] / bit) % 10     //
		tmp[bitCounts[num]-1] = array[i] //计算排序的位置，其实就是看当前值前面要在temp的第几位，根据当前bitCounts[num]的到前面要空出几个位置
		bitCounts[num]--
	}
	for i := 0; i < length; i++ {
		array[i] = tmp[i] //保存数组
	}
	return array
}

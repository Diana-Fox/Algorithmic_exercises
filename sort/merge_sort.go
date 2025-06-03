package sort

// MergeSort 归并排序
type MergeSort[T any] struct {
}

func NewMergeSort[T any]() *MergeSort[T] {
	return &MergeSort[T]{}
}
func (m *MergeSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	if len(array) < 2 { //递归基
		return array
	}
	mid := len(array) / 2                //先切分
	leftArr := m.Sort(array[:mid], f)    //0到中间是左边，继续分割
	rightArr := m.Sort(array[mid:], f)   //中间到结尾是右边，继续分割
	return m.merge(leftArr, rightArr, f) //最后进行合并
}
func (m *MergeSort[T]) merge(lArray []T, rArray []T, f func(a T, b T) int) []T {
	leftIndex := 0
	rightIndex := 0
	arr := []T{}                                              //要返回的集合
	for leftIndex < len(lArray) && rightIndex < len(rArray) { //两边都没越界
		if f(lArray[leftIndex], rArray[rightIndex]) < 0 { //取左边
			arr = append(arr, lArray[leftIndex])
			leftIndex++
		} else if f(lArray[leftIndex], rArray[rightIndex]) > 0 { //取右边
			arr = append(arr, rArray[rightIndex])
			rightIndex++
		} else { //两变都要
			arr = append(arr, lArray[leftIndex])
			leftIndex++
			arr = append(arr, rArray[rightIndex])
			rightIndex++
		}
	}
	//左边，剩余部分
	for leftIndex < len(lArray) {
		arr = append(arr, lArray[leftIndex])
		leftIndex++
	}
	//右边，剩余部分
	for rightIndex < len(rArray) {
		arr = append(arr, rArray[rightIndex])
		rightIndex++
	}
	return arr
}

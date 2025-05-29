package sort

// InsertSort     插入排序
type InsertSort[T any] struct {
	//SelectSort(array []T) []T
}

func NewInsertSort[T any]() *InsertSort[T] {
	return &InsertSort[T]{}
}
func (s *InsertSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	if len(array) < 2 {
		return array
	}
	for i := 1; i < len(array); i++ {
		backup := array[i]                      //备份插入的数据
		j := i - 1                              //要插入的位置在这
		for j >= 0 && f(backup, array[j]) < 0 { //backup更小
			array[j+1] = array[j] //往后挪动一位
			j--                   //去下一个位置
		}
		array[j+1] = backup
	}
	return array
}

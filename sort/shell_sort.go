package sort

type ShellSort[T any] struct {
}

func NewShellSort[T any]() *ShellSort[T] {
	return &ShellSort[T]{}
}

// Sort 希尔一般用在并发场合
func (s *ShellSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	if len(array) < 2 {
		return array
	}
	gap := len(array) / 2 //步长
	for gap > 0 {         //
		for i := 0; i < gap; i++ { //做个排序
			s.ShellSortStep(array, i, gap, f)
		}
		gap = gap / 2 //步长每次缩小一半
	}
	return array
}
func (s *ShellSort[T]) ShellSortStep(arr []T, start int, gap int, f func(a T, b T) int) []T {
	for i := start + gap; i < len(arr); i += gap { //编写一个插入排序
		backup := arr[i]
		j := i - gap
		for j >= 0 && f(backup, arr[j]) < 0 {
			arr[j+gap] = arr[j]
			j -= gap
		}
		arr[j+gap] = backup
	}
	return arr
}

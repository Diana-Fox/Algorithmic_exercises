package sort

// Sort 排序的接口
type Sort[T any] interface {
	Sort(array []T, f func(a T, b T) bool) []T //
}

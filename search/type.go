package search

type Search[T any] interface {
	Search(array []T, value T, f func(a T, b T) int) int //要查询的值
}

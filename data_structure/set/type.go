package set

// SetInterface set接口,便于以后扩展
type SetInterface[T any] interface {
	Add(value T) bool
	Strings() []T
}

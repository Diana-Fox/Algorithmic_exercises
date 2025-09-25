package sort

// GnomeSort 侏儒排序
type GnomeSort[T any] struct {
}

func NewGnomeSort[T any]() *GnomeSort[T] {
	return &GnomeSort[T]{}
}

func (g *GnomeSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	i := 1
	for i < len(array) {
		if f(array[i], array[i-1]) > 0 {
			i++ //符合顺序
		} else {
			array[i], array[i-1] = array[i-1], array[i] //交换
			if i > 1 {                                  //回到前一个位置，看当前这个位置的情况
				i--
			}
		}
	}
	return array
}

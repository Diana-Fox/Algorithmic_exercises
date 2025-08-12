package sort

// CocktailSort 鸡尾酒排序
type CocktailSort[T any] struct {
}

func NewCocktailSort[T any]() Sort[T] {
	return &CocktailSort[T]{}
}

// Sort 鸡尾酒排序，核心就是左边一下，右边一下
func (c *CocktailSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	for i := 0; i < len(array)/2; i++ {
		left := 0               //左指针
		right := len(array) - 1 //右指针
		for left < right {
			if f(array[left], array[left+1]) > 0 {
				array[left], array[left+1] = array[left+1], array[left] //交换一下
			}
			left++
			if f(array[right-1], array[right]) > 0 {
				array[right], array[right-1] = array[right-1], array[right]
			}
			right--
		}
	}
	return array
}

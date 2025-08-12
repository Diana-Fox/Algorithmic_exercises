package int_sort

import (
	"fmt"
	"testing"
)

func TestRadixSort_Sort(t *testing.T) {
	arr := []int{5, 8, 13, 19, 21}
	fmt.Println(NewRadixSort().Sort(arr))
}

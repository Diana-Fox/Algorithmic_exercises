package int_sort

import (
	"fmt"
	"testing"
)

func TestRadixSort_Sort(t *testing.T) {
	//type args struct {
	//	array []int
	//}
	//arr := []int{10, 9, 23, 43, 40, 30, 14, 25}
	arr := []int{5, 8, 13, 19, 21}
	//want := []int{9, 10, 23, 30, 40, 43}
	fmt.Println(NewRadixSort().Sort(arr))
	//tests := []struct {
	//	name string
	//	args args
	//	want []int
	//}{
	//	{name: "桶排序", args: args{array: arr}, want: want},
	//}
	//sort := NewRadixSort()
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		fmt.Println(sort.Sort(tt.args.array))
	//	})
	//}
}

package sort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	//参数实体
	type args[T any] struct {
		array []T
		f     func(a T, b T) int
	}
	type testCase[T any] struct {
		name string
		s    Sort[T]
		args args[T]
		want []T
	}
	arr := []int{10, 9, 23, 43, 40, 30}
	want := []int{9, 10, 23, 30, 40, 43}
	tests := []testCase[int]{
		//{name: "选择排序", s: NewSelectSort[int_sort](), args: args[int_sort]{arr, func(a int_sort, b int_sort) int_sort {
		//	return a - b
		//}}, want: want},
		//{name: "插入排序", s: NewInsertSort[int_sort](), args: args[int_sort]{arr, func(a int_sort, b int_sort) int_sort {
		//	return a - b
		//}}, want: want},
		//{name: "冒泡排序", s: NewBubbleSort[int_sort](), args: args[int_sort]{arr, func(a int_sort, b int_sort) int_sort {
		//	return a - b
		//}}, want: want},
		//{name: "堆排序", s: NewHeapSort[int_sort](), args: args[int_sort]{arr, func(a int_sort, b int_sort) int_sort {
		//	return a - b
		//}}, want: want},
		//{name: "奇偶排序", s: NewOddEvenSort[int_sort](), args: args[int_sort]{arr, func(a int_sort, b int_sort) int_sort {
		//	return a - b
		//}}, want: want},
		//{name: "归并排序", s: NewMergeSort[int_sort](), args: args[int_sort]{arr, func(a int_sort, b int_sort) int_sort {
		//	return a - b
		//}}, want: want},
		//{name: "希尔排序", s: NewShellSort[int_sort](), args: args[int_sort]{arr, func(a int_sort, b int_sort) int_sort {
		//	return a - b
		//}}, want: want},
		//{name: "锦标赛排序", s: NewTournamentSort[int](), args: args[int]{arr, func(a int, b int) int {
		//	return a - b
		//}}, want: want},
		//{name: "鸡尾酒排序", s: NewCocktailSort[int](), args: args[int]{arr, func(a int, b int) int {
		//	return a - b
		//}}, want: want},
		{name: "快速排序", s: NewQuickSort[int](), args: args[int]{arr, func(a int, b int) int {
			return a - b
		}}, want: want},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Sort(tt.args.array, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name=%v Sort() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

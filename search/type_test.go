package search

import (
	"reflect"
	"testing"
)

// 生成测试的方法
func TestSearch(t *testing.T) {
	type args[T any] struct {
		array []T
		f     func(a T, b T) int
	}
	type testCase[T any] struct {
		name string
		s    Search[T]
		args args[T]
		want int
	}
	arr := []int{9, 10, 23, 30, 40, 43} //用有序的数组
	tests := []testCase[int]{
		{
			name: "二分查找",
			s:    NewBinarySearch[int](),
			args: args[int]{arr, func(a int, b int) int {
				return a - b
			}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Search(tt.args.array, 23, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("name=%v Sort() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

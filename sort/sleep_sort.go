package sort

import "time"

// SleepSort 休眠排序
type SleepSort[T any] struct {
	flag      bool   //标识
	container chan T //管道
	count     int
}

func NewSleepSort[T any](flag bool, container chan T) *SleepSort[T] {
	return &SleepSort[T]{
		flag:      flag,
		container: container,
		count:     0,
	}
}

func (s *SleepSort[T]) Sort(array []T, f func(a T, b T) int) []T {
	for i := 0; i < len(array); i++ {
		go s.toSleep(array[i], f)
	}
	go s.listen(array)
	for s.flag {
		time.Sleep(time.Millisecond * 10000)
	}
	return array
}

// 睡眠多久
func (s *SleepSort[T]) toSleep(value T, f func(a T, b T) int) {
	data := f(value, value)                            //得到应该休眠的次数，本处f应该与其他的f不同，b用不上，但是为了保持接口一致性，故如此
	time.Sleep(time.Duration(data) * time.Millisecond) //todo 注意，如果两个数大小相近，可能会因为系统调度的原因导致排序出错
	s.container <- value                               //管道输入
}

// 监听
func (s *SleepSort[T]) listen(array []T) {
	size := len(array)
	for s.flag {
		select {
		case value := <-s.container:
			array[s.count] = value
			s.count++ //计数器
			if s.count >= size {
				s.flag = false
				break
			}
		}
	}
}

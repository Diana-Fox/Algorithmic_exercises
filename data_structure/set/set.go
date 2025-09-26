package set

type Set[T any] struct {
	buf  []T          //实际数据
	num  int          //数据长度
	hash map[any]bool //由于官方map不允许使用泛型，只能使用any了
}

func (s *Set[T]) Strings() []T {
	return s.buf
}

// 新建可变set
func NewSet[T any]() SetInterface[T] {
	return &Set[T]{
		buf:  []T{},
		num:  0,
		hash: make(map[any]bool),
	}
}

func (s *Set[T]) Add(value T) bool {
	if s.isExist(value) {
		return false
	}
	s.buf = append(s.buf, value)
	s.num++
	s.hash[value] = true
	return true
}

func (s *Set[T]) isExist(value T) bool {
	return s.hash[value]
}

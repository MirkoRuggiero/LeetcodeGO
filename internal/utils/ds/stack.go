package ds

type Stack[T any] struct {
	Push   func(T)
	Peek   func() T
	Pop    func() T
	Length func() int
}

func NewStack[T any]() Stack[T] {
	slice := make([]T, 0)
	return Stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Peek: func() T {
			return slice[len(slice)-1]
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
	}
}

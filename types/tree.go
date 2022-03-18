package types

type Tree[T any] struct {
	Left  *Tree[T]
	Right *Tree[T]
	Value T
}

func (t Tree[T]) GetValue() T {
	return t.Value
}

func (t Tree[_]) Count() int {
	acc := 1
	if t.Left != nil {
		acc = acc + t.Left.Count()
	}
	if t.Right != nil {
		acc = acc + t.Left.Count()
	}

	return acc
}

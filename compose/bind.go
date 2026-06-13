package compose

type Bind[T any] interface {
	Get() T
	Set(v T)
}

type Setter[T any] func(v T)
type Getter[T any] func() T
type BindObj[T any] struct {
	value  T
	setter Setter[T]
	getter Getter[T]
}

func (b *BindObj[T]) Get() T {
	if b.getter != nil {
		return b.getter()
	}
	return b.value
}

func (b *BindObj[T]) Set(v T) {
	b.value = v
	if b.setter != nil {
		b.setter(b.value)
	}
}

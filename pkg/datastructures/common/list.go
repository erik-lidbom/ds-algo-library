package common

type List[T any] interface {
	Collection
	Add(index int, elem T) error
	Get(index int) (T, error)
	Set(index int, elem T) error
	Remove(index int) (T, error)
}

package common

type Set[T any] interface {
	Collection
	Add(elem T) error
	Remove(elem T) (T, error)
	Contains(elem T) bool
}
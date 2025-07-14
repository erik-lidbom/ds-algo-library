package common

type List interface {
	Collection
	Add(index int, elem any) (error)
	Get(index int) (any, error)
	Set(index int, elem any) (error)
	Remove(index int) (any, error)
}
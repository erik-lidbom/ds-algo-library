package common

type Stack interface {
	Collection
	Push(x any)
	Pop() (any, error)
	Peek() (any, error)
}
package common

type Tree interface {
	Collection
	Insert(item any) error
	Search(item any) (any, bool)
	Delete(item any) error
	TraversePreOrder() ([]any, error)
	TraversePostOrder() ([]any, error)
	TraverseInOrder() ([]any, error)
	Clear() error
}
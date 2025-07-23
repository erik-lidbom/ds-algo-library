package common

type Map[K comparable, V any] interface {
	Collection
	Put(key K, value V)
	Get(key K) (V, error)
	Remove(key K) (V, error)
	ContainsKey(key K) bool
}
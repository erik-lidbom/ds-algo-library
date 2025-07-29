package common

import "ds-algorithms/pkg/datastructures/array"

type Map[K comparable, V any] interface {
	Collection
	Put(key K, value V)
	Get(key K) (V, error)
	Remove(key K) (V, error)
	ContainsKey(key K) bool
}

type SortedMap[K comparable, V any] interface {
	Map[K, V]
	FirstKey() (K, error)
	LastKey() (K, error)
	FloorKey(key K) (K, error)
	CeilingKey(key K) (K, error)
	LowerKey(key K) (K, error)
	HigherKey(key K) (K, error)
	KeysBetween(key1, key2 K) *array.ArrayList[K]
}

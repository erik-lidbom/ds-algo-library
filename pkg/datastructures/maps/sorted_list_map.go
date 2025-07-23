package maps

import (
	"cmp"
	search "ds-algorithms/pkg/algorithms/searching"
	"ds-algorithms/pkg/datastructures/array"
	"fmt"
)

type SortedListMap[K cmp.Ordered, V comparable] struct {
	keys *array.ArrayList[K]
	values *array.ArrayList[V]
	size int
}

func NewSortedListMap[K cmp.Ordered, V comparable]() *SortedListMap[K, V] {
	return &SortedListMap[K, V]{
		keys: array.NewArrayList[K](),
		values: array.NewArrayList[V](),
		size: 0,
	}
}

func (slm *SortedListMap[K, V]) Size() int {
	return slm.size
}

func (slm *SortedListMap[K, V]) IsEmpty() bool {
	return slm.size == 0
}

func (slm *SortedListMap[K, V]) Put(key K, value V) {

	duplicatedVal, index := search.FindInsertionPoint(slm.keys, key)
	
	if duplicatedVal {
		slm.values.Set(index, value)
		return
	}

	slm.keys.Add(index, key)	
	slm.values.Add(index, value)
	slm.size++	
}

func (slm *SortedListMap[K, V]) Get(key K) (V, error) {
	var zero V

	containsKey, index := search.BinarySearchArrayList(slm.keys, key)
	if containsKey {
		val, _ := slm.values.Get(index)
		return val, nil
	}

	return zero, fmt.Errorf("cannot retrieve value: key '%v' not found in the map", key)
}

func (slm *SortedListMap[K, V]) Remove(key K) (V, error) {
	var zero V

	containsKey, index := search.BinarySearchArrayList(slm.keys, key)

	if containsKey {
		slm.keys.Remove(index)
		removedVal, _ := slm.values.Remove(index)
		slm.size--
		return removedVal, nil
	}
	
	return zero, fmt.Errorf("key %v not found in the map", key)
}

func (slm *SortedListMap[K, V]) ContainsKey(key K) bool {
	res, _ := search.BinarySearchArrayList(slm.keys, key)
	return res
}

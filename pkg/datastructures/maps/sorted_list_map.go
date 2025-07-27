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

func (slm *SortedListMap[K, V]) FirstKey() (K, error) {
	var zero K

	if slm.IsEmpty() {
		return zero, fmt.Errorf("map is empty")
	}

	key, err := slm.keys.Get(0)
	if err != nil {
		return zero, fmt.Errorf("internal error: failed to retrieve first key: %w", err)
	}

	return key, nil
}

func (slm *SortedListMap[K, V]) LastKey() (K, error) {
	var zero K

	if slm.IsEmpty() {
		return zero, fmt.Errorf("map is empty")
	}

	key, err := slm.keys.Get(slm.Size() - 1)
	if err != nil {
		return zero, fmt.Errorf("internal error: failed to retrieve last key: %w", err)
	}

	return key, nil
}

func (slm *SortedListMap[K, V]) FloorKey(key K) (K, error) {
	var zero K

	if slm.IsEmpty() {
		return zero, fmt.Errorf("map is empty")
	}

	found, index := search.FindInsertionPoint(slm.keys, key)

	if found {
		floorKey, err := slm.keys.Get(index)
		if err != nil {
			return zero, fmt.Errorf("internal error: failed to retrieve floorKey at index %d: %w", index, err)
		}

		return floorKey, nil
	}

	if index == 0 {
		return zero, fmt.Errorf("no element <= %v found in map", key)
	}

	floorKey, err := slm.keys.Get(index - 1)
	if err != nil {
        return zero, fmt.Errorf("internal error: failed to retrieve floorKey at index %d-1: %w", index, err)
    }
	return floorKey, nil
}

func (slm *SortedListMap[K, V]) CeilingKey(key K) (K, error) {
	var zero K

	if slm.IsEmpty() {
		return zero, fmt.Errorf("map is empty")
	}

	_, index := search.FindInsertionPoint(slm.keys, key)

	if index == slm.Size() {
		return zero, fmt.Errorf("no key >= %v found in map", key)
	}

	ceilingKey, err := slm.keys.Get(index)
	if err != nil {
        return zero, fmt.Errorf("internal error: failed to retrieve ceilingKey at index %d: %w", index, err)
    }

	return ceilingKey, nil
}

func (slm *SortedListMap[K, V]) LowerKey(key K) (K, error) {
	var zero K

	if slm.IsEmpty() {
		return zero, fmt.Errorf("map is empty")
	}

	_, index := search.FindInsertionPoint(slm.keys, key)
	if index == 0 {
		return zero, fmt.Errorf("no key < %v found in map", key)
	}

	lowerKey, err := slm.keys.Get(index - 1)
	if err != nil {
		return zero, fmt.Errorf("internal error: failed to retrieve lower key at index %d: %w", index-1, err)
	}

	return lowerKey, nil
}

func (slm *SortedListMap[K, V]) HigherKey(key K) (K, error) {
	var zero K

	if slm.IsEmpty() {
		return zero, fmt.Errorf("map is empty")
	}

	_, index := search.FindUpperBound(slm.keys, key)

	if index == slm.Size() {
		return zero, fmt.Errorf("no element > %v found in map", key)
	}

	higherKey, err := slm.keys.Get(index)
	if err != nil {
		return zero, fmt.Errorf("internal error: failed to retrieve higherKey at index %d: %w", index-1, err)
	}

	return higherKey, nil
}

func (slm *SortedListMap[K, V]) KeysBetween(key1, key2 K) *array.ArrayList[K] {
	arr := array.NewArrayList[K]()

	if slm.IsEmpty() {
		return arr
	}

	_, startIndex := search.FindInsertionPoint(slm.keys, key1)
	_, endIndex := search.FindUpperBound(slm.keys, key2)

	if startIndex >= slm.Size() || startIndex >= endIndex {
		return arr
	}

	for i := startIndex; startIndex < endIndex; i++ {
		currVal, _ := slm.keys.Get(i)
		arr.Add(i, currVal)
	}
	return arr
}

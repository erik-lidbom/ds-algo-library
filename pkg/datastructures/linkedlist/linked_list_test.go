package linkedlist

import (
	"testing"
)

func TestLinkedList_NewAndBasic(t *testing.T) {
	ll := &LinkedList{}
	if ll.Size() != 0 {
		t.Errorf("Expected size 0, got %d", ll.Size())
	}
	if !ll.IsEmpty() {
		t.Errorf("Expected list to be empty")
	}
}

func TestLinkedList_AddGetSetRemove(t *testing.T) {
	ll := &LinkedList{}
	// Add at head
	err := ll.Add(0, 10)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Add at tail
	err = ll.Add(1, 20)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Add in middle
	err = ll.Add(1, 15)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if ll.Size() != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size())
	}
	// Get
	val, err := ll.Get(0)
	if err != nil || val != 10 {
		t.Errorf("Expected 10 at index 0, got %v (err: %v)", val, err)
	}
	val, err = ll.Get(1)
	if err != nil || val != 15 {
		t.Errorf("Expected 15 at index 1, got %v (err: %v)", val, err)
	}
	val, err = ll.Get(2)
	if err != nil || val != 20 {
		t.Errorf("Expected 20 at index 2, got %v (err: %v)", val, err)
	}
	// Set
	err = ll.Set(1, 17)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	val, _ = ll.Get(1)
	if val != 17 {
		t.Errorf("Expected 17 at index 1 after set, got %v", val)
	}
	// Remove middle
	val, err = ll.Remove(1)
	if err != nil || val != 17 {
		t.Errorf("Expected to remove 17 at index 1, got %v (err: %v)", val, err)
	}
	// Remove head
	val, err = ll.Remove(0)
	if err != nil || val != 10 {
		t.Errorf("Expected to remove 10 at index 0, got %v (err: %v)", val, err)
	}
	// Remove tail
	val, err = ll.Remove(0)
	if err != nil || val != 20 {
		t.Errorf("Expected to remove 20 at index 0, got %v (err: %v)", val, err)
	}
	if !ll.IsEmpty() {
		t.Errorf("Expected list to be empty after removals")
	}
}

func TestLinkedList_Errors(t *testing.T) {
	ll := &LinkedList{}
	// Add out of bounds
	err := ll.Add(1, 1)
	if err == nil {
		t.Errorf("Expected error for add out of bounds")
	}
	// Get out of bounds
	_, err = ll.Get(0)
	if err == nil {
		t.Errorf("Expected error for get out of bounds")
	}
	// Set out of bounds
	err = ll.Set(0, 1)
	if err == nil {
		t.Errorf("Expected error for set out of bounds")
	}
	// Remove out of bounds
	_, err = ll.Remove(0)
	if err == nil {
		t.Errorf("Expected error for remove out of bounds")
	}
}

func BenchmarkLinkedList_Add(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ll := &LinkedList{}
		for i := 0; i < 1000; i++ {
			ll.Add(ll.Size(), i)
		}
	}
}

func BenchmarkLinkedList_Get(b *testing.B) {
	ll := &LinkedList{}
	for i := 0; i < 1000; i++ {
		ll.Add(ll.Size(), i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ll.Get(n % 1000)
	}
}

func BenchmarkLinkedList_Set(b *testing.B) {
	ll := &LinkedList{}
	for i := 0; i < 1000; i++ {
		ll.Add(ll.Size(), i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ll.Set(n%1000, n)
	}
}

func BenchmarkLinkedList_Remove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ll := &LinkedList{}
		for i := 0; i < 1000; i++ {
			ll.Add(ll.Size(), i)
		}
		for i := 0; i < 1000; i++ {
			ll.Remove(0)
		}
	}
}

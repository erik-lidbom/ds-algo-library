package linkedlist

import (
	"testing"
)

func TestDoubleDeque_NewAndBasic(t *testing.T) {
	dd := &DoubleDeque{}
	if dd.Size() != 0 {
		t.Errorf("Expected size 0, got %d", dd.Size())
	}
	if !dd.IsEmpty() {
		t.Errorf("Expected deque to be empty")
	}
}

func TestDoubleDeque_AddRemovePeek(t *testing.T) {
	dd := &DoubleDeque{}
	// AddFirst, AddLast
	dd.AddFirst(1)
	dd.AddLast(2)
	dd.AddFirst(0)
	if dd.Size() != 3 {
		t.Errorf("Expected size 3, got %d", dd.Size())
	}
	// PeekFirst, PeekLast
	val, err := dd.PeekFirst()
	if err != nil || val != 0 {
		t.Errorf("Expected to peekFirst 0, got %v (err: %v)", val, err)
	}
	val, err = dd.PeekLast()
	if err != nil || val != 2 {
		t.Errorf("Expected to peekLast 2, got %v (err: %v)", val, err)
	}
	// RemoveFirst
	val, err = dd.RemoveFirst()
	if err != nil || val != 0 {
		t.Errorf("Expected to removeFirst 0, got %v (err: %v)", val, err)
	}
	// RemoveLast
	val, err = dd.RemoveLast()
	if err != nil || val != 2 {
		t.Errorf("Expected to removeLast 2, got %v (err: %v)", val, err)
	}
	// RemoveLast (now only 1 left)
	val, err = dd.RemoveLast()
	if err != nil || val != 1 {
		t.Errorf("Expected to removeLast 1, got %v (err: %v)", val, err)
	}
	if !dd.IsEmpty() {
		t.Errorf("Expected deque to be empty after removals")
	}
}

func TestDoubleDeque_Errors(t *testing.T) {
	dd := &DoubleDeque{}
	_, err := dd.RemoveFirst()
	if err == nil {
		t.Errorf("Expected error for removeFirst on empty deque")
	}
	_, err = dd.RemoveLast()
	if err == nil {
		t.Errorf("Expected error for removeLast on empty deque")
	}
	_, err = dd.PeekFirst()
	if err == nil {
		t.Errorf("Expected error for peekFirst on empty deque")
	}
	_, err = dd.PeekLast()
	if err == nil {
		t.Errorf("Expected error for peekLast on empty deque")
	}
}

func BenchmarkDoubleDeque_AddFirst(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dd := &DoubleDeque{}
		for i := 0; i < 1000; i++ {
			dd.AddFirst(i)
		}
	}
}

func BenchmarkDoubleDeque_AddLast(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dd := &DoubleDeque{}
		for i := 0; i < 1000; i++ {
			dd.AddLast(i)
		}
	}
}

func BenchmarkDoubleDeque_RemoveFirst(b *testing.B) {
	dd := &DoubleDeque{}
	for i := 0; i < 1000; i++ {
		dd.AddLast(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dd.RemoveFirst()
	}
}

func BenchmarkDoubleDeque_RemoveLast(b *testing.B) {
	dd := &DoubleDeque{}
	for i := 0; i < 1000; i++ {
		dd.AddLast(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dd.RemoveLast()
	}
}

func BenchmarkDoubleDeque_PeekFirst(b *testing.B) {
	dd := &DoubleDeque{}
	for i := 0; i < 1000; i++ {
		dd.AddLast(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dd.PeekFirst()
	}
}

func BenchmarkDoubleDeque_PeekLast(b *testing.B) {
	dd := &DoubleDeque{}
	for i := 0; i < 1000; i++ {
		dd.AddLast(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		dd.PeekLast()
	}
}

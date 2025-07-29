package linkedlist

import (
	"fmt"
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	sl := NewSinglyLinkedList()
	if sl.Size() != 0 {
		t.Errorf("Expected size 0, got %d", sl.Size())
	}
	if !sl.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}

func TestSinglyLinkedList_PushPop(t *testing.T) {
	sl := NewSinglyLinkedList()
	sl.Push(1)
	sl.Push(2)
	sl.Push(3)

	if sl.Size() != 3 {
		t.Errorf("Expected size 3, got %d", sl.Size())
	}
	val, err := sl.Pop()

	if err != nil || val != 3 {
		t.Errorf("Expected to pop 3, got %v (err: %v)", val, err)
	}
	val, err = sl.Pop()

	if err != nil || val != 2 {
		fmt.Println("bam bam")

		t.Errorf("Expected to pop 2, got %v (err: %v)", val, err)
	}
	val, err = sl.Pop()
	if err != nil || val != 1 {
		t.Errorf("Expected to pop 1, got %v (err: %v)", val, err)
	}
	if !sl.IsEmpty() {
		t.Errorf("Expected stack to be empty after pops")
	}
}

func TestSinglyLinkedList_PopEmpty(t *testing.T) {
	sl := NewSinglyLinkedList()
	_, err := sl.Pop()
	if err == nil {
		t.Errorf("Expected error when popping empty stack")
	}
}

func TestSinglyLinkedList_Peek(t *testing.T) {
	sl := NewSinglyLinkedList()
	sl.Push("a")
	sl.Push("b")
	val, err := sl.Peek()
	if err != nil || val != "b" {
		t.Errorf("Expected to peek 'b', got %v (err: %v)", val, err)
	}
	if sl.Size() != 2 {
		t.Errorf("Peek should not change size")
	}
}

func TestSinglyLinkedList_PeekEmpty(t *testing.T) {
	sl := NewSinglyLinkedList()
	_, err := sl.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking empty stack")
	}
}

func BenchmarkSinglyLinkedList_Push(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := NewSinglyLinkedList()
		for i := 0; i < 1000; i++ {
			sl.Push(i)
		}
	}
}

func BenchmarkSinglyLinkedList_Pop(b *testing.B) {
	sl := NewSinglyLinkedList()
	for i := 0; i < 1000; i++ {
		sl.Push(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sl.Pop()
	}
}

func BenchmarkSinglyLinkedList_Peek(b *testing.B) {
	sl := NewSinglyLinkedList()
	for i := 0; i < 1000; i++ {
		sl.Push(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sl.Peek()
	}
}

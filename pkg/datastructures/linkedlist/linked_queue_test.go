package linkedlist

import (
	"testing"
)

func TestNewLinkedQueue(t *testing.T) {
	q := NewLinkedQueue()
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}
}

func TestLinkedQueue_EnqueueDequeue(t *testing.T) {
	q := NewLinkedQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
	// Dequeue returns the node, so check .val
	node, err := q.Dequeue()
	if err != nil || node.(*Node).val != 1 {
		t.Errorf("Expected to dequeue 1, got %v (err: %v)", node, err)
	}
	node, err = q.Dequeue()
	if err != nil || node.(*Node).val != 2 {
		t.Errorf("Expected to dequeue 2, got %v (err: %v)", node, err)
	}
	node, err = q.Dequeue()
	if err != nil || node.(*Node).val != 3 {
		t.Errorf("Expected to dequeue 3, got %v (err: %v)", node, err)
	}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty after dequeues")
	}
}

func TestLinkedQueue_DequeueEmpty(t *testing.T) {
	q := NewLinkedQueue()
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Expected error when dequeuing empty queue")
	}
}

func TestLinkedQueue_Peek(t *testing.T) {
	q := NewLinkedQueue()
	q.Enqueue("x")
	q.Enqueue("y")
	val, err := q.Peek()
	if err != nil || val != "x" {
		t.Errorf("Expected to peek 'x', got %v (err: %v)", val, err)
	}
	if q.Size() != 2 {
		t.Errorf("Peek should not change size")
	}
}

func TestLinkedQueue_PeekEmpty(t *testing.T) {
	q := NewLinkedQueue()
	_, err := q.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking empty queue")
	}
}

func BenchmarkLinkedQueue_Enqueue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := NewLinkedQueue()
		for i := 0; i < 1000; i++ {
			q.Enqueue(i)
		}
	}
}

func BenchmarkLinkedQueue_Dequeue(b *testing.B) {
	q := NewLinkedQueue()
	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		q.Dequeue()
	}
}

func BenchmarkLinkedQueue_Peek(b *testing.B) {
	q := NewLinkedQueue()
	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		q.Peek()
	}
} 
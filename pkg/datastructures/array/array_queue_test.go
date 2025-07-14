package array

import (
	"testing"
)

func TestNewArrayQueue(t *testing.T) {
	q := NewArrayQueue()
	if q.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", q.Size())
	}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty")
	}
	if len(q.arr) != 10 {
		t.Errorf("Expected initial capacity to be 10, got %d", len(q.arr))
	}
}

func TestArrayQueue_EnqueueDequeue(t *testing.T) {
	q := NewArrayQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", q.Size())
	}

	val, err := q.Dequeue()
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %v", val)
	}

	val, _ = q.Dequeue()
	if val != 2 {
		t.Errorf("Expected 2, got %v", val)
	}

	val, _ = q.Dequeue()
	if val != 3 {
		t.Errorf("Expected 3, got %v", val)
	}

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty after dequeues")
	}
}

func TestArrayQueue_DequeueEmpty(t *testing.T) {
	q := NewArrayQueue()
	_, err := q.Dequeue()
	if err == nil {
		t.Errorf("Expected error when dequeuing empty queue, got nil")
	}
}

func TestArrayQueue_Peek(t *testing.T) {
	q := NewArrayQueue()
	q.Enqueue("x")
	q.Enqueue("y")
	val, err := q.Peek()
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	if val != "x" {
		t.Errorf("Expected 'x', got %v", val)
	}
	if q.Size() != 2 {
		t.Errorf("Peek should not change size")
	}
}

func TestArrayQueue_PeekEmpty(t *testing.T) {
	q := NewArrayQueue()
	_, err := q.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking empty queue, got nil")
	}
}

func TestArrayQueue_ResizeAndShrink(t *testing.T) {
	q := NewArrayQueue()
	// Enqueue enough to trigger resize
	for i := 0; i < 20; i++ {
		q.Enqueue(i)
	}
	if len(q.arr) < 20 {
		t.Errorf("Expected capacity >= 20 after enqueues, got %d", len(q.arr))
	}
	// Dequeue enough to trigger shrink
	for i := 0; i < 15; i++ {
		q.Dequeue()
	}
	if len(q.arr) >= 20 {
		t.Errorf("Expected capacity to shrink, got %d", len(q.arr))
	}
}

func TestArrayQueue_CircularBuffer(t *testing.T) {
	q := NewArrayQueue()
	for i := 0; i < 10; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 5; i++ {
		q.Dequeue()
	}
	for i := 10; i < 15; i++ {
		q.Enqueue(i)
	}
	// Now the buffer should have wrapped around
	for i := 5; i < 15; i++ {
		val, err := q.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if val != i {
			t.Errorf("Expected %d, got %v", i, val)
		}
	}
	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty after wraparound dequeues")
	}
}

func BenchmarkArrayQueue_Enqueue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		q := NewArrayQueue()
		for i := 0; i < 1000; i++ {
			q.Enqueue(i)
		}
	}
}

func BenchmarkArrayQueue_Dequeue(b *testing.B) {
	q := NewArrayQueue()
	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		q.Dequeue()
	}
}

func BenchmarkArrayQueue_Peek(b *testing.B) {
	q := NewArrayQueue()
	for i := 0; i < 1000; i++ {
		q.Enqueue(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		q.Peek()
	}
} 
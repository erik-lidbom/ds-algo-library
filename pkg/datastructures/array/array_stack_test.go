package array

import (
	"testing"
)

func TestNewArrayStack(t *testing.T) {
	stack := NewArrayStack()
	if stack.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", stack.Size())
	}
	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
	if len(stack.arr) != 1 {
		t.Errorf("Expected initial capacity to be 1, got %d", len(stack.arr))
	}
}

func TestArrayStack_PushAndPop(t *testing.T) {
	stack := NewArrayStack()
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", stack.Size())
	}

	val, err := stack.Pop()
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	if val != 30 {
		t.Errorf("Expected 30, got %v", val)
	}

	val, _ = stack.Pop()
	if val != 20 {
		t.Errorf("Expected 20, got %v", val)
	}

	val, _ = stack.Pop()
	if val != 10 {
		t.Errorf("Expected 10, got %v", val)
	}

	if !stack.IsEmpty() {
		t.Errorf("Expected stack to be empty after pops")
	}
}

func TestArrayStack_PopEmpty(t *testing.T) {
	stack := NewArrayStack()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Expected error when popping empty stack, got nil")
	}
}

func TestArrayStack_Peek(t *testing.T) {
	stack := NewArrayStack()
	stack.Push("a")
	stack.Push("b")
	val, err := stack.Peek()
	if err != nil {
		t.Errorf("Did not expect error, got %v", err)
	}
	if val != "b" {
		t.Errorf("Expected 'b', got %v", val)
	}
	if stack.Size() != 2 {
		t.Errorf("Peek should not change size")
	}
}

func TestArrayStack_PeekEmpty(t *testing.T) {
	stack := NewArrayStack()
	_, err := stack.Peek()
	if err == nil {
		t.Errorf("Expected error when peeking empty stack, got nil")
	}
}

func TestArrayStack_ResizeAndShrink(t *testing.T) {
	stack := NewArrayStack()
	// Push enough to trigger resize
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	if len(stack.arr) < 10 {
		t.Errorf("Expected capacity >= 10 after pushes, got %d", len(stack.arr))
	}
	// Pop enough to trigger shrink
	for i := 0; i < 8; i++ {
		stack.Pop()
	}
	if len(stack.arr) >= 10 {
		t.Errorf("Expected capacity to shrink, got %d", len(stack.arr))
	}
}

func BenchmarkArrayStack_Push(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stack := NewArrayStack()
		for i := 0; i < 1000; i++ {
			stack.Push(i)
		}
	}
}

func BenchmarkArrayStack_Pop(b *testing.B) {
	stack := NewArrayStack()
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		stack.Pop()
	}
}

func BenchmarkArrayStack_Peek(b *testing.B) {
	stack := NewArrayStack()
	for i := 0; i < 1000; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		stack.Peek()
	}
} 
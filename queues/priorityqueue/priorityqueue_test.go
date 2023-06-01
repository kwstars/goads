package priorityqueue

import "testing"

var (
	IntMinHeap = func(a, b int) int8 {
		if a < b {
			return 1
		}
		return 0
	}
)

func TestQueue(t *testing.T) {
	queue := New(IntMinHeap)

	if !queue.Empty() {
		t.Errorf("Expected new queue to be empty")
	}

	if size := queue.Size(); size != 0 {
		t.Errorf("Expected new queue size to be 0, got %v", size)
	}

	_, err := queue.Peek()
	if err == nil {
		t.Errorf("Expected Peek() to return an error for empty queue")
	}

	_, err = queue.Dequeue()
	if err == nil {
		t.Errorf("Expected Dequeue() to return an error for empty queue")
	}

	queue.Enqueue(10)

	if queue.Empty() {
		t.Errorf("Expected queue not to be empty after Enqueue")
	}

	if size := queue.Size(); size != 1 {
		t.Errorf("Expected queue size to be 1, got %v", size)
	}

	item, err := queue.Peek()
	if err != nil {
		t.Errorf("Expected Peek() to succeed, got error: %v", err)
	}
	if item != 10 {
		t.Errorf("Expected Peek() to return 10, got %v", item)
	}

	item, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Expected Dequeue() to succeed, got error: %v", err)
	}
	if item != 10 {
		t.Errorf("Expected Dequeue() to return 10, got %v", item)
	}

	if !queue.Empty() {
		t.Errorf("Expected queue to be empty after Dequeue")
	}

	queue.Enqueue(20)
	queue.Enqueue(30)
	queue.Clear()

	if !queue.Empty() {
		t.Errorf("Expected queue to be empty after Clear")
	}
}

package dts

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue("test1")
	q.Enqueue("test2")
	if q.Size() != 2 {
		t.Error("Queue size should be 2")
	}
}

func TestDequeueAll(t *testing.T) {
	q := NewQueue()
	q.Enqueue("test1")
	q.Enqueue("test2")
	q.DequeueAll()
	if q.Size() != 0 {
		t.Error("Queue size should be 0")
	}
}

func TestIsNotEmpty(t *testing.T) {
	q := NewQueue()
	q.Enqueue("test1")
	q.Enqueue("test2")
	if !q.IsNotEmpty() {
		t.Error("Queue should not be empty")
	}
}

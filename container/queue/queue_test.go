package queue

import "testing"

func TestQueueOperations(t *testing.T) {
	t.Run("Formation of a queue", func(t *testing.T) {
		q1 := New(1, 2, 3)
		q2 := New[int]()

		if q1.Len() != 3 {
			t.Errorf("Expected 3, got %d", q1.Len())
		}
		if q1.IsEmpty() {
			t.Error("q1 expected to not be empty")
		}
		if !q2.IsEmpty() {
			t.Error("q2 expected to be empty")
		}
	})

	t.Run("Check Enqueue and Dequeue", func(t *testing.T) {
		q1 := New(1, 2, 3)
		q2 := New[int]()

		q1.Enqueue(4)
		q2.Enqueue(7)
		q2.Enqueue(8)

		val, ok := q1.GetFront()
		if !ok {
			t.Fatal("GetFront() returned ok=false")
		}
		if val != 1 {
			t.Errorf("Got GetFront() = %d, expected 1", val)
		}

		rear, ok := q1.GetRear()
		if !ok {
			t.Fatal("GetRear() returned ok=false")
		}
		if rear != 4 {
			t.Errorf("Got GetRear() = %d, expected 4", rear)
		}
		if q2.IsEmpty() {
			t.Error("q2 expected to not be empty")
		}

		p, ok := q1.Dequeue()

		if !ok {
			t.Fatal("Dequeue() returned ok=false")
		}
		if p != 1 {
			t.Errorf("Got Dequeue() = %d, expected 1", p)
		}

		val, ok = q1.GetFront()
		if !ok {
			t.Fatal("GetFront() returned ok=false")
		}
		if val != 2 {
			t.Errorf("Got GetFront() = %d, expected 2", val)
		}

		q2.Dequeue()
		q2.Dequeue()

		_, ok = q2.Dequeue()
		if ok {
			t.Error("Dequeue() on an empty queue should return ok=false")
		}

		q2.Enqueue(9)

		if q2.Len() != 1 {
			t.Errorf("Got Len() = %d, expected 1", q2.Len())
		}
	})

	t.Run("Other Operations", func(t *testing.T) {
		q := New[int]()

		if !q.IsEmpty() {
			t.Error("new queue expected to be empty")
		}

		q.Enqueue(3)
		q.Enqueue(6)

		if q.Len() != 2 {
			t.Errorf("Got Len() = %d, expected 2", q.Len())
		}

		val, ok := q.GetFront()
		if !ok {
			t.Fatal("GetFront() returned ok=false")
		}
		if val != 3 {
			t.Errorf("Got GetFront() = %d, expected 3", val)
		}

		val, ok = q.GetRear()
		if !ok {
			t.Fatal("GetRear() returned ok=false")
		}
		if val != 6 {
			t.Errorf("Got GetRear() = %d, expected 6", val)
		}

		q.Dequeue()
		q.Dequeue()

		_, ok = q.GetRear()
		if ok {
			t.Error("GetRear() on an empty queue should return ok=false")
		}

		_, ok = q.GetFront()
		if ok {
			t.Error("GetFront() on an empty queue should return ok=false")
		}
	})

	t.Run("Memory Compaction", func(t *testing.T) {
		q := New[int]()

		// Add enough elements to trigger the queue's internal memory compaction.
		for i := 0; i < 300; i++ {
			q.Enqueue(i)
		}

		// Remove elements and verify FIFO order remains correct after compaction.
		for i := 0; i < 200; i++ {
			val, ok := q.Dequeue()

			if !ok {
				t.Fatalf("Dequeue() returned ok=false at %d", i)
			}
			if val != i {
				t.Fatalf("Got Dequeue() = %d, expected %d", val, i)
			}
		}

		if q.Len() != 100 {
			t.Errorf("Got Len() = %d, expected 100", q.Len())
		}

		front, ok := q.GetFront()
		if !ok {
			t.Fatal("Got GetFront() returned ok=false")
		}
		if front != 200 {
			t.Errorf("Got GetFront() = %d, expected 200", front)
		}
	})
}

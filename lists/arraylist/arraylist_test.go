package arraylist

import "testing"

func TestArrayList(t *testing.T) {
	// Define a comparison function for integers
	cmp := func(a, b int) int8 {
		if a == b {
			return 0
		} else if a < b {
			return -1
		} else {
			return 1
		}
	}

	// Create an ArrayList with initial capacity 5
	list := NewArrayList(cmp, WithInitialCapacity[int](5))

	// Test Empty
	if !list.Empty() {
		t.Error("Empty: new list should be empty")
	}

	// Test Append
	for i := 1; i <= 5; i++ {
		list.Append(i)
		if list.Size() != i {
			t.Errorf("Append: expected size = %d, got = %d", i, list.Size())
		}
	}

	// Test Get and IndexOf
	for i := 1; i <= 5; i++ {
		val, err := list.Get(i - 1)
		if err != nil {
			t.Errorf("Get: expected no error, got error: %v", err)
		}
		if val != i {
			t.Errorf("Get: expected element = %d, got = %d", i, val)
		}

		index := list.IndexOf(i)
		if index != i-1 {
			t.Errorf("IndexOf: expected index = %d, got = %d", i-1, index)
		}
	}

	// Test Remove
	err := list.Remove(2)
	if err != nil {
		t.Errorf("Remove: expected no error, got error: %v", err)
	}
	val, err := list.Get(2)
	if err != nil {
		t.Errorf("Get after Remove: expected no error, got error: %v", err)
	}
	if val != 4 {
		t.Errorf("Get after Remove: expected element = %d, got = %d", 4, val)
	}

	// Test PopAtIndex
	val, err = list.PopAtIndex(2)
	if err != nil {
		t.Errorf("PopAtIndex: expected no error, got error: %v", err)
	}
	if val != 4 {
		t.Errorf("PopAtIndex: expected element = %d, got = %d", 4, val)
	}

	// Test Add
	err = list.Add(2, 99)
	if err != nil {
		t.Errorf("Add: expected no error, got error: %v", err)
	}
	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Get after Add: expected no error, got error: %v", err)
	}
	if val != 99 {
		t.Errorf("Get after Add: expected element = %d, got = %d", 99, val)
	}

	// Test Contains
	if !list.Contains(99) {
		t.Errorf("Contains: list should contain %d", 99)
	}

	// Test Set
	err = list.Set(2, 88)
	if err != nil {
		t.Errorf("Set: expected no error, got error: %v", err)
	}
	val, err = list.Get(2)
	if err != nil {
		t.Errorf("Get after Set: expected no error, got error: %v", err)
	}
	if val != 88 {
		t.Errorf("Get after Set: expected element = %d, got = %d", 88, val)
	}

	// Test RemoveIf
	list.RemoveIf(func(v int) bool { return v%2 == 0 }) // remove all even numbers
	for _, v := range list.Copy() {
		if v%2 == 0 {
			t.Errorf("RemoveIf: list should not contain even number %d", v)
		}
	}

	// Test Clear
	list.Clear()
	if !list.Empty() {
		t.Error("Clear: list should be empty after Clear")
	}
}

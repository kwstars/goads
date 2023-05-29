package arraylist

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	// Test with default capacity
	list := NewArrayList[int](nil)
	if list.Size() != 0 {
		t.Error("NewArrayList with default capacity should have size 0")
	}

	// Test with custom capacity
	capacity := 10
	list = NewArrayList[int](nil, WithInitialCapacity[int](capacity))
	if cap(list.elements) != capacity {
		t.Errorf("NewArrayList with custom capacity should have capacity %d", capacity)
	}
}

func TestArrayList_Empty(t *testing.T) {
	t.Run("Empty list - Returns true", func(t *testing.T) {
		// Create an empty ArrayList
		list := NewArrayList[int](nil)

		isEmpty := list.Empty()
		assert.True(t, isEmpty)
	})

	t.Run("Non-empty list - Returns false", func(t *testing.T) {
		// Create an ArrayList with elements
		list := NewArrayList[int](nil)
		list.Append(10)
		list.Append(20)

		isEmpty := list.Empty()
		assert.False(t, isEmpty)
	})
}

func TestArrayList_Full(t *testing.T) {
	t.Run("Returns false", func(t *testing.T) {
		// Create an ArrayList
		list := NewArrayList[int](nil)

		isFull := list.Full()
		assert.False(t, isFull)
	})
}

func TestArrayList_Insert(t *testing.T) {
	t.Run("Normal case - Insert an element at the specified index", func(t *testing.T) {
		// Create an ArrayList with initial elements
		list := NewArrayList[int](nil)
		list.Append(10)
		list.Append(20)
		list.Append(30)

		err := list.Insert(1, 15)
		assert.NoError(t, err)

		expected := []int{10, 15, 20, 30}
		assert.Equal(t, expected, list.elements)
	})

	t.Run("Edge case - Insert an element at the beginning of the list", func(t *testing.T) {
		// Create an ArrayList with initial elements
		list := NewArrayList[int](nil)
		list.Append(10)
		list.Append(20)
		list.Append(30)

		err := list.Insert(0, 5)
		assert.NoError(t, err)

		expected := []int{5, 10, 20, 30}
		assert.Equal(t, expected, list.elements)
	})

	t.Run("Edge case - Insert an element at the end of the list", func(t *testing.T) {
		// Create an ArrayList with initial elements
		list := NewArrayList[int](nil)
		list.Append(10)
		list.Append(20)
		list.Append(30)

		err := list.Insert(list.Size(), 35)
		assert.NoError(t, err)

		expected := []int{10, 20, 30, 35}
		assert.Equal(t, expected, list.elements)
	})

	t.Run("Error case - Insert an element at an invalid index", func(t *testing.T) {
		// Create an empty ArrayList
		list := NewArrayList[int](nil)

		err := list.Insert(1, 10)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrIndexOutOfRange.Error())
	})
}

func TestArrayList_Get(t *testing.T) {
	// Create an ArrayList with initial elements
	list := NewArrayList[int](nil)
	list.Append(10)
	list.Append(20)
	list.Append(30)

	t.Run("Normal case - Retrieve an element at a valid index", func(t *testing.T) {
		element, err := list.Get(1)
		assert.NoError(t, err)
		assert.Equal(t, 20, element)
	})

	t.Run("Edge case - Retrieve an element at the beginning of the list", func(t *testing.T) {
		element, err := list.Get(0)
		assert.NoError(t, err)
		assert.Equal(t, 10, element)
	})

	t.Run("Edge case - Retrieve an element at the end of the list", func(t *testing.T) {
		element, err := list.Get(list.Size() - 1)
		assert.NoError(t, err)
		assert.Equal(t, 30, element)
	})

	t.Run("Error case - Retrieve an element at an invalid index", func(t *testing.T) {
		// Attempt to retrieve an element at an index outside the valid range
		_, err := list.Get(3)
		assert.Error(t, err)
		assert.EqualError(t, err, ErrIndexOutOfRange.Error())
	})
}

func TestArrayList_Append(t *testing.T) {
	t.Run("Normal case - Insert an element to the end of the list", func(t *testing.T) {
		// Create an empty ArrayList
		list := NewArrayList[int](nil)

		list.Append(10)
		expected := []int{10}
		assert.Equal(t, expected, list.elements)

		list.Append(20)
		expected = []int{10, 20}
		assert.Equal(t, expected, list.elements)
	})

	t.Run("Edge case - Insert an element to the end of an existing list", func(t *testing.T) {
		// Create an ArrayList with existing elements
		list := NewArrayList[int](nil)
		list.Append(10)
		list.Append(20)
		list.Append(30)
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, list.elements)
	})
}
func TestArrayList_AppendFront(t *testing.T) {
	t.Run("Normal case - Insert an element to the front of the list", func(t *testing.T) {
		// Create an empty ArrayList
		list := NewArrayList[int](nil)

		list.Prepend(10)
		expected := []int{10}
		assert.Equal(t, expected, list.elements)
	})

	t.Run("Insert more elements to the front", func(t *testing.T) {
		// Create an ArrayList with initial elements
		list := NewArrayList[int](nil)
		list.Append(20)
		list.Append(30)

		list.Prepend(10)
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, list.elements)

		list.Prepend(5)
		expected = []int{5, 10, 20, 30}
		assert.Equal(t, expected, list.elements)
	})

	t.Run("Edge case - Insert an element to the front of an existing list", func(t *testing.T) {
		// Create an ArrayList with existing elements
		list := NewArrayList[int](nil)
		list.Append(20)
		list.Append(30)

		list.Prepend(10)
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, list.elements)
	})
}

func TestArrayList_InsertAll(t *testing.T) {
	// Create an ArrayList with initial elements
	list := NewArrayList[int](nil)
	list.Append(10)
	list.Append(20)
	list.Append(30)

	t.Run("Normal case - Insert multiple elements at the specified index", func(t *testing.T) {
		elements := []int{40, 50}
		err := list.InsertAll(1, elements)
		if err != nil {
			t.Errorf("Error Inserting elements: %v", err)
		}
		expected := []int{10, 40, 50, 20, 30}
		assert.Equal(t, list.elements, expected)
	})

	t.Run("Edge case - Insert multiple elements at the beginning of the list", func(t *testing.T) {
		elements := []int{5, 6, 7}
		err := list.InsertAll(0, elements)
		if err != nil {
			t.Errorf("Error Inserting elements: %v", err)
		}
		expected := []int{5, 6, 7, 10, 40, 50, 20, 30}
		assert.Equal(t, list.elements, expected)
	})

	t.Run("Edge case - Insert multiple elements at the end of the list", func(t *testing.T) {
		elements := []int{60, 70, 80}
		err := list.InsertAll(list.Size(), elements)
		if err != nil {
			t.Errorf("Error Inserting elements: %v", err)
		}
		expected := []int{5, 6, 7, 10, 40, 50, 20, 30, 60, 70, 80}
		assert.Equal(t, list.elements, expected)
	})

	t.Run("Error case - Insert multiple elements at an invalid index", func(t *testing.T) {
		elements := []int{90, 100}
		err := list.InsertAll(100, elements)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}

func TestArrayList_IndexOf(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}

	// Create an ArrayList with multiple elements
	list := NewArrayList[int](cmp)
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Append(40)
	list.Append(50)
	list.Append(30) // Insert a duplicate element

	// Normal case: Find an existing element and return its index
	index := list.IndexOf(20)
	if index != 1 {
		t.Errorf("Expected index 1, but got %d", index)
	}

	// Normal case: Find a duplicate element and return the index of its first occurrence
	index = list.IndexOf(30)
	if index != 2 {
		t.Errorf("Expected index 2, but got %d", index)
	}

	// Edge case: Find an element that doesn't exist in the list
	index = list.IndexOf(100)
	if index != -1 {
		t.Errorf("Expected index -1, but got %d", index)
	}
}

func TestArrayList_LastIndexOf(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestLastIndexOf/element_present", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5, 3}
		index := arrayList.LastIndexOf(3)
		if index != 5 {
			t.Errorf("LastIndexOf() = %v, want %v", index, 5)
		}
	})

	t.Run("TestLastIndexOf/element_absent", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		index := arrayList.LastIndexOf(7)
		if index != -1 {
			t.Errorf("LastIndexOf() = %v, want %v", index, -1)
		}
	})

	t.Run("TestLastIndexOf/multiple_occurrences", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 2, 5, 2}
		index := arrayList.LastIndexOf(2)
		if index != 5 {
			t.Errorf("LastIndexOf() = %v, want %v", index, 5)
		}
	})

	t.Run("TestLastIndexOf/empty_list", func(t *testing.T) {
		arrayList.elements = []int{}
		index := arrayList.LastIndexOf(1)
		if index != -1 {
			t.Errorf("LastIndexOf() = %v, want %v", index, -1)
		}
	})
}

func TestArrayList_Remove(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestRemove/valid_index", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		err := arrayList.Remove(2)
		if err != nil {
			t.Errorf("Remove() error = %v, want nil", err)
		}
		if len(arrayList.elements) != 4 {
			t.Errorf("Remove() remaining elements count = %v, want %v", len(arrayList.elements), 4)
		}
		if arrayList.elements[2] != 4 {
			t.Errorf("Remove() next element = %v, want %v", arrayList.elements[2], 4)
		}
	})

	t.Run("TestRemove/index_out_of_range", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		err := arrayList.Remove(7)
		if err == nil || err.Error() != ErrIndexOutOfRange.Error() {
			t.Errorf("Remove() error = %v, want %v", err, ErrIndexOutOfRange)
		}
	})

	t.Run("TestRemove/negative_index", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		err := arrayList.Remove(-1)
		if err == nil || err.Error() != ErrIndexOutOfRange.Error() {
			t.Errorf("Remove() error = %v, want %v", err, ErrIndexOutOfRange)
		}
	})
}

func TestArrayList_RemoveUnorderedAtIndex(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestRemoveUnorderedAtIndex/valid_index", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		err := arrayList.RemoveUnorderedAtIndex(2)
		if err != nil {
			t.Errorf("RemoveUnorderedAtIndex() error = %v, want nil", err)
		}
		if len(arrayList.elements) != 4 {
			t.Errorf("RemoveUnorderedAtIndex() remaining elements count = %v, want %v", len(arrayList.elements), 4)
		}
		if arrayList.elements[2] != 5 {
			t.Errorf("RemoveUnorderedAtIndex() moved element = %v, want %v", arrayList.elements[2], 5)
		}
	})

	t.Run("TestRemoveUnorderedAtIndex/index_out_of_range", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		err := arrayList.RemoveUnorderedAtIndex(7)
		if err == nil || err.Error() != ErrIndexOutOfRange.Error() {
			t.Errorf("RemoveUnorderedAtIndex() error = %v, want %v", err, ErrIndexOutOfRange)
		}
	})

	t.Run("TestRemoveUnorderedAtIndex/negative_index", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		err := arrayList.RemoveUnorderedAtIndex(-1)
		if err == nil || err.Error() != ErrIndexOutOfRange.Error() {
			t.Errorf("RemoveUnorderedAtIndex() error = %v, want %v", err, ErrIndexOutOfRange)
		}
	})
}

func TestArrayList_PopAtIndex(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestPopAtIndex/valid_index", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		popElement, err := arrayList.PopAtIndex(2)
		if err != nil {
			t.Errorf("PopAtIndex() error = %v, want nil", err)
		}
		if popElement != 3 {
			t.Errorf("PopAtIndex() = %v, want %v", popElement, 3)
		}
		if len(arrayList.elements) != 4 {
			t.Errorf("PopAtIndex() remaining elements count = %v, want %v", len(arrayList.elements), 4)
		}
	})

	t.Run("TestPopAtIndex/index_out_of_range", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		_, err := arrayList.PopAtIndex(7)
		if err == nil || err.Error() != ErrIndexOutOfRange.Error() {
			t.Errorf("PopAtIndex() error = %v, want %v", err, ErrIndexOutOfRange)
		}
	})

	t.Run("TestPopAtIndex/negative_index", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		_, err := arrayList.PopAtIndex(-1)
		if err == nil || err.Error() != ErrIndexOutOfRange.Error() {
			t.Errorf("PopAtIndex() error = %v, want %v", err, ErrIndexOutOfRange)
		}
	})
}

func TestArrayList_Pop(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestPop/valid_pop", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		popElement, err := arrayList.Pop()
		if err != nil {
			t.Errorf("Pop() error = %v, want nil", err)
		}
		if popElement != 5 {
			t.Errorf("Pop() = %v, want %v", popElement, 5)
		}
		if len(arrayList.elements) != 4 {
			t.Errorf("Pop() remaining elements count = %v, want %v", len(arrayList.elements), 4)
		}
	})

	t.Run("TestPop/empty_list", func(t *testing.T) {
		arrayList.elements = []int{}
		_, err := arrayList.Pop()
		if err == nil || err.Error() != ErrPopFromEmptyList.Error() {
			t.Errorf("Pop() error = %v, want %v", err, ErrPopFromEmptyList)
		}
	})
}

func TestArrayList_PopFront(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestPopFront/valid_pop", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		popElement, err := arrayList.PopFront()
		if err != nil {
			t.Errorf("PopFront() error = %v, want nil", err)
		}
		if popElement != 1 {
			t.Errorf("PopFront() = %v, want %v", popElement, 1)
		}
		if arrayList.elements[0] != 2 {
			t.Errorf("PopFront() remaining first element = %v, want %v", arrayList.elements[0], 2)
		}
	})

	t.Run("TestPopFront/empty_list", func(t *testing.T) {
		arrayList.elements = []int{}
		_, err := arrayList.PopFront()
		if err == nil || err.Error() != ErrPopFromEmptyList.Error() {
			t.Errorf("PopFront() error = %v, want %v", err, ErrPopFromEmptyList)
		}
	})
}

func TestArrayList_Set(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestSet/valid_index", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		err := arrayList.Set(2, 6)
		if err != nil {
			t.Errorf("Set() error = %v, want nil", err)
		}
		if arrayList.elements[2] != 6 {
			t.Errorf("Set() = %v, want %v", arrayList.elements[2], 6)
		}
	})

	t.Run("TestSet/invalid_index_negative", func(t *testing.T) {
		err := arrayList.Set(-1, 6)
		if err == nil || err.Error() != ErrIndexOutOfRange.Error() {
			t.Errorf("Set() error = %v, want %v", err, ErrIndexOutOfRange)
		}
	})

	t.Run("TestSet/invalid_index_out_of_bounds", func(t *testing.T) {
		err := arrayList.Set(len(arrayList.elements), 6)
		if err == nil || err.Error() != ErrIndexOutOfRange.Error() {
			t.Errorf("Set() error = %v, want %v", err, ErrIndexOutOfRange)
		}
	})
}

func TestArrayList_Clear(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestClear/empty_list", func(t *testing.T) {
		arrayList.Clear()
		if len(arrayList.elements) != 0 {
			t.Errorf("Clear() = %v, want %v", len(arrayList.elements), 0)
		}
	})

	t.Run("TestClear/non_empty_list", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		arrayList.Clear()
		if len(arrayList.elements) != 0 {
			t.Errorf("Clear() = %v, want %v", len(arrayList.elements), 0)
		}
	})
}

func TestArrayList_Contains(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)
	arrayList.elements = []int{1, 2, 3, 4, 5}

	t.Run("TestContains/single_element_present", func(t *testing.T) {
		result := arrayList.Contains(3)
		if !result {
			t.Errorf("Contains() = %v, want %v", result, true)
		}
	})

	t.Run("TestContains/multiple_elements_present", func(t *testing.T) {
		result := arrayList.Contains(1, 2, 3)
		if !result {
			t.Errorf("Contains() = %v, want %v", result, true)
		}
	})

	t.Run("TestContains/single_element_absent", func(t *testing.T) {
		result := arrayList.Contains(6)
		if result {
			t.Errorf("Contains() = %v, want %v", result, false)
		}
	})

	t.Run("TestContains/multiple_elements_absent", func(t *testing.T) {
		result := arrayList.Contains(6, 7, 8)
		if result {
			t.Errorf("Contains() = %v, want %v", result, false)
		}
	})

	t.Run("TestContains/some_elements_absent", func(t *testing.T) {
		result := arrayList.Contains(4, 5, 6)
		if result {
			t.Errorf("Contains() = %v, want %v", result, false)
		}
	})
}

func TestArrayList_RemoveRange(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)
	arrayList.elements = []int{1, 2, 3, 4, 5}

	t.Run("TestRemoveRange/valid_range", func(t *testing.T) {
		err := arrayList.RemoveRange(1, 4)
		if err != nil {
			t.Fatalf("RemoveRange() error = %v, wantErr %v", err, nil)
		}
		if !reflect.DeepEqual(arrayList.elements, []int{1, 5}) {
			t.Errorf("RemoveRange() = %v, want %v", arrayList.elements, []int{1, 5})
		}
	})

	t.Run("TestRemoveRange/invalid_range", func(t *testing.T) {
		err := arrayList.RemoveRange(4, 1)
		if err == nil {
			t.Fatalf("RemoveRange() error = %v, wantErr %v", err, ErrFormIndexMustBeLessThanToIndex)
		}
	})

	t.Run("TestRemoveRange/fromIndex_out_of_range", func(t *testing.T) {
		err := arrayList.RemoveRange(-1, 4)
		if err == nil {
			t.Fatalf("RemoveRange() error = %v, wantErr %v", err, ErrIndexOutOfRange)
		}
	})

	t.Run("TestRemoveRange/toIndex_out_of_range", func(t *testing.T) {
		err := arrayList.RemoveRange(1, 6)
		if err == nil {
			t.Fatalf("RemoveRange() error = %v, wantErr %v", err, ErrIndexOutOfRange)
		}
	})
}

func TestArrayList_SubList(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)
	arrayList.elements = []int{1, 2, 3, 4, 5}

	t.Run("TestSubList/valid_range", func(t *testing.T) {
		subList, err := arrayList.SubList(1, 4)
		if err != nil {
			t.Fatalf("SubList() error = %v, wantErr %v", err, nil)
		}
		if !reflect.DeepEqual(subList, []int{2, 3, 4}) {
			t.Errorf("SubList() = %v, want %v", subList, []int{2, 3, 4})
		}
	})

	t.Run("TestSubList/invalid_range", func(t *testing.T) {
		_, err := arrayList.SubList(4, 1)
		if err == nil {
			t.Fatalf("SubList() error = %v, wantErr %v", err, ErrFormIndexMustBeLessThanToIndex)
		}
	})

	t.Run("TestSubList/fromIndex_out_of_range", func(t *testing.T) {
		_, err := arrayList.SubList(-1, 4)
		if err == nil {
			t.Fatalf("SubList() error = %v, wantErr %v", err, ErrIndexOutOfRange)
		}
	})

	t.Run("TestSubList/toIndex_out_of_range", func(t *testing.T) {
		_, err := arrayList.SubList(1, 6)
		if err == nil {
			t.Fatalf("SubList() error = %v, wantErr %v", err, ErrIndexOutOfRange)
		}
	})
}

func TestArrayList_Reverse(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestReverse/odd_number_of_elements", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		arrayList.Reverse()
		if !reflect.DeepEqual(arrayList.elements, []int{5, 4, 3, 2, 1}) {
			t.Errorf("Reverse() = %v, want %v", arrayList.elements, []int{5, 4, 3, 2, 1})
		}
	})

	t.Run("TestReverse/even_number_of_elements", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4}
		arrayList.Reverse()
		if !reflect.DeepEqual(arrayList.elements, []int{4, 3, 2, 1}) {
			t.Errorf("Reverse() = %v, want %v", arrayList.elements, []int{4, 3, 2, 1})
		}
	})

	t.Run("TestReverse/single_element", func(t *testing.T) {
		arrayList.elements = []int{1}
		arrayList.Reverse()
		if !reflect.DeepEqual(arrayList.elements, []int{1}) {
			t.Errorf("Reverse() = %v, want %v", arrayList.elements, []int{1})
		}
	})

	t.Run("TestReverse/empty_arrayList", func(t *testing.T) {
		arrayList.elements = []int{}
		arrayList.Reverse()
		if !reflect.DeepEqual(arrayList.elements, []int{}) {
			t.Errorf("Reverse() = %v, want %v", arrayList.elements, []int{})
		}
	})
}

func TestArrayList_RemoveIf(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestRemoveIf/all_elements_meet_predicate", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		predicate := func(i int) bool { return i < 6 }
		result := arrayList.RemoveIf(predicate)
		if result != true {
			t.Errorf("RemoveIf() = %v, want %v", result, true)
		}
		if !reflect.DeepEqual(arrayList.elements, []int{}) {
			t.Errorf("RemoveIf() = %v, want %v", arrayList.elements, []int{})
		}
	})

	t.Run("TestRemoveIf/no_elements_meet_predicate", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		predicate := func(i int) bool { return i > 5 }
		result := arrayList.RemoveIf(predicate)
		if result != false {
			t.Errorf("RemoveIf() = %v, want %v", result, false)
		}
		if !reflect.DeepEqual(arrayList.elements, []int{1, 2, 3, 4, 5}) {
			t.Errorf("RemoveIf() = %v, want %v", arrayList.elements, []int{1, 2, 3, 4, 5})
		}
	})

	t.Run("TestRemoveIf/some_elements_meet_predicate", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		predicate := func(i int) bool { return i%2 == 0 }
		result := arrayList.RemoveIf(predicate)
		if result != true {
			t.Errorf("RemoveIf() = %v, want %v", result, true)
		}
		if !reflect.DeepEqual(arrayList.elements, []int{1, 3, 5}) {
			t.Errorf("RemoveIf() = %v, want %v", arrayList.elements, []int{1, 3, 5})
		}
	})

	t.Run("TestRemoveIf/empty_arrayList", func(t *testing.T) {
		arrayList.elements = []int{}
		predicate := func(i int) bool { return i%2 == 0 }
		result := arrayList.RemoveIf(predicate)
		if result != false {
			t.Errorf("RemoveIf() = %v, want %v", result, false)
		}
		if !reflect.DeepEqual(arrayList.elements, []int{}) {
			t.Errorf("RemoveIf() = %v, want %v", arrayList.elements, []int{})
		}
	})
}

func TestArrayList_Sort(t *testing.T) {
	list := NewArrayList[int](func(a, b int) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	})

	// Insert unsorted elements to the list
	list.Append(5)
	list.Append(2)
	list.Append(8)
	list.Append(1)
	list.Append(3)

	// Test Sort
	list.Sort()

	// Verify the elements are sorted in ascending order
	expected := []int{1, 2, 3, 5, 8}
	for i := 0; i < list.Size(); i++ {
		element, _ := list.Get(i)
		if element != expected[i] {
			t.Errorf("Sort() failed: got %d, expected %d", element, expected[i])
		}
	}
}

func TestArrayList_Copy(t *testing.T) {
	list := NewArrayList[int](nil)

	// Insert elements to the list
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test Copy
	copiedList := list.Copy()

	// Verify the length of the copied list is the same
	if len(copiedList) != list.Size() {
		t.Error("Copy() failed to copy all elements")
	}

	// Verify the elements in the copied list are the same
	for i := 0; i < list.Size(); i++ {
		element1, _ := list.Get(i)
		element2 := copiedList[i]
		if element1 != element2 {
			t.Errorf("Copy() failed to copy element: got %d, expected %d", element2, element1)
		}
	}

	// Modify the original list and verify the copied list remains unchanged
	_ = list.Set(0, 100)
	element := copiedList[0]
	if element == 100 {
		t.Error("Copy() failed to create a deep copy")
	}
}

func TestArrayList_BatchSplit(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestBatchSplit/batchSize_is_negative", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.BatchSplit(-1)
		if !reflect.DeepEqual(result, [][]int{}) {
			t.Errorf("BatchSplit() = %v, want %v", result, [][]int{})
		}
	})

	t.Run("TestBatchSplit/batchSize_is_zero", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.BatchSplit(0)
		if !reflect.DeepEqual(result, [][]int{}) {
			t.Errorf("BatchSplit() = %v, want %v", result, [][]int{})
		}
	})

	t.Run("TestBatchSplit/batchSize_is_one", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.BatchSplit(1)
		if !reflect.DeepEqual(result, [][]int{{1}, {2}, {3}, {4}, {5}}) {
			t.Errorf("BatchSplit() = %v, want %v", result, [][]int{{1}, {2}, {3}, {4}, {5}})
		}
	})

	t.Run("TestBatchSplit/batchSize_greater_than_length_of_arrayList", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.BatchSplit(10)
		if !reflect.DeepEqual(result, [][]int{{1, 2, 3, 4, 5}}) {
			t.Errorf("BatchSplit() = %v, want %v", result, [][]int{{1, 2, 3, 4, 5}})
		}
	})

	t.Run("TestBatchSplit/batchSize_less_than_length_of_arrayList", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := arrayList.BatchSplit(3)
		if !reflect.DeepEqual(result, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}) {
			t.Errorf("BatchSplit() = %v, want %v", result, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
		}
	})

	t.Run("TestBatchSplit/batchSize_cannot_divide_arrayList_evenly", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5, 6, 7}
		result := arrayList.BatchSplit(3)
		if !reflect.DeepEqual(result, [][]int{{1, 2, 3}, {4, 5, 6}, {7}}) {
			t.Errorf("BatchSplit() = %v, want %v", result, [][]int{{1, 2, 3}, {4, 5, 6}, {7}})
		}
	})
}

func TestArrayList_SlidingWindows(t *testing.T) {
	// Create a comparator function for integers
	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}
	arrayList := NewArrayList(cmp)

	t.Run("TestSlidingWindows/size_is_negative", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.SlidingWindows(-1)
		if !reflect.DeepEqual(result, [][]int{}) {
			t.Errorf("SlidingWindows() = %v, want %v", result, [][]int{})
		}
	})

	t.Run("TestSlidingWindows/size_is_zero", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.SlidingWindows(0)
		if !reflect.DeepEqual(result, [][]int{}) {
			t.Errorf("SlidingWindows() = %v, want %v", result, [][]int{})
		}
	})

	t.Run("TestSlidingWindows/size_is_one", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.SlidingWindows(1)
		if !reflect.DeepEqual(result, [][]int{{1}, {2}, {3}, {4}, {5}}) {
			t.Errorf("SlidingWindows() = %v, want %v", result, [][]int{{1}, {2}, {3}, {4}, {5}})
		}
	})

	t.Run("TestSlidingWindows/size_greater_than_length_of_arrayList", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.SlidingWindows(10)
		if !reflect.DeepEqual(result, [][]int{{1, 2, 3, 4, 5}}) {
			t.Errorf("SlidingWindows() = %v, want %v", result, [][]int{{1, 2, 3, 4, 5}})
		}
	})

	t.Run("TestSlidingWindows/size_less_than_length_of_arrayList", func(t *testing.T) {
		arrayList.elements = []int{1, 2, 3, 4, 5}
		result := arrayList.SlidingWindows(3)
		if !reflect.DeepEqual(result, [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}) {
			t.Errorf("SlidingWindows() = %v, want %v", result, [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}})
		}
	})
}

func TestArrayList_BringElementToFront(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		element  int
		expected []int
	}{
		{
			name:     "empty list",
			input:    []int{},
			element:  1,
			expected: []int{1},
		},
		{
			name:     "element at front",
			input:    []int{1, 2, 3, 4, 5},
			element:  1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "element at end",
			input:    []int{2, 3, 4, 5, 1},
			element:  1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "element in middle",
			input:    []int{2, 1, 3, 4, 5},
			element:  1,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "element not in list",
			input:    []int{2, 3, 4, 5},
			element:  1,
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Initialize ArrayList with the test case input
			list := NewArrayList(cmp, WithInitialCapacity[int](len(tc.input)))
			list.elements = append(list.elements, tc.input...)

			list.BringElementToFront(tc.element)

			if !reflect.DeepEqual(list.elements, tc.expected) {
				t.Errorf("BringElementToFront() = %v, want %v", list.elements, tc.expected)
			}
		})
	}
}

func TestArrayList_RemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "all duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "mixed duplicates and non-duplicates",
			input:    []int{1, 2, 2, 3, 4, 4, 4, 5, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	cmp := func(a, b int) int8 {
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Initialize ArrayList with the test case input
			list := NewArrayList(cmp, WithInitialCapacity[int](len(tc.input)))
			list.elements = append(list.elements, tc.input...)

			list.RemoveDuplicates()

			if !reflect.DeepEqual(list.elements, tc.expected) {
				t.Errorf("RemoveDuplicates() = %v, want %v", list.elements, tc.expected)
			}
		})
	}
}

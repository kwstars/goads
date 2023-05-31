package doublylinkedlist

import (
	"errors"
	"github.com/kwstars/goads/pkg/common"
	"reflect"
	"testing"
)

func TestList_Append(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		wantSize int
		wantTail int
	}{
		{"append one element", []int{1}, 1, 1},
		{"append two elements", []int{1, 2}, 2, 2},
		{"append three elements", []int{1, 2, 3}, 3, 3},
		{"append multiple same elements", []int{1, 1, 1}, 3, 1},
		{"append zero", []int{0}, 1, 0},
		{"append negative numbers", []int{-1, -2, -3}, 3, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			if list.size != tt.wantSize {
				t.Errorf("Expected size of %d, but got %d", tt.wantSize, list.size)
			}

			if list.tail.prev.value != tt.wantTail {
				t.Errorf("Expected tail value of %v, but got %v", tt.wantTail, list.tail.prev.value)
			}
		})
	}
}

func TestList_Prepend(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		wantSize int
		wantHead int
	}{
		{"prepend one element", []int{1}, 1, 1},
		{"prepend two elements", []int{1, 2}, 2, 2},
		{"prepend three elements", []int{1, 2, 3}, 3, 3},
		{"prepend multiple same elements", []int{1, 1, 1}, 3, 1},
		{"prepend zero", []int{0}, 1, 0},
		{"prepend negative numbers", []int{-1, -2, -3}, 3, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Prepend(v)
			}

			if list.size != tt.wantSize {
				t.Errorf("Expected size of %d, but got %d", tt.wantSize, list.size)
			}

			if list.head.next.value != tt.wantHead {
				t.Errorf("Expected head value of %v, but got %v", tt.wantHead, list.head.next.value)
			}
		})
	}
}

func TestList_Get(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		index     int
		wantValue int
		wantError error
	}{
		{"get first element", []int{1, 2, 3}, 0, 1, nil},
		{"get middle element", []int{1, 2, 3}, 1, 2, nil},
		{"get last element", []int{1, 2, 3}, 2, 3, nil},
		{"get out of range element", []int{1, 2, 3}, 3, 0, ErrIndexOutOfRange},
		{"get element from empty list", []int{}, 0, 0, ErrIndexOutOfRange},
		{"get negative index", []int{1, 2, 3}, -1, 0, ErrIndexOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			gotValue, err := list.Get(tt.index)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Expected error = %v, but got error = %v", tt.wantError, err != nil)
			}

			if gotValue != tt.wantValue {
				t.Errorf("Expected value of %v, but got %v", tt.wantValue, gotValue)
			}
		})
	}
}

func TestList_Insert(t *testing.T) {
	tests := []struct {
		name        string
		values      []int
		index       int
		insertValue int
		wantSize    int
		wantError   error
		wantList    []int
	}{
		{"insert at beginning", []int{1, 2, 3}, 0, 0, 4, nil, []int{0, 1, 2, 3}},
		{"insert in middle", []int{1, 2, 3}, 1, 0, 4, nil, []int{1, 0, 2, 3}},
		{"insert at end", []int{1, 2, 3}, 3, 0, 4, nil, []int{1, 2, 3, 0}},
		{"insert out of range", []int{1, 2, 3}, 4, 0, 3, ErrIndexOutOfRange, []int{1, 2, 3}},
		{"insert negative index", []int{1, 2, 3}, -1, 0, 3, ErrIndexOutOfRange, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.Insert(tt.index, tt.insertValue)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Expected error = %v, but got error = %v", tt.wantError, err)
			}

			if list.size != tt.wantSize {
				t.Errorf("Expected size of %d, but got %d", tt.wantSize, list.size)
			}

			for i := 0; i < list.size; i++ {
				gotValue, _ := list.Get(i)
				if gotValue != tt.wantList[i] {
					t.Errorf("At index %d: expected value of %v, but got %v", i, tt.wantList[i], gotValue)
				}
			}
		})
	}
}

func TestList_InsertAll(t *testing.T) {
	tests := []struct {
		name         string
		values       []int
		index        int
		insertValues []int
		wantSize     int
		wantError    error
		wantList     []int
	}{
		{"insert at beginning", []int{1, 2, 3}, 0, []int{0, -1}, 5, nil, []int{0, -1, 1, 2, 3}},
		{"insert in middle", []int{1, 2, 3}, 1, []int{0, -1}, 5, nil, []int{1, 0, -1, 2, 3}},
		{"insert at end", []int{1, 2, 3}, 3, []int{0, -1}, 5, nil, []int{1, 2, 3, 0, -1}},
		{"insert out of range", []int{1, 2, 3}, 4, []int{0, -1}, 3, ErrIndexOutOfRange, []int{1, 2, 3}},
		{"insert negative index", []int{1, 2, 3}, -1, []int{0, -1}, 3, ErrIndexOutOfRange, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.InsertAll(tt.index, tt.insertValues)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Expected error = %v, but got error = %v", tt.wantError, err)
			}

			if list.size != tt.wantSize {
				t.Errorf("Expected size of %d, but got %d", tt.wantSize, list.size)
			}

			for i := 0; i < list.size; i++ {
				gotValue, _ := list.Get(i)
				if gotValue != tt.wantList[i] {
					t.Errorf("At index %d: expected value of %v, but got %v", i, tt.wantList[i], gotValue)
				}
			}
		})
	}
}

func TestList_Clear(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		wantSize  int
		wantError error
	}{
		{"clear empty list", []int{}, 0, nil},
		{"clear list with one element", []int{1}, 0, nil},
		{"clear list with multiple elements", []int{1, 2, 3}, 0, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			list.Clear()

			if list.size != tt.wantSize {
				t.Errorf("Expected size of %d, but got %d", tt.wantSize, list.size)
			}

			// Attempt to get an element from the list should return ErrIndexOutOfRange error
			_, err := list.Get(0)
			if !errors.Is(err, ErrIndexOutOfRange) {
				t.Errorf("Expected error = %v, but got error = %v", ErrIndexOutOfRange, err)
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		index     int
		wantSize  int
		wantError error
		wantList  []int
	}{
		{"remove from beginning", []int{1, 2, 3}, 0, 2, nil, []int{2, 3}},
		{"remove from middle", []int{1, 2, 3}, 1, 2, nil, []int{1, 3}},
		{"remove from end", []int{1, 2, 3}, 2, 2, nil, []int{1, 2}},
		{"remove out of range", []int{1, 2, 3}, 3, 3, ErrIndexOutOfRange, []int{1, 2, 3}},
		{"remove negative index", []int{1, 2, 3}, -1, 3, ErrIndexOutOfRange, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.Remove(tt.index)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Expected error = %v, but got error = %v", tt.wantError, err)
			}

			if list.size != tt.wantSize {
				t.Errorf("Expected size of %d, but got %d", tt.wantSize, list.size)
			}

			for i := 0; i < list.size; i++ {
				gotValue, _ := list.Get(i)
				if gotValue != tt.wantList[i] {
					t.Errorf("At index %d: expected value of %v, but got %v", i, tt.wantList[i], gotValue)
				}
			}
		})
	}
}

func TestList_RemoveRange(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		fromIndex int
		toIndex   int
		wantSize  int
		wantError error
		wantList  []int
	}{
		{"remove middle range", []int{1, 2, 3, 4, 5}, 1, 3, 3, nil, []int{1, 4, 5}},
		{"remove from beginning", []int{1, 2, 3, 4, 5}, 0, 2, 3, nil, []int{3, 4, 5}},
		{"remove to end", []int{1, 2, 3, 4, 5}, 2, 4, 3, nil, []int{1, 2, 5}},
		{"remove full range", []int{1, 2, 3, 4, 5}, 0, 5, 0, nil, []int{}},
		{"fromIndex out of range", []int{1, 2, 3, 4, 5}, -1, 3, 5, ErrIndexOutOfRange, []int{1, 2, 3, 4, 5}},
		{"toIndex out of range", []int{1, 2, 3, 4, 5}, 1, 6, 5, ErrIndexOutOfRange, []int{1, 2, 3, 4, 5}},
		{"fromIndex greater than toIndex", []int{1, 2, 3, 4, 5}, 3, 1, 5, ErrFormIndexMustBeLessThanToIndex, []int{1, 2, 3, 4, 5}},
		{"fromIndex equals toIndex", []int{1, 2, 3, 4, 5}, 2, 2, 5, nil, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.RemoveRange(tt.fromIndex, tt.toIndex)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Expected error = %v, but got error = %v", tt.wantError, err)
			}

			if list.size != tt.wantSize {
				t.Errorf("Expected size of %d, but got %d", tt.wantSize, list.size)
			}

			for i := 0; i < list.size; i++ {
				gotValue, _ := list.Get(i)
				if gotValue != tt.wantList[i] {
					t.Errorf("At index %d: expected value of %v, but got %v", i, tt.wantList[i], gotValue)
				}
			}
		})
	}
}

func TestList_Set(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		setIndex  int
		setValue  int
		wantError error
		wantList  []int
	}{
		{"set first element", []int{1, 2, 3, 4, 5}, 0, 9, nil, []int{9, 2, 3, 4, 5}},
		{"set middle element", []int{1, 2, 3, 4, 5}, 2, 9, nil, []int{1, 2, 9, 4, 5}},
		{"set last element", []int{1, 2, 3, 4, 5}, 4, 9, nil, []int{1, 2, 3, 4, 9}},
		{"set with negative index", []int{1, 2, 3, 4, 5}, -1, 9, ErrIndexOutOfRange, []int{1, 2, 3, 4, 5}},
		{"set with out of range index", []int{1, 2, 3, 4, 5}, 5, 9, ErrIndexOutOfRange, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.Set(tt.setIndex, tt.setValue)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Expected error = %v, but got error = %v", tt.wantError, err)
			}

			for i := 0; i < list.size; i++ {
				gotValue, _ := list.Get(i)
				if gotValue != tt.wantList[i] {
					t.Errorf("At index %d: expected value of %v, but got %v", i, tt.wantList[i], gotValue)
				}
			}
		})
	}
}

func TestList_IndexOf(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		target    int
		wantIndex int
	}{
		{"find first element", []int{1, 2, 3, 4, 5}, 1, 0},
		{"find middle element", []int{1, 2, 3, 4, 5}, 3, 2},
		{"find last element", []int{1, 2, 3, 4, 5}, 5, 4},
		{"find non-existing element", []int{1, 2, 3, 4, 5}, 6, -1},
		{"find element in empty list", []int{}, 1, -1},
		{"find repeated element", []int{1, 2, 2, 3, 3, 3}, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			gotIndex := list.IndexOf(tt.target)
			if gotIndex != tt.wantIndex {
				t.Errorf("Expected index of %d, but got %d", tt.wantIndex, gotIndex)
			}
		})
	}
}

func TestList_LastIndexOf(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		target    int
		wantIndex int
	}{
		{"find last element", []int{1, 2, 3, 4, 5}, 5, 4},
		{"find middle element", []int{1, 2, 3, 4, 5}, 3, 2},
		{"find first element", []int{1, 2, 3, 4, 5}, 1, 0},
		{"find non-existing element", []int{1, 2, 3, 4, 5}, 6, -1},
		{"find element in empty list", []int{}, 1, -1},
		{"find repeated element", []int{1, 2, 2, 3, 3, 3}, 3, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			gotIndex := list.LastIndexOf(tt.target)
			if gotIndex != tt.wantIndex {
				t.Errorf("Expected last index of %d, but got %d", tt.wantIndex, gotIndex)
			}
		})
	}
}

func TestList_SubList(t *testing.T) {
	tests := []struct {
		name      string
		values    []int
		fromIndex int
		toIndex   int
		want      []int
		wantError error
	}{
		{"valid sublist from start", []int{1, 2, 3, 4, 5}, 0, 3, []int{1, 2, 3}, nil},
		{"valid sublist from middle", []int{1, 2, 3, 4, 5}, 2, 4, []int{3, 4}, nil},
		{"valid sublist till end", []int{1, 2, 3, 4, 5}, 2, 5, []int{3, 4, 5}, nil},
		{"invalid fromIndex", []int{1, 2, 3, 4, 5}, -1, 3, nil, ErrIndexOutOfRange},
		{"invalid toIndex", []int{1, 2, 3, 4, 5}, 2, 6, nil, ErrIndexOutOfRange},
		{"fromIndex > toIndex", []int{1, 2, 3, 4, 5}, 4, 2, nil, ErrFormIndexMustBeLessThanToIndex},
		{"fromIndex == toIndex", []int{1, 2, 3, 4, 5}, 2, 2, []int{}, nil},
		{"empty list", []int{}, 0, 0, []int{}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			got, err := list.SubList(tt.fromIndex, tt.toIndex)
			if !errors.Is(err, tt.wantError) {
				t.Errorf("Expected error %v, but got %v", tt.wantError, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected sublist %v, but got %v", tt.want, got)
			}
		})
	}
}

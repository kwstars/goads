package arraylist

import (
	"github.com/kwstars/goads/pkg/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_New(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		cmp      func(a, b int) int8
		wantLen  int
		wantCap  int
	}{
		{
			name:    "without initial capacity",
			cmp:     common.IntsCompare,
			wantLen: 0,
			wantCap: 0,
		},
		{
			name:     "with initial capacity of 10",
			capacity: 10,
			cmp:      common.IntsCompare,
			wantLen:  0,
			wantCap:  10,
		},
		{
			name:     "with initial capacity of 100",
			capacity: 100,
			cmp:      common.IntsCompare,
			wantLen:  0,
			wantCap:  100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opts []Option[int]
			if tt.capacity != 0 {
				opts = append(opts, WithInitialCapacity[int](tt.capacity))
			}

			l := New[int](tt.cmp, opts...)
			assert.Equal(t, tt.wantLen, len(l.elements), "unexpected list length")
			assert.Equal(t, tt.wantCap, cap(l.elements), "unexpected list capacity")
		})
	}
}
func TestList_Empty(t *testing.T) {
	t.Run("Empty list - Returns true", func(t *testing.T) {
		// Create an empty ArrayList
		list := New[int](nil)

		isEmpty := list.Empty()
		assert.True(t, isEmpty)
	})

	t.Run("Non-empty list - Returns false", func(t *testing.T) {
		// Create an ArrayList with elements
		list := New[int](nil)
		list.Append(10)
		list.Append(20)

		isEmpty := list.Empty()
		assert.False(t, isEmpty)
	})
}

func TestArrayList_Full(t *testing.T) {
	t.Run("Returns false", func(t *testing.T) {
		// Create an ArrayList
		list := New[int](nil)

		isFull := list.Full()
		assert.False(t, isFull)
	})
}

func TestList_Append(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		append   int
		want     []int
	}{
		{
			name:     "append to an empty list",
			elements: []int{},
			append:   1,
			want:     []int{1},
		},
		{
			name:     "append to a non-empty list",
			elements: []int{1, 2, 3},
			append:   4,
			want:     []int{1, 2, 3, 4},
		},
		{
			name:     "append zero to a list",
			elements: []int{1, 2, 3},
			append:   0,
			want:     []int{1, 2, 3, 0},
		},
		{
			name:     "append negative number to a list",
			elements: []int{1, 2, 3},
			append:   -1,
			want:     []int{1, 2, 3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)+1))
			for _, element := range tt.elements {
				l.Append(element)
			}

			l.Append(tt.append)
			assert.Equal(t, tt.want, l.elements)
		})
	}
}

func TestList_Prepend(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		prepend  int
		want     []int
	}{
		{
			name:     "prepend to an empty list",
			elements: []int{},
			prepend:  1,
			want:     []int{1},
		},
		{
			name:     "prepend to a non-empty list",
			elements: []int{1, 2, 3},
			prepend:  4,
			want:     []int{4, 1, 2, 3},
		},
		{
			name:     "prepend zero to a list",
			elements: []int{1, 2, 3},
			prepend:  0,
			want:     []int{0, 1, 2, 3},
		},
		{
			name:     "prepend negative number to a list",
			elements: []int{1, 2, 3},
			prepend:  -1,
			want:     []int{-1, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)+1))
			for _, element := range tt.elements {
				l.Append(element)
			}

			l.Prepend(tt.prepend)
			assert.Equal(t, tt.want, l.elements)
		})
	}
}

func TestList_Insert(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		insertIdx int
		insertVal int
		want      []int
		wantErr   error
	}{
		{
			name:      "insert into an empty list",
			elements:  []int{},
			insertIdx: 0,
			insertVal: 1,
			want:      []int{1},
			wantErr:   nil,
		},
		{
			name:      "insert at the beginning of a list",
			elements:  []int{1, 2, 3},
			insertIdx: 0,
			insertVal: 4,
			want:      []int{4, 1, 2, 3},
			wantErr:   nil,
		},
		{
			name:      "insert in the middle of a list",
			elements:  []int{1, 2, 3},
			insertIdx: 1,
			insertVal: 4,
			want:      []int{1, 4, 2, 3},
			wantErr:   nil,
		},
		{
			name:      "insert at the end of a list",
			elements:  []int{1, 2, 3},
			insertIdx: 3,
			insertVal: 4,
			want:      []int{1, 2, 3, 4},
			wantErr:   nil,
		},
		{
			name:      "insert at an index larger than list length",
			elements:  []int{1, 2, 3},
			insertIdx: 4,
			insertVal: 4,
			want:      nil,
			wantErr:   ErrIndexOutOfRange,
		},
		{
			name:      "insert at a negative index",
			elements:  []int{1, 2, 3},
			insertIdx: -1,
			insertVal: 4,
			want:      nil,
			wantErr:   ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)+1))
			for _, element := range tt.elements {
				l.Append(element)
			}

			err := l.Insert(tt.insertIdx, tt.insertVal)
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestList_InsertAll(t *testing.T) {
	tests := []struct {
		name        string
		elements    []int
		insertIdx   int
		insertElems []int
		want        []int
		wantErr     error
	}{
		{
			name:        "insert multiple elements into an empty list",
			elements:    []int{},
			insertIdx:   0,
			insertElems: []int{1, 2, 3},
			want:        []int{1, 2, 3},
			wantErr:     nil,
		},
		{
			name:        "insert multiple elements at the beginning of a list",
			elements:    []int{1, 2, 3},
			insertIdx:   0,
			insertElems: []int{4, 5},
			want:        []int{4, 5, 1, 2, 3},
			wantErr:     nil,
		},
		{
			name:        "insert multiple elements in the middle of a list",
			elements:    []int{1, 2, 3},
			insertIdx:   1,
			insertElems: []int{4, 5},
			want:        []int{1, 4, 5, 2, 3},
			wantErr:     nil,
		},
		{
			name:        "insert multiple elements at the end of a list",
			elements:    []int{1, 2, 3},
			insertIdx:   3,
			insertElems: []int{4, 5},
			want:        []int{1, 2, 3, 4, 5},
			wantErr:     nil,
		},
		{
			name:        "insert multiple elements at an index larger than list length",
			elements:    []int{1, 2, 3},
			insertIdx:   4,
			insertElems: []int{4, 5},
			want:        []int{1, 2, 3},
			wantErr:     ErrIndexOutOfRange,
		},
		{
			name:        "insert multiple elements at a negative index",
			elements:    []int{1, 2, 3},
			insertIdx:   -1,
			insertElems: []int{4, 5},
			want:        []int{1, 2, 3},
			wantErr:     ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)+len(tt.insertElems)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			err := l.InsertAll(tt.insertIdx, tt.insertElems)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, l.elements)
		})
	}
}

func TestList_Get(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		getIdx    int
		want      int
		wantError error
	}{
		{
			name:      "get from an empty list",
			elements:  []int{},
			getIdx:    0,
			want:      0,
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "get at the beginning of a list",
			elements:  []int{1, 2, 3},
			getIdx:    0,
			want:      1,
			wantError: nil,
		},
		{
			name:      "get in the middle of a list",
			elements:  []int{1, 2, 3},
			getIdx:    1,
			want:      2,
			wantError: nil,
		},
		{
			name:      "get at the end of a list",
			elements:  []int{1, 2, 3},
			getIdx:    2,
			want:      3,
			wantError: nil,
		},
		{
			name:      "get at an index larger than list length",
			elements:  []int{1, 2, 3},
			getIdx:    3,
			want:      0,
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "get at a negative index",
			elements:  []int{1, 2, 3},
			getIdx:    -1,
			want:      0,
			wantError: ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got, err := l.Get(tt.getIdx)
			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		find     int
		want     int
	}{
		{
			name:     "find element in an empty list",
			elements: []int{},
			find:     1,
			want:     -1,
		},
		{
			name:     "find element at the beginning of a list",
			elements: []int{1, 2, 3},
			find:     1,
			want:     0,
		},
		{
			name:     "find element in the middle of a list",
			elements: []int{1, 2, 3},
			find:     2,
			want:     1,
		},
		{
			name:     "find element at the end of a list",
			elements: []int{1, 2, 3},
			find:     3,
			want:     2,
		},
		{
			name:     "find non-existing element in a list",
			elements: []int{1, 2, 3},
			find:     4,
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got := l.IndexOf(tt.find)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_LastIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		find     int
		want     int
	}{
		{
			name:     "find element in an empty list",
			elements: []int{},
			find:     1,
			want:     -1,
		},
		{
			name:     "find element at the beginning of a list",
			elements: []int{1, 2, 3, 1},
			find:     1,
			want:     3,
		},
		{
			name:     "find element in the middle of a list",
			elements: []int{1, 2, 3, 2, 4},
			find:     2,
			want:     3,
		},
		{
			name:     "find element at the end of a list",
			elements: []int{1, 2, 3, 4},
			find:     4,
			want:     3,
		},
		{
			name:     "find non-existing element in a list",
			elements: []int{1, 2, 3},
			find:     5,
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got := l.LastIndexOf(tt.find)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_Remove(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		removeIdx int
		want      []int
		wantError error
	}{
		{
			name:      "remove from an empty list",
			elements:  []int{},
			removeIdx: 0,
			want:      []int{},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "remove at the beginning of a list",
			elements:  []int{1, 2, 3},
			removeIdx: 0,
			want:      []int{2, 3},
			wantError: nil,
		},
		{
			name:      "remove in the middle of a list",
			elements:  []int{1, 2, 3},
			removeIdx: 1,
			want:      []int{1, 3},
			wantError: nil,
		},
		{
			name:      "remove at the end of a list",
			elements:  []int{1, 2, 3},
			removeIdx: 2,
			want:      []int{1, 2},
			wantError: nil,
		},
		{
			name:      "remove at an index larger than list length",
			elements:  []int{1, 2, 3},
			removeIdx: 3,
			want:      []int{1, 2, 3},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "remove at a negative index",
			elements:  []int{1, 2, 3},
			removeIdx: -1,
			want:      []int{1, 2, 3},
			wantError: ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}
			err := l.Remove(tt.removeIdx)
			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.want, l.elements)
		})
	}
}

func TestList_RemoveUnorderedAtIndex(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		removeIdx int
		want      []int
		wantError error
	}{
		{
			name:      "remove from an empty list",
			elements:  []int{},
			removeIdx: 0,
			want:      []int{},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "remove at the beginning of a list",
			elements:  []int{1, 2, 3},
			removeIdx: 0,
			want:      []int{3, 2},
			wantError: nil,
		},
		{
			name:      "remove in the middle of a list",
			elements:  []int{1, 2, 3},
			removeIdx: 1,
			want:      []int{1, 3},
			wantError: nil,
		},
		{
			name:      "remove at the end of a list",
			elements:  []int{1, 2, 3},
			removeIdx: 2,
			want:      []int{1, 2},
			wantError: nil,
		},
		{
			name:      "remove at an index larger than list length",
			elements:  []int{1, 2, 3},
			removeIdx: 3,
			want:      []int{1, 2, 3},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "remove at a negative index",
			elements:  []int{1, 2, 3},
			removeIdx: -1,
			want:      []int{1, 2, 3},
			wantError: ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			err := l.RemoveUnorderedAtIndex(tt.removeIdx)
			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.want, l.elements)
		})
	}
}

func TestList_PopAtIndex(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		popIdx    int
		want      int
		wantList  []int
		wantError error
	}{
		{
			name:      "pop from an empty list",
			elements:  []int{},
			popIdx:    0,
			want:      0,
			wantList:  []int{},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "pop at the beginning of a list",
			elements:  []int{1, 2, 3},
			popIdx:    0,
			want:      1,
			wantList:  []int{2, 3},
			wantError: nil,
		},
		{
			name:      "pop in the middle of a list",
			elements:  []int{1, 2, 3},
			popIdx:    1,
			want:      2,
			wantList:  []int{1, 3},
			wantError: nil,
		},
		{
			name:      "pop at the end of a list",
			elements:  []int{1, 2, 3},
			popIdx:    2,
			want:      3,
			wantList:  []int{1, 2},
			wantError: nil,
		},
		{
			name:      "pop at an index larger than list length",
			elements:  []int{1, 2, 3},
			popIdx:    3,
			want:      0,
			wantList:  []int{1, 2, 3},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "pop at a negative index",
			elements:  []int{1, 2, 3},
			popIdx:    -1,
			want:      0,
			wantList:  []int{1, 2, 3},
			wantError: ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got, err := l.PopAtIndex(tt.popIdx)
			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantList, l.elements)
		})
	}
}

func TestList_Pop(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		want      int
		wantList  []int
		wantError error
	}{
		{
			name:      "pop from an empty list",
			elements:  []int{},
			want:      0,
			wantList:  []int{},
			wantError: ErrEmptyList,
		},
		{
			name:      "pop from a list with one element",
			elements:  []int{1},
			want:      1,
			wantList:  []int{},
			wantError: nil,
		},
		{
			name:      "pop from a list with multiple elements",
			elements:  []int{1, 2, 3},
			want:      3,
			wantList:  []int{1, 2},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got, err := l.Pop()
			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantList, l.elements)
		})
	}
}

func TestList_PopFront(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		want      int
		wantList  []int
		wantError error
	}{
		{
			name:      "pop front from an empty list",
			elements:  []int{},
			want:      0,
			wantList:  []int{},
			wantError: ErrEmptyList,
		},
		{
			name:      "pop front from a list with one element",
			elements:  []int{1},
			want:      1,
			wantList:  []int{},
			wantError: nil,
		},
		{
			name:      "pop front from a list with multiple elements",
			elements:  []int{1, 2, 3},
			want:      1,
			wantList:  []int{2, 3},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got, err := l.PopFront()
			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantList, l.elements)
		})
	}
}

func TestList_Set(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		index     int
		value     int
		wantList  []int
		wantError error
	}{
		{
			name:      "set in an empty list",
			elements:  []int{},
			index:     0,
			value:     1,
			wantList:  []int{},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "set at a negative index",
			elements:  []int{1},
			index:     -1,
			value:     2,
			wantList:  []int{1},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "set at an out-of-range index",
			elements:  []int{1},
			index:     1,
			value:     2,
			wantList:  []int{1},
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "set at a valid index",
			elements:  []int{1, 2, 3},
			index:     1,
			value:     4,
			wantList:  []int{1, 4, 3},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			err := l.Set(tt.index, tt.value)
			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.wantList, l.elements)
		})
	}
}

func TestList_Clear(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
	}{
		{
			name:     "clear an empty list",
			elements: []int{},
		},
		{
			name:     "clear a list with one element",
			elements: []int{1},
		},
		{
			name:     "clear a list with multiple elements",
			elements: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			l.Clear()

			assert.Equal(t, 0, len(l.elements))
		})
	}
}

func TestList_Contains(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		values   []int
		want     bool
	}{
		{
			name:     "empty list contains nothing",
			elements: []int{},
			values:   []int{1},
			want:     false,
		},
		{
			name:     "single element list contains its element",
			elements: []int{1},
			values:   []int{1},
			want:     true,
		},
		{
			name:     "single element list does not contain other elements",
			elements: []int{1},
			values:   []int{2},
			want:     false,
		},
		{
			name:     "multiple elements list contains its elements",
			elements: []int{1, 2, 3},
			values:   []int{1, 3},
			want:     true,
		},
		{
			name:     "multiple elements list does not contain other elements",
			elements: []int{1, 2, 3},
			values:   []int{4},
			want:     false,
		},
		{
			name:     "multiple elements list contains some but not all values",
			elements: []int{1, 2, 3},
			values:   []int{1, 4},
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got := l.Contains(tt.values...)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_RemoveRange(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		fromIndex int
		toIndex   int
		want      []int
		wantError error
	}{
		{
			name:      "remove range from an empty list",
			elements:  []int{},
			fromIndex: 0,
			toIndex:   0,
			want:      []int{},
			wantError: ErrFormIndexMustBeLessThanToIndex,
		},
		{
			name:      "remove range with fromIndex == toIndex",
			elements:  []int{1, 2, 3},
			fromIndex: 1,
			toIndex:   1,
			want:      []int{1, 2, 3},
			wantError: ErrFormIndexMustBeLessThanToIndex,
		},
		{
			name:      "remove range with fromIndex > toIndex",
			elements:  []int{1, 2, 3},
			fromIndex: 2,
			toIndex:   1,
			want:      []int{1, 2, 3},
			wantError: ErrFormIndexMustBeLessThanToIndex,
		},
		{
			name:      "remove range from single element list",
			elements:  []int{1},
			fromIndex: 0,
			toIndex:   1,
			want:      []int{},
			wantError: nil,
		},
		{
			name:      "remove range from multiple elements list",
			elements:  []int{1, 2, 3, 4, 5},
			fromIndex: 1,
			toIndex:   4,
			want:      []int{1, 5},
			wantError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			err := l.RemoveRange(tt.fromIndex, tt.toIndex)

			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.want, l.elements)
		})
	}
}

func TestList_SubList(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		fromIndex int
		toIndex   int
		want      []int
		wantError error
	}{
		{
			name:      "sub list from an empty list",
			elements:  []int{},
			fromIndex: 0,
			toIndex:   0,
			want:      nil,
			wantError: ErrFormIndexMustBeLessThanToIndex,
		},
		{
			name:      "sub list with fromIndex == toIndex",
			elements:  []int{1, 2, 3},
			fromIndex: 1,
			toIndex:   1,
			want:      nil,
			wantError: ErrFormIndexMustBeLessThanToIndex,
		},
		{
			name:      "sub list with fromIndex > toIndex",
			elements:  []int{1, 2, 3},
			fromIndex: 2,
			toIndex:   1,
			want:      nil,
			wantError: ErrFormIndexMustBeLessThanToIndex,
		},
		{
			name:      "sub list from single element list",
			elements:  []int{1},
			fromIndex: 0,
			toIndex:   1,
			want:      []int{1},
			wantError: nil,
		},
		{
			name:      "sub list from multiple elements list",
			elements:  []int{1, 2, 3, 4, 5},
			fromIndex: 1,
			toIndex:   4,
			want:      []int{2, 3, 4},
			wantError: nil,
		},
		{
			name:      "sub list with toIndex out of range",
			elements:  []int{1, 2, 3, 4, 5},
			fromIndex: 1,
			toIndex:   6,
			want:      nil,
			wantError: ErrIndexOutOfRange,
		},
		{
			name:      "sub list with fromIndex out of range",
			elements:  []int{1, 2, 3, 4, 5},
			fromIndex: -1,
			toIndex:   4,
			want:      nil,
			wantError: ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got, err := l.SubList(tt.fromIndex, tt.toIndex)

			assert.ErrorIs(t, err, tt.wantError)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		want     []int
	}{
		{
			name:     "reverse an empty list",
			elements: []int{},
			want:     []int{},
		},
		{
			name:     "reverse a single element list",
			elements: []int{1},
			want:     []int{1},
		},
		{
			name:     "reverse a two elements list",
			elements: []int{1, 2},
			want:     []int{2, 1},
		},
		{
			name:     "reverse a multiple elements list",
			elements: []int{1, 2, 3, 4, 5},
			want:     []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			l.Reverse()

			assert.Equal(t, tt.want, l.elements)
		})
	}
}

func TestList_RemoveIf(t *testing.T) {
	tests := []struct {
		name      string
		elements  []int
		predicate func(int) bool
		want      []int
		hasChange bool
	}{
		{
			name:      "remove from an empty list",
			elements:  []int{},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int{},
			hasChange: false,
		},
		{
			name:      "remove none from a list",
			elements:  []int{1, 2, 3, 4, 5},
			predicate: func(x int) bool { return x > 5 },
			want:      []int{1, 2, 3, 4, 5},
			hasChange: false,
		},
		{
			name:      "remove some from a list",
			elements:  []int{1, 2, 3, 4, 5},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      []int{1, 3, 5},
			hasChange: true,
		},
		{
			name:      "remove all from a list",
			elements:  []int{1, 2, 3, 4, 5},
			predicate: func(x int) bool { return x >= 1 },
			want:      []int{},
			hasChange: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			changed := l.RemoveIf(tt.predicate)

			assert.Equal(t, tt.want, l.elements)
			assert.Equal(t, tt.hasChange, changed)
		})
	}
}

func TestList_Sort(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		want     []int
	}{
		{
			name:     "sort empty list",
			elements: []int{},
			want:     []int{},
		},
		{
			name:     "sort list with one element",
			elements: []int{1},
			want:     []int{1},
		},
		{
			name:     "sort list with multiple elements",
			elements: []int{5, 1, 3, 2, 4},
			want:     []int{1, 2, 3, 4, 5},
		},
		{
			name:     "sort list with duplicate elements",
			elements: []int{5, 1, 3, 2, 4, 3, 2},
			want:     []int{1, 2, 2, 3, 3, 4, 5},
		},
		{
			name:     "sort already sorted list",
			elements: []int{1, 2, 3, 4, 5},
			want:     []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			l.Sort()

			assert.Equal(t, tt.want, l.elements)
		})
	}
}

func TestList_Copy(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
	}{
		{
			name:     "copy empty list",
			elements: []int{},
		},
		{
			name:     "copy list with one element",
			elements: []int{1},
		},
		{
			name:     "copy list with multiple elements",
			elements: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "copy list with duplicate elements",
			elements: []int{5, 1, 3, 2, 4, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			got := l.Copy()

			assert.Equal(t, l.elements, got)

			// Change original list and make sure the copy doesn't change
			l.Append(10)
			assert.NotEqual(t, l.elements, got)
		})
	}
}

func TestList_BatchSplit(t *testing.T) {
	tests := []struct {
		name        string
		elements    []int
		batchSize   int
		wantBatches [][]int
	}{
		{
			name:        "split empty list",
			elements:    []int{},
			batchSize:   2,
			wantBatches: [][]int{},
		},
		{
			name:      "split list with one element",
			elements:  []int{1},
			batchSize: 2,
			wantBatches: [][]int{
				{1},
			},
		},
		{
			name:      "split list with multiple elements, size less than batch size",
			elements:  []int{1, 2, 3},
			batchSize: 5,
			wantBatches: [][]int{
				{1, 2, 3},
			},
		},
		{
			name:      "split list with multiple elements, size equals batch size",
			elements:  []int{1, 2, 3, 4, 5},
			batchSize: 5,
			wantBatches: [][]int{
				{1, 2, 3, 4, 5},
			},
		},
		{
			name:      "split list with multiple elements, size greater than batch size",
			elements:  []int{1, 2, 3, 4, 5, 6},
			batchSize: 2,
			wantBatches: [][]int{
				{1, 2},
				{3, 4},
				{5, 6},
			},
		},
		{
			name:      "split list with multiple elements, size is not a multiple of batch size",
			elements:  []int{1, 2, 3, 4, 5, 6, 7},
			batchSize: 3,
			wantBatches: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7},
			},
		},
		{
			name:        "split list with negative batch size",
			elements:    []int{1, 2, 3, 4, 5},
			batchSize:   -1,
			wantBatches: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			gotBatches := l.BatchSplit(tt.batchSize)

			assert.Equal(t, tt.wantBatches, gotBatches)
		})
	}
}

func TestList_SlidingWindows(t *testing.T) {
	tests := []struct {
		name        string
		elements    []int
		windowSize  int
		wantWindows [][]int
	}{
		{
			name:        "empty list",
			elements:    []int{},
			windowSize:  2,
			wantWindows: [][]int{},
		},
		{
			name:       "list with one element",
			elements:   []int{1},
			windowSize: 2,
			wantWindows: [][]int{
				{1},
			},
		},
		{
			name:       "list with multiple elements, size less than window size",
			elements:   []int{1, 2, 3},
			windowSize: 5,
			wantWindows: [][]int{
				{1, 2, 3},
			},
		},
		{
			name:       "list with multiple elements, size equals window size",
			elements:   []int{1, 2, 3, 4, 5},
			windowSize: 5,
			wantWindows: [][]int{
				{1, 2, 3, 4, 5},
			},
		},
		{
			name:       "list with multiple elements, size greater than window size",
			elements:   []int{1, 2, 3, 4, 5, 6},
			windowSize: 2,
			wantWindows: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
				{4, 5},
				{5, 6},
			},
		},
		{
			name:        "list with negative window size",
			elements:    []int{1, 2, 3, 4, 5},
			windowSize:  -1,
			wantWindows: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			gotWindows := l.SlidingWindows(tt.windowSize)

			assert.Equal(t, tt.wantWindows, gotWindows)
		})
	}
}

func TestList_BringElementToFront(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		element  int
		want     []int
	}{
		{
			name:     "empty list",
			elements: []int{},
			element:  1,
			want:     []int{1},
		},
		{
			name:     "list with one element, element is at front",
			elements: []int{1},
			element:  1,
			want:     []int{1},
		},
		{
			name:     "list with one element, element is not in list",
			elements: []int{1},
			element:  2,
			want:     []int{2, 1},
		},
		{
			name:     "list with multiple elements, element is in list",
			elements: []int{1, 2, 3},
			element:  2,
			want:     []int{2, 1, 3},
		},
		{
			name:     "list with multiple elements, element is not in list",
			elements: []int{1, 2, 3},
			element:  4,
			want:     []int{4, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New[int](common.IntsCompare, WithInitialCapacity[int](len(tt.elements)))
			for _, element := range tt.elements {
				l.Append(element)
			}

			l.BringElementToFront(tt.element)

			assert.Equal(t, tt.want, l.elements)
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

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			list := New(common.IntsCompare, WithInitialCapacity[int](len(tc.input)))
			list.elements = append(list.elements, tc.input...)

			list.RemoveDuplicates()
			assert.Equal(t, tc.expected, list.elements)
		})
	}
}

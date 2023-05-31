package singlylinkedlist

import (
	"github.com/kwstars/goads/pkg/common"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_New(t *testing.T) {
	tests := []struct {
		name string
		cmp  interface{}
	}{
		{"int list", common.IntComparator},
		{"string list", common.StringComparator},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch cmp := tt.cmp.(type) {
			case func(a, b int) int8:
				list := New(cmp)
				assert.NotNil(t, list)
				assert.NotNil(t, list.head)
				assert.Equal(t, list.head, list.tail)
				assert.Equal(t, 0, list.head.value)
				assert.Nil(t, list.head.next)
			case func(a, b string) int8:
				list := New(cmp)
				assert.NotNil(t, list)
				assert.NotNil(t, list.head)
				assert.Equal(t, list.head, list.tail)
				assert.Equal(t, "", list.head.value)
				assert.Nil(t, list.head.next)
			}
		})
	}
}

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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			assert.Equal(t, tt.wantSize, list.Size())
			assert.Equal(t, tt.wantTail, list.tail.value)
		})
	}
}

func TestList_Prepend(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		wantSize int
		wantHead int
		wantTail int
	}{
		{"prepend one element", []int{1}, 1, 1, 1},
		{"prepend two elements", []int{1, 2}, 2, 2, 1},
		{"prepend three elements", []int{1, 2, 3}, 3, 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Prepend(v)
			}

			assert.Equal(t, tt.wantSize, list.Size())
			assert.Equal(t, tt.wantHead, list.head.next.value)
			if list.tail != list.head {
				assert.Equal(t, tt.wantTail, list.tail.value)
			}
		})
	}
}

func TestList_Get(t *testing.T) {
	tests := []struct {
		name     string
		values   []int
		getIndex int
		want     int
		wantErr  error
	}{
		{"get first element", []int{1}, 0, 1, nil},
		{"get middle element", []int{1, 2, 3}, 1, 2, nil},
		{"get last element", []int{1, 2, 3}, 2, 3, nil},
		{"get from empty list", []int{}, 0, 0, ErrIndexOutOfRange},
		{"get out of range element", []int{1, 2, 3}, 3, 0, ErrIndexOutOfRange},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)

			for _, v := range tt.values {
				list.Append(v)
			}

			got, err := list.Get(tt.getIndex)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_Insert(t *testing.T) {
	tests := []struct {
		name    string
		values  []int
		inserts []struct {
			index int
			value int
		}
		want    []int
		wantErr error
	}{
		{
			name:   "insert to an empty list",
			values: []int{},
			inserts: []struct {
				index int
				value int
			}{{0, 1}},
			want:    []int{1},
			wantErr: nil,
		},
		{
			name:   "insert at the start",
			values: []int{2, 3},
			inserts: []struct {
				index int
				value int
			}{{0, 1}},
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name:   "insert in the middle",
			values: []int{1, 3},
			inserts: []struct {
				index int
				value int
			}{{1, 2}},
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name:   "insert at the end",
			values: []int{1, 2},
			inserts: []struct {
				index int
				value int
			}{{2, 3}},
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
		{
			name:   "insert out of range",
			values: []int{1, 2, 3},
			inserts: []struct {
				index int
				value int
			}{{4, 4}},
			want:    []int{1, 2, 3},
			wantErr: ErrIndexOutOfRange,
		},
		{
			name:   "multiple inserts",
			values: []int{1},
			inserts: []struct {
				index int
				value int
			}{{1, 3}, {1, 2}},
			want:    []int{1, 2, 3},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			for _, insert := range tt.inserts {
				err := list.Insert(insert.index, insert.value)
				assert.ErrorIs(t, err, tt.wantErr)
			}

			for i, v := range tt.want {
				got, err := list.Get(i)
				assert.NoError(t, err)
				assert.Equal(t, v, got)
			}
		})
	}
}

func TestList_InsertAll(t *testing.T) {
	tests := []struct {
		name    string
		values  []int
		inserts struct {
			index  int
			values []int
		}
		want    []int
		wantErr error
	}{
		{
			"insert all to an empty list",
			[]int{},
			struct {
				index  int
				values []int
			}{0, []int{1, 2, 3}},
			[]int{1, 2, 3},
			nil,
		},
		{
			"insert all at the start",
			[]int{3, 4, 5},
			struct {
				index  int
				values []int
			}{0, []int{1, 2}},
			[]int{1, 2, 3, 4, 5},
			nil,
		},
		{
			"insert all in the middle",
			[]int{1, 5},
			struct {
				index  int
				values []int
			}{1, []int{2, 3, 4}},
			[]int{1, 2, 3, 4, 5},
			nil,
		},
		{
			"insert all at the end",
			[]int{1, 2},
			struct {
				index  int
				values []int
			}{2, []int{3, 4, 5}},
			[]int{1, 2, 3, 4, 5},
			nil,
		},
		{
			"insert all out of range",
			[]int{1, 2, 3},
			struct {
				index  int
				values []int
			}{4, []int{4, 5}},
			[]int{1, 2, 3},
			ErrIndexOutOfRange,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.InsertAll(tt.inserts.index, tt.inserts.values)
			assert.ErrorIs(t, err, tt.wantErr)

			for i, v := range tt.want {
				got, err := list.Get(i)
				assert.NoError(t, err)
				assert.Equal(t, v, got)
			}
		})
	}
}

func TestList_Clear(t *testing.T) {
	tests := []struct {
		name   string
		values []int
	}{
		{
			"clear an empty list",
			[]int{},
		},
		{
			"clear a list with one element",
			[]int{1},
		},
		{
			"clear a list with multiple elements",
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			list.Clear()
			assert.True(t, list.Empty())
			assert.Equal(t, 0, list.Size())
		})
	}
}

func TestList_Remove(t *testing.T) {
	tests := []struct {
		name        string
		values      []int
		removeIndex int
		wantErr     error
		wantList    []int
	}{
		{
			name:        "remove from an empty list",
			values:      []int{},
			removeIndex: 0,
			wantErr:     ErrIndexOutOfRange,
			wantList:    []int{},
		},
		{
			name:        "remove from a list with one element",
			values:      []int{1},
			removeIndex: 0,
			wantErr:     nil,
			wantList:    []int{},
		},
		{
			name:        "remove from the beginning of the list",
			values:      []int{1, 2, 3, 4, 5},
			removeIndex: 0,
			wantErr:     nil,
			wantList:    []int{2, 3, 4, 5},
		},
		{
			name:        "remove from the middle of the list",
			values:      []int{1, 2, 3, 4, 5},
			removeIndex: 2,
			wantErr:     nil,
			wantList:    []int{1, 2, 4, 5},
		},
		{
			name:        "remove from the end of the list",
			values:      []int{1, 2, 3, 4, 5},
			removeIndex: 4,
			wantErr:     nil,
			wantList:    []int{1, 2, 3, 4},
		},
		{
			name:        "remove out of range",
			values:      []int{1, 2, 3, 4, 5},
			removeIndex: 5,
			wantErr:     ErrIndexOutOfRange,
			wantList:    []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.Remove(tt.removeIndex)
			assert.ErrorIs(t, err, tt.wantErr)

			for i, v := range tt.wantList {
				got, err := list.Get(i)
				assert.NoError(t, err)
				assert.Equal(t, v, got)
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
		wantErr   error
		wantList  []int
	}{
		{
			name:      "remove range from an empty list",
			values:    []int{},
			fromIndex: 0,
			toIndex:   1,
			wantErr:   ErrIndexOutOfRange,
			wantList:  []int{},
		},
		{
			name:      "remove range from a list with one element",
			values:    []int{1},
			fromIndex: 0,
			toIndex:   0,
			wantErr:   nil,
			wantList:  []int{},
		},
		{
			name:      "remove range from the beginning of the list",
			values:    []int{1, 2, 3, 4, 5},
			fromIndex: 0,
			toIndex:   1,
			wantErr:   nil,
			wantList:  []int{3, 4, 5},
		},
		{
			name:      "remove range from the middle of the list",
			values:    []int{1, 2, 3, 4, 5},
			fromIndex: 1,
			toIndex:   3,
			wantErr:   nil,
			wantList:  []int{1, 5},
		},
		{
			name:      "remove range from the end of the list",
			values:    []int{1, 2, 3, 4, 5},
			fromIndex: 3,
			toIndex:   4,
			wantErr:   nil,
			wantList:  []int{1, 2, 3},
		},
		{
			name:      "remove range out of range",
			values:    []int{1, 2, 3, 4, 5},
			fromIndex: 4,
			toIndex:   5,
			wantErr:   ErrIndexOutOfRange,
			wantList:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.RemoveRange(tt.fromIndex, tt.toIndex)
			assert.ErrorIs(t, err, tt.wantErr)

			for i, v := range tt.wantList {
				got, err := list.Get(i)
				assert.NoError(t, err)
				assert.Equal(t, v, got)
			}
		})
	}
}

func TestList_Set(t *testing.T) {
	tests := []struct {
		name    string
		values  []int
		index   int
		setVal  int
		wantErr error
		want    int
	}{
		{
			name:    "set an element in an empty list",
			values:  []int{},
			index:   0,
			setVal:  1,
			wantErr: ErrIndexOutOfRange,
			want:    0,
		},
		{
			name:    "set an element in a list with one element",
			values:  []int{1},
			index:   0,
			setVal:  2,
			wantErr: nil,
			want:    2,
		},
		{
			name:    "set an element at the beginning of the list",
			values:  []int{1, 2, 3},
			index:   0,
			setVal:  4,
			wantErr: nil,
			want:    4,
		},
		{
			name:    "set an element in the middle of the list",
			values:  []int{1, 2, 3},
			index:   1,
			setVal:  4,
			wantErr: nil,
			want:    4,
		},
		{
			name:    "set an element at the end of the list",
			values:  []int{1, 2, 3},
			index:   2,
			setVal:  4,
			wantErr: nil,
			want:    4,
		},
		{
			name:    "set an element out of range",
			values:  []int{1, 2, 3},
			index:   3,
			setVal:  4,
			wantErr: ErrIndexOutOfRange,
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			err := list.Set(tt.index, tt.setVal)
			assert.ErrorIs(t, err, tt.wantErr)

			if err == nil {
				got, err := list.Get(tt.index)
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestList_IndexOf(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		find   int
		want   int
	}{
		{
			name:   "find an element in an empty list",
			values: []int{},
			find:   1,
			want:   -1,
		},
		{
			name:   "find an non-existent element in the list",
			values: []int{1, 2, 3},
			find:   4,
			want:   -1,
		},
		{
			name:   "find an existing element in the list",
			values: []int{1, 2, 3},
			find:   2,
			want:   1,
		},
		{
			name:   "find an element in a list with duplicates",
			values: []int{1, 2, 2, 3},
			find:   2,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			got := list.IndexOf(tt.find)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestList_LastIndexOf(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		find   int
		want   int
	}{
		{
			name:   "find an element in an empty list",
			values: []int{},
			find:   1,
			want:   -1,
		},
		{
			name:   "find an non-existent element in the list",
			values: []int{1, 2, 3},
			find:   4,
			want:   -1,
		},
		{
			name:   "find an existing element in the list",
			values: []int{1, 2, 3},
			find:   2,
			want:   1,
		},
		{
			name:   "find an element in a list with duplicates",
			values: []int{1, 2, 2, 3},
			find:   2,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			got := list.LastIndexOf(tt.find)
			assert.Equal(t, tt.want, got)
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
		wantErr   error
	}{
		{
			name:      "index out of bounds",
			values:    []int{1, 2, 3},
			fromIndex: -1,
			toIndex:   3,
			want:      nil,
			wantErr:   ErrIndexOutOfRange,
		},
		{
			name:      "fromIndex greater than toIndex",
			values:    []int{1, 2, 3},
			fromIndex: 2,
			toIndex:   1,
			want:      nil,
			wantErr:   ErrIndexOutOfRange,
		},
		{
			name:      "valid sublist within the list",
			values:    []int{1, 2, 3, 4, 5},
			fromIndex: 1,
			toIndex:   3,
			want:      []int{2, 3, 4},
			wantErr:   nil,
		},
		{
			name:      "fromIndex equal to toIndex",
			values:    []int{1, 2, 3, 4, 5},
			fromIndex: 2,
			toIndex:   2,
			want:      []int{3},
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := New(common.IntComparator)
			for _, v := range tt.values {
				list.Append(v)
			}

			got, err := list.SubList(tt.fromIndex, tt.toIndex)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

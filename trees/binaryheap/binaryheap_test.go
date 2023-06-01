package binaryheap

import (
	"errors"
	"math"
	"testing"

	"github.com/kwstars/goads/pkg/common"
)

var (
	IntMaxHeap = func(a, b int) int8 {
		if a > b {
			return 1
		}
		return 0
	}
	IntMinHeap = func(a, b int) int8 {
		if a < b {
			return 1
		}
		return 0
	}
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name     string
		options  []Option[int]
		expCap   int
		expLen   int
		expError error
	}{
		{
			name:     "without initial capacity",
			options:  nil,
			expCap:   0,
			expLen:   0,
			expError: nil,
		},
		{
			name:     "with initial capacity",
			options:  []Option[int]{WithInitialCapacity[int](10)},
			expCap:   10,
			expLen:   0,
			expError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			heap := New(IntMinHeap, tc.options...)
			if len(heap.data) != tc.expLen {
				t.Errorf("expected length of heap data to be %d, got %d", tc.expLen, len(heap.data))
			}
			if cap(heap.data) != tc.expCap {
				t.Errorf("expected capacity of heap data to be %d, got %d", tc.expCap, cap(heap.data))
			}
		})
	}
}

func TestBinaryHeap_Push(t *testing.T) {
	testCases := []struct {
		name     string
		cmp      common.Comparator[int, int]
		input    []int
		pushVal  int
		expected []int
	}{
		// MinHeap
		{name: "push to an empty minheap", cmp: IntMinHeap, input: []int{}, pushVal: 1, expected: []int{1}},
		/*push to a non-empty minheap
		     1            1                1
		    / \          / \             / \
		  3    5  -->   3   5    -->    2   5
		               /               /
		              2               3
		*/
		{name: "push to a non-empty minheap", cmp: IntMinHeap, input: []int{1, 3, 5}, pushVal: 2, expected: []int{1, 2, 5, 3}},
		/*push a larger value to a non-empty minheap
		    1             1
		   / \           / \
		  3   5   -->   3   5
		               /
		             6
		*/
		{name: "push a larger value to a non-empty minheap", cmp: IntMinHeap, input: []int{1, 3, 5}, pushVal: 6, expected: []int{1, 3, 5, 6}},
		{name: "push a large amount of numbers", cmp: IntMinHeap, input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, pushVal: 10, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		/*push a negative number
		    1            1            -1
		   / \          / \           / \
		  2   3  -->   2   3   -->   1   3
		               /             /
		              -1            2
		*/
		{name: "push a negative number", cmp: IntMinHeap, input: []int{1, 2, 3}, pushVal: -1, expected: []int{-1, 1, 3, 2}},
		{name: "pushing IntMin", cmp: IntMinHeap, input: []int{1, 2, 3}, pushVal: math.MinInt32, expected: []int{math.MinInt32, 1, 3, 2}},
		/*push a repeated number
		    1           1             1
		   / \         / \           / \
		  2   3  -->  2   3   -->   1   3
		             /             /
		            1             2
		*/
		{name: "push a repeated number", cmp: IntMinHeap, input: []int{1, 2, 3}, pushVal: 1, expected: []int{1, 1, 3, 2}},

		// MaxHeap
		{name: "push to an empty maxheap", cmp: IntMaxHeap, input: []int{}, pushVal: 1, expected: []int{1}},
		/*push to a non-empty maxheap
		     5              5               5
		    / \            / \             / \
		  1     3   -->   1   3    -->    2   3
		                 /               /
		                2               1
		*/
		{name: "push to a non-empty maxheap", cmp: IntMaxHeap, input: []int{1, 3, 5}, pushVal: 2, expected: []int{5, 2, 3, 1}},
		/*push a larger value to a non-empty maxheap
		     5             5               6
		    / \           / \             / \
		  1    3   -->   1   3    -->    5   3
		                /               /
		               6               1
		*/
		{name: "push a larger value to a non-empty maxheap", cmp: IntMaxHeap, input: []int{1, 3, 5}, pushVal: 6, expected: []int{6, 5, 3, 1}},
		{name: "push a large amount of numbers", cmp: IntMaxHeap, input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, pushVal: 10, expected: []int{10, 9, 6, 7, 8, 2, 5, 1, 4, 3}},
		/*pushing IntMax
		   3             3           MaxInt
		  / \            / \          / \
		1     2   -->   1   2 -->    3   2
		               /            /
		           MaxInt          1
		*/
		{name: "pushing IntMax", cmp: IntMaxHeap, input: []int{1, 2, 3}, pushVal: math.MaxInt32, expected: []int{math.MaxInt32, 3, 2, 1}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			heap := New(tc.cmp, WithInitialCapacity[int](len(tc.input)+1))
			for _, val := range tc.input {
				heap.Push(val)
			}

			heap.Push(tc.pushVal)

			if len(heap.data) != len(tc.expected) {
				t.Errorf("Expected heap size %d, got %d", len(tc.expected), len(heap.data))
			}

			for i, v := range heap.data {
				if v != tc.expected[i] {
					t.Errorf("Expected value at index %d to be %d, got %d", i, tc.expected[i], v)
				}
			}
		})
	}
}

func TestBinaryHeap_Pop(t *testing.T) {
	testCases := []struct {
		name        string
		cmp         common.Comparator[int, int]
		input       []int
		expectedVal int
		expectedErr error
		expected    []int
	}{
		// MinHeap
		{name: "pop from an empty minheap", cmp: IntMinHeap, input: []int{}, expectedErr: ErrHeapEmpty, expected: []int{}},
		/*pop from a non-empty minheap
		     1                      5         3
		    / \         / \        /         /
		  3    5 -->   3   5 -->  3    -->  5
		*/
		{name: "pop from a non-empty minheap", cmp: IntMinHeap, input: []int{1, 3, 5}, expectedVal: 1, expected: []int{3, 5}},
		/*pop from a heap with negative numbers
		     -3                           -2
		    /  \           /  \           /
		  -1   -2 -->    -1   -2   -->  -1
		*/
		{name: "pop from a heap with negative numbers", cmp: IntMinHeap, input: []int{-1, -2, -3}, expectedVal: -3, expected: []int{-2, -1}},
		{name: "pop from a heap with repeated elements", cmp: IntMinHeap, input: []int{1, 1, 1}, expectedVal: 1, expected: []int{1, 1}},
		{name: "pop from a larger minheap", cmp: IntMinHeap, input: []int{1, 2, 3, 5}, expectedVal: 1, expected: []int{2, 5, 3}},

		// MaxHeap
		{name: "pop from an empty maxheap", cmp: IntMaxHeap, input: []int{}, expectedErr: ErrHeapEmpty, expected: []int{}},
		/*pop from a non-empty maxheap
		    5           3
		   / \         /
		  1   3  -->  1
		*/
		{name: "pop from a non-empty maxheap", cmp: IntMaxHeap, input: []int{1, 3, 5}, expectedVal: 5, expected: []int{3, 1}},
		/*pop from a larger maxheap
		      5                            2            4
		     / \            / \           /  \         / \
		    4   3   -->    4   3   -->   4   3   -->  2   3
		   /              /
		  2              2
		*/
		{name: "pop from a larger maxheap", cmp: IntMaxHeap, input: []int{5, 4, 3, 2}, expectedVal: 5, expected: []int{4, 2, 3}},
		{name: "Pop from max heap where right child is larger", cmp: IntMaxHeap, input: []int{7, 6, 5, 4, 3, 2, 1}, expectedVal: 7, expected: []int{6, 4, 5, 1, 3, 2}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			heap := New(tc.cmp, WithInitialCapacity[int](len(tc.input)+1))
			for _, val := range tc.input {
				heap.Push(val)
			}

			val, err := heap.Pop()

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

			if val != tc.expectedVal {
				t.Errorf("Expected pop value to be %d, got %d", tc.expectedVal, val)
			}

			if len(heap.data) != len(tc.expected) {
				t.Errorf("Expected heap size %d, got %d", len(tc.expected), len(heap.data))
			}

			for i, v := range heap.data {
				if v != tc.expected[i] {
					t.Errorf("Expected value at index %d to be %d, got %d", i, tc.expected[i], v)
				}
			}
		})
	}
}

func TestBinaryHeap(t *testing.T) {
	heap := New(IntMinHeap, WithInitialCapacity[int](10))
	for i := 0; i < 10; i++ {
		heap.Push(i)
	}

	tests := []struct {
		name          string
		i             int
		expectedLeft  int
		expectedRight int
		expectedHasL  bool
		expectedHasR  bool
	}{
		{name: "Node at index 0", i: 0, expectedLeft: 1, expectedRight: 2, expectedHasL: true, expectedHasR: true},
		{name: "Node at index 4", i: 4, expectedLeft: 9, expectedRight: 10, expectedHasL: true, expectedHasR: false},
		{name: "Node at index 5", i: 5, expectedLeft: 11, expectedRight: 12, expectedHasL: false, expectedHasR: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if left := heap.leftChild(tt.i); left != tt.expectedLeft {
				t.Errorf("Expected left child index: %d, got: %d", tt.expectedLeft, left)
			}
			if right := heap.rightChild(tt.i); right != tt.expectedRight {
				t.Errorf("Expected right child index: %d, got: %d", tt.expectedRight, right)
			}
			if hasL := heap.hasLeftChild(tt.i); hasL != tt.expectedHasL {
				t.Errorf("Expected has left child: %v, got: %v", tt.expectedHasL, hasL)
			}
			if hasR := heap.hasRightChild(tt.i); hasR != tt.expectedHasR {
				t.Errorf("Expected has right child: %v, got: %v", tt.expectedHasR, hasR)
			}
		})
	}
}

func TestBinaryHeap_Peek(t *testing.T) {
	testCases := []struct {
		name     string
		cmp      common.Comparator[int, int]
		input    []int
		expected int
		err      error
	}{
		{name: "Peek from an empty heap", cmp: IntMinHeap, input: []int{}, expected: 0, err: ErrHeapEmpty},
		{name: "Peek from a min heap", cmp: IntMinHeap, input: []int{1, 3, 5, 2, 6}, expected: 1, err: nil},
		{name: "Peek from a max heap", cmp: IntMaxHeap, input: []int{1, 3, 5, 2, 6}, expected: 6, err: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			heap := New(tc.cmp, WithInitialCapacity[int](len(tc.input)))
			for _, val := range tc.input {
				heap.Push(val)
			}

			val, err := heap.Peek()

			if !errors.Is(err, tc.err) {
				t.Errorf("Expected error %v, got %v", tc.err, err)
			}

			if val != tc.expected {
				t.Errorf("Expected value to be %d, got %d", tc.expected, val)
			}
		})
	}
}

func TestBinaryHeap_Size(t *testing.T) {
	testCases := []struct {
		name     string
		cmp      common.Comparator[int, int]
		input    []int
		expected int
	}{
		{name: "Size of an empty heap", cmp: IntMinHeap, input: []int{}, expected: 0},
		{name: "Size of a heap with one element", cmp: IntMinHeap, input: []int{1}, expected: 1},
		{name: "Size of a heap with multiple elements", cmp: IntMinHeap, input: []int{1, 3, 5, 2, 6}, expected: 5},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			heap := New(tc.cmp, WithInitialCapacity[int](len(tc.input)))
			for _, val := range tc.input {
				heap.Push(val)
			}

			if size := heap.Size(); size != tc.expected {
				t.Errorf("Expected size to be %d, got %d", tc.expected, size)
			}
		})
	}
}

func TestBinaryHeap_IsEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		cmp      common.Comparator[int, int]
		input    []int
		expected bool
	}{
		{name: "Empty heap", cmp: IntMinHeap, input: []int{}, expected: true},
		{name: "Heap with one element", cmp: IntMinHeap, input: []int{1}, expected: false},
		{name: "Heap with multiple elements", cmp: IntMinHeap, input: []int{1, 3, 5, 2, 6}, expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			heap := New(tc.cmp, WithInitialCapacity[int](len(tc.input)))
			for _, val := range tc.input {
				heap.Push(val)
			}

			if isEmpty := heap.IsEmpty(); isEmpty != tc.expected {
				t.Errorf("Expected IsEmpty to return %t, got %t", tc.expected, isEmpty)
			}
		})
	}
}

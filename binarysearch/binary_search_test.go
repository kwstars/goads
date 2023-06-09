package binarysearch

import (
	"testing"
)

func TestFindExact(t *testing.T) {
	t.Parallel()
	type args[T1 any, T2 comparable] struct {
		arr    []T1
		target T2
		cmp    func(a T1, b T2) int8
	}
	type testCase[T1 any, T2 comparable] struct {
		name string
		args args[T1, T2]
		want int
	}

	intsTests := []testCase[int, int]{
		{
			name: "search x in int slice",
			args: args[int, int]{
				arr:    []int{1, 2, 3, 4, 5},
				target: 4,
				cmp:    IntComparator},
			want: 3,
		},
	}

	t.Run("search target", func(t *testing.T) {
		for _, tt := range intsTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := FindExact(tt.args.arr, tt.args.target, tt.args.cmp); got != tt.want {
					t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	rangeTests := []testCase[Range, int]{
		{
			name: "start of range",
			args: args[Range, int]{
				arr:    []Range{{Start: 1, End: 2}, {Start: 3, End: 4}, {Start: 5, End: 100}},
				target: 1,
				cmp:    rangeComparator,
			},
			want: 0,
		},
		{
			name: "end of range",
			args: args[Range, int]{
				arr:    []Range{{Start: 1, End: 2}, {Start: 3, End: 4}, {Start: 5, End: 100}},
				target: 100,
				cmp:    rangeComparator,
			},
			want: 2,
		},
		{
			name: "target not found in start",
			args: args[Range, int]{
				arr:    []Range{{Start: 1, End: 2}, {Start: 3, End: 4}, {Start: 5, End: 100}},
				target: 101,
				cmp:    rangeComparator,
			},
			want: -1,
		},
		{
			name: "target not found in end ",
			args: args[Range, int]{
				arr:    []Range{{Start: 1, End: 2}, {Start: 3, End: 4}, {Start: 5, End: 100}},
				target: 0,
				cmp:    rangeComparator,
			},
			want: -1,
		},
		{
			name: "target not found",
			args: args[Range, int]{
				arr:    []Range{{Start: 1, End: 2}, {Start: 7, End: 10}, {Start: 20, End: 100}},
				target: 15,
				cmp:    rangeComparator,
			},
			want: -1,
		},
	}

	t.Run("search in range", func(t *testing.T) {
		for _, tt := range rangeTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := FindExact(tt.args.arr, tt.args.target, tt.args.cmp); got != tt.want {
					t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestFindFirstGreaterOrEqual(t *testing.T) {
	t.Parallel()

	type args[T1 any, T2 comparable] struct {
		arr    []T1
		target T2
		cmp    func(a T1, b T2) int8
	}
	type testCase[T1 any, T2 comparable] struct {
		name string
		args args[T1, T2]
		want int
	}

	intTests := []testCase[int, int]{
		{name: "find middle value", args: args[int, int]{arr: []int{1, 2, 5, 6, 7, 8, 9}, target: 4, cmp: IntComparator}, want: 2},
		{name: "not found", args: args[int, int]{arr: []int{1, 2, 5, 6, 7, 8, 9}, target: 10, cmp: IntComparator}, want: -1},
		{name: "find leftmost value", args: args[int, int]{arr: []int{1, 2, 5, 6, 7, 8, 9}, target: 0, cmp: IntComparator}, want: 0},
		{name: "null slice", args: args[int, int]{arr: []int{}, target: 6, cmp: IntComparator}, want: -1},
	}

	t.Run("search in int slice", func(t *testing.T) {
		for _, tt := range intTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := FindFirstGreaterOrEqual(tt.args.arr, tt.args.target, tt.args.cmp); got != tt.want {
					t.Errorf("FindNextBiggest() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestFindLastLessOrEqual(t *testing.T) {
	t.Parallel()

	type args[T1 any, T2 comparable] struct {
		arr    []T1
		target T2
		cmp    func(a T1, b T2) int8
	}
	type testCase[T1 any, T2 comparable] struct {
		name string
		args args[T1, T2]
		want int
	}
	intTests := []testCase[int, int]{
		{name: "find middle value", args: args[int, int]{arr: []int{1, 2, 5, 6, 7, 8, 9}, target: 4, cmp: IntComparator}, want: 1},
		{name: "find middle value", args: args[int, int]{arr: []int{1, 2, 5, 6, 7, 8, 9}, target: 5, cmp: IntComparator}, want: 2},
		{name: "find rightmost value", args: args[int, int]{arr: []int{1, 2, 5, 6, 7, 8, 9}, target: 10, cmp: IntComparator}, want: 6},
		{name: "find leftmost value", args: args[int, int]{arr: []int{1, 2, 5, 6, 7, 8, 9}, target: 0, cmp: IntComparator}, want: -1},
		{name: "null slice", args: args[int, int]{arr: []int{}, target: 5, cmp: IntComparator}, want: -1},
	}
	t.Run("search in int slice", func(t *testing.T) {
		for _, tt := range intTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := FindLastLessOrEqual(tt.args.arr, tt.args.target, tt.args.cmp); got != tt.want {
					t.Errorf("FindLastNotBiggerThan() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

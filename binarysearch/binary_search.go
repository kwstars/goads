package binarysearch

import "github.com/kwstars/goads/pkg/common"

// FindExact performs binary search to find the index of the element equal to the target.
// It searches for the target in arr using the provided comparison function cmp.
// cmp should return a negative number if a < b, zero if a == b, and a positive number if a > b.
// If the target is found, it returns the index of the target. Otherwise, it returns -1.
func FindExact[T1 any, T2 comparable](arr []T1, target T2, cmp common.Comparator[T1, T2]) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		comparison := cmp(arr[mid], target)
		if comparison == 0 {
			return mid
		}
		if comparison < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// FindFirstGreaterOrEqual finds the index of the first element that is greater than or equal to the target.
// It uses the provided comparison function cmp to compare elements.
// cmp should return a negative number if a < b, zero if a == b, and a positive number if a > b.
// If such an element is not found, it returns -1.
func FindFirstGreaterOrEqual[T1 any, T2 comparable](arr []T1, target T2, cmp common.Comparator[T1, T2]) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)>>1
		if cmp(arr[mid], target) >= 0 {
			if mid == 0 || cmp(arr[mid-1], target) < 0 {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// FindLastLessOrEqual finds the index of the last element that is less than or equal to the target.
// It uses the provided comparison function cmp to compare elements.
// cmp should return a negative number if a < b, zero if a == b, and a positive number if a > b.
// If such an element is not found, it returns -1.
func FindLastLessOrEqual[T1 any, T2 comparable](arr []T1, target T2, cmp common.Comparator[T1, T2]) int {
	left, right := 0, len(arr)-1
	result := -1

	for left <= right {
		mid := left + (right-left)>>1
		if cmp(arr[mid], target) <= 0 {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

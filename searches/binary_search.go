// author: wongoo@apache.org

package searches

import "github.com/vogo/goalg/compare"

// BinarySearch binary search the target, return the index when first found.
func BinarySearch(arr []compare.Comparer, target compare.Comparer) int {
	l, r, mid := 0, len(arr)-1, 0

	// the left and right index may be equal
	for l <= r {

		// find the mid
		// not use (l+r)/2 to avoid overflow
		// and the expression `l + (r-l)/2` is better to explain.
		mid = l + (r-l)/2

		switch target.Compare(arr[mid]) {
		case 0:
			return mid
		case -1:
			r = mid - 1
		case 1:
			l = mid + 1
		}
	}

	// not found
	return -1
}

// BinarySearchLeftBound binary search the target, return the left bound index when found duplicates.
func BinarySearchLeftBound(arr []compare.Lesser, target compare.Lesser) int {
	if len(arr) == 0 {
		return -1
	}

	l, r, mid := 0, len(arr)-1, 0

	// exit loop when left == right
	for l < r {

		// find the mid
		// `l + (r-l)/2` get the left index when the left and right are neighbor,
		mid = l + (r-l)/2

		// check whether the target is in the right half.
		if arr[mid].Less(target) {
			l = mid + 1
		} else {
			// arr[mid] >= target
			// keep binary search in the left half even when equals to close to the left bound.
			r = mid
		}
	}

	// check whether target is equal to the final left index.
	if !target.Less(arr[l]) && !arr[l].Less(target) {
		return l
	}

	return -1
}

// BinarySearchRightBound binary search the target, return the right bound index when found duplicates.
func BinarySearchRightBound(arr []compare.Lesser, target compare.Lesser) int {
	if len(arr) == 0 {
		return -1
	}

	l, r, mid := 0, len(arr)-1, 0

	// exit loop when left == right
	for l < r {

		// find the mid
		// `l + (r-l+1)/2` get the right index when the left and right are neighbor,
		mid = l + (r-l+1)/2

		// check whether the target is in the left half.
		if target.Less(arr[mid]) {
			r = mid - 1
		} else {
			// arr[mid] <= target
			// keep binary search in the right half even when equals to close to the right bound.
			l = mid
		}
	}

	// check whether target is equal to the final right index.
	if !target.Less(arr[r]) && !arr[r].Less(target) {
		return l
	}

	return -1
}

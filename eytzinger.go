// Package eytzinger implements Eytzinger Binary Search for any ordered type.
package eytzinger

import (
	"slices"

	"golang.org/x/exp/constraints"
)

// Search does a binary search of an array sorted in Eytzinger order.
func Search[T constraints.Ordered](a []T, x T) int {
	var index int
	for index < len(a) {
		k := a[index]
		if x == k {
			return index
		}
		index = index<<1 | 1
		if k < x {
			index++
		}
	}
	return -1
}

// Sort sorts the given array into Eytzinger order.
func Sort[T constraints.Ordered](a []T) {
	slices.Sort(a)
	sorted := make([]T, len(a))
	eytzinger[T](a, sorted, 0, 1)
	copy(a, sorted)
}

func eytzinger[T constraints.Ordered](in, out []T, i, k int) int {
	if k <= len(in) {
		i = eytzinger(in, out, i, 2*k)
		out[k-1] = in[i]
		i++
		i = eytzinger(in, out, i, 2*k+1)
	}
	return i
}

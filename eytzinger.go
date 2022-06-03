// Package eytzinger implements Eytzinger Binary Search for any ordered type.
package eytzinger

import (
	"golang.org/x/exp/constraints"
	"sort"
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
	s := sorter[T]{
		items: a,
	}
	sort.Sort(s)
	sorted := make([]T, len(a))
	eytzinger[T](a, sorted, 0, 1)
	copy(a, sorted)
}

type sorter[T constraints.Ordered] struct {
	items []T
}

func (s sorter[T]) Len() int           { return len(s.items) }
func (s sorter[T]) Less(i, j int) bool { return s.items[i] < s.items[j] }
func (s sorter[T]) Swap(i, j int)      { s.items[i], s.items[j] = s.items[j], s.items[i] }

func eytzinger[T constraints.Ordered](in, out []T, i, k int) int {
	if k <= len(in) {
		i = eytzinger(in, out, i, 2*k)
		out[k-1] = in[i]
		i++
		i = eytzinger(in, out, i, 2*k+1)
	}
	return i
}

/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/7/19
 */

package lol

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// SortSlice sorts a slice. Just a wrapper for slices.Sort.
func SortSlice[S ~[]E, E constraints.Ordered](s S) {
	slices.Sort(s)
}

// MergeSlice merges multiple slices without removing duplicates or shuffling the elements.
//
// Play: https://go.dev/play/p/ARhoTg83WK8
func MergeSlice[T any](ss ...[]T) []T {
	if len(ss) == 0 {
		return []T{}
	}
	size := 0
	for _, v := range ss {
		size += len(v)
	}

	if size == 0 {
		return []T{}
	}

	var res = make([]T, 0, size)
	for _, v := range ss {
		res = append(res, v...)
	}
	return res
}

// UniqSlice merges multiple slices with removing duplicates the elements.
//
// Play: https://go.dev/play/p/clvg0gFoBQs
func UniqSlice[T comparable](ss ...[]T) []T {
	if len(ss) == 0 {
		return []T{}
	}
	size := 0
	for _, v := range ss {
		size += len(v)
	}

	if size == 0 {
		return []T{}
	}

	var (
		res   = make([]T, 0, size)
		exist = make(map[T]struct{}, size)
	)
	for _, s := range ss {
		for _, v := range s {
			if _, ok := exist[v]; ok {
				continue
			}
			exist[v] = struct{}{}
			res = append(res, v)
		}
	}
	return res
}

// Union returns the union of multiple slices.
func Union[T comparable](ss ...[]T) []T {
	return UniqSlice[T](ss...)
}

// Intersection returns the intersection of multiple slices.
//
// Play: https://go.dev/play/p/tqI_pu_-khj
func Intersection[T comparable](ss ...[]T) []T {
	if len(ss) == 0 {
		return []T{}
	}
	var n, min, max = len(ss), len(ss[0]), len(ss[0])
	for _, v := range ss {
		l := len(v)
		if l > max {
			max = l
		}
		if l < min {
			min = l
		}
	}
	if max == 0 {
		return []T{}
	}

	var (
		res   = make([]T, 0, min)
		count = make(map[T]uint, max)
	)
	for _, s := range ss {
		for _, v := range s {
			count[v] += 1
			if count[v] == uint(n) {
				res = append(res, v)
			}
		}
	}
	return res
}

// Difference returns the difference of s1-s2.
//
// Play: https://go.dev/play/p/fB7HJFYbsCB
func Difference[T comparable](s1, s2 []T) []T {
	if len(s1) == 0 {
		return []T{}
	}
	if len(s2) == 0 {
		return s1
	}

	exist := make(map[T]struct{}, len(s2))
	for _, v := range s2 {
		exist[v] = struct{}{}
	}
	res := make([]T, 0, len(s1)/2)
	for _, v := range s1 {
		if _, ok := exist[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}

// Map maps a function over a slice.
//
// Play: https://go.dev/play/p/ePBYrs1YqDz
func Map[T any](list []T, fn func(v T) T) []T {
	if len(list) == 0 {
		return []T{}
	}
	if fn == nil {
		return list
	}

	for i := range list {
		list[i] = fn(list[i])
	}

	return list
}

// MapTo is same as Map, but it can return another type slice.
//
// Special case: when fn is nil, it returns the empty slice.
//
// PLay: https://go.dev/play/p/wClBNz0Fjxt
func MapTo[T any, R any](list []T, fn func(v T) R) []R {
	if len(list) == 0 {
		return []R{}
	}
	if fn == nil {
		return []R{}
	}

	var res = make([]R, len(list))
	for i := range list {
		res[i] = fn(list[i])
	}

	return res
}

// Reduce accumulates and combines elements through a fn into a single value.
//
// Special case: when fn is nil, it returns the initial value
//
// Play: https://go.dev/play/p/gI8Mcvk4NGr
func Reduce[T any, R any](list []T, fn func(sum R, item T, index int) R, initial R) R {
	if fn == nil {
		return initial
	}
	for i, v := range list {
		initial = fn(initial, v, i)
	}

	return initial
}

// ReduceRight is like the Reduce, but the order is reserve.
//
// Special case: when fn is nil, it returns the initial value
//
// Play: https://go.dev/play/p/n1rGGeg1KFf
func ReduceRight[T any, R any](list []T, fn func(sum R, item T, index int) R, initial R) R {
	if fn == nil {
		return initial
	}
	for i := len(list) - 1; i >= 0; i-- {
		initial = fn(initial, list[i], i)
	}

	return initial
}

// Include returns true if the slice includes the element.
//
// Play: https://go.dev/play/p/6VW3_rG4AIX
func Include[T comparable](list []T, ele T) bool {
	for _, v := range list {
		if v == ele {
			return true
		}
	}
	return false
}

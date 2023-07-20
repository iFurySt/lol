/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/7/19
 */

package lol

// MergeSlice merges multiple slices without removing duplicates or shuffling the elements
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

// UniqSlice merges multiple slices with removing duplicates the elements
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

// Map maps a function over a slice
//
// https://go.dev/play/p/ePBYrs1YqDz
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

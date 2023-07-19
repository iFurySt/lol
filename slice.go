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

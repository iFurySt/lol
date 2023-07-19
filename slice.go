/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/7/19
 */

package lol

// MergeSlice merges multiple slices without removing duplicates or shuffling the elements
//
// Play: https://go.dev/play/p/yAdsTYoAyLn
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

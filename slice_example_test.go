/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/7/19
 */

package lol

import "fmt"

func ExampleMergeSlice() {
	res1 := MergeSlice[int]([]int{1, 4, 7}, []int{2, 5, 8}, []int{3, 6, 9})
	res2 := MergeSlice[int]([]int{1}, []int{2, 5, 8}, []int{3, 6}, []int{4, 7, 9, 10})
	fmt.Println(res1)
	fmt.Println(res2)
	// Output:
	// [1 4 7 2 5 8 3 6 9]
	// [1 2 5 8 3 6 4 7 9 10]
}

func ExampleUniqSlice() {
	res1 := UniqSlice[int]([]int{1, 4, 7}, []int{4, 1, 2}, []int{7, 1, 3})
	res2 := UniqSlice[int]([]int{4, 7, 9, 10})
	fmt.Println(res1)
	fmt.Println(res2)
	// Output:
	// [1 4 7 2 3]
	// [4 7 9 10]
}

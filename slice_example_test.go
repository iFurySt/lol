/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/7/19
 */

package lol

import (
	"fmt"
	"sort"
)

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

func ExampleMap() {
	res := Map([]int{1, 3, 5}, func(i int) int { return i * 2 })
	fmt.Println(res)
	// Output:
	// [2 6 10]
}

func ExampleReduce() {
	res1 := Reduce([]int{1, 7, 3, 4}, func(s float64, v, i int) float64 {
		return s * float64(v)
	}, 2)
	res2 := Reduce([]string{"1", "7", "3", "4"}, func(s, v string, i int) string {
		return fmt.Sprintf("%s %d.%s", s, i, v)
	}, "start:")
	fmt.Println(res1)
	fmt.Println(res2)
	// Output:
	// 168
	// start: 0.1 1.7 2.3 3.4
}

func ExampleReduceRight() {
	res1 := ReduceRight([]int{1, 7, 3, 4}, func(s float64, v, i int) float64 {
		return s * float64(v)
	}, 2)
	res2 := ReduceRight([]string{"1", "7", "3", "4"}, func(s, v string, i int) string {
		return fmt.Sprintf("%s %d.%s", s, i, v)
	}, "reverse:")
	fmt.Println(res1)
	fmt.Println(res2)
	// Output:
	// 168
	// reverse: 3.4 2.3 1.7 0.1
}

func ExampleKeys() {
	res1 := Keys(map[int]struct{}{
		1: {},
		7: {},
		3: {},
		4: {},
	})
	sort.Ints(res1)
	res2 := Keys(map[string]struct{}{
		"a": {},
		"b": {},
		"1": {},
		"2": {},
	})
	sort.Strings(res2)
	fmt.Println(res1)
	fmt.Println(res2)
	// Output:
	// [1 3 4 7]
	// [1 2 a b]
}

func ExampleValues() {
	res1 := Values(map[int]int{
		1: 3,
		7: 5,
		3: 9,
		4: 7,
	})
	sort.Ints(res1)
	res2 := Values(map[int]string{
		1: "3",
		7: "5",
		3: "9",
		4: "7",
	})
	sort.Strings(res2)
	fmt.Println(res1)
	fmt.Println(res2)
	// Output:
	// [3 5 7 9]
	// [3 5 7 9]
}

/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/7/19
 */

package lol

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testMultipleSlice struct {
	name   string
	inputs [][]int
	want   []int
}

func TestMergeSlice(t *testing.T) {
	cases := []testMultipleSlice{
		{
			"nil",
			[][]int{nil},
			[]int{},
		},
		{
			"empty",
			[][]int{},
			[]int{},
		},
		{
			"empties",
			[][]int{{}, {}},
			[]int{},
		},
		{
			"same length",
			[][]int{{1, 3, 5}, {2, 4, 6}},
			[]int{1, 3, 5, 2, 4, 6},
		},
		{
			"different length",
			[][]int{{1, 3, 5, 6}, {2, 4}},
			[]int{1, 3, 5, 6, 2, 4},
		},
		{
			"many slices",
			[][]int{{1}, {2, 5, 8}, {3, 6}, {4, 7, 9, 10}},
			[]int{1, 2, 5, 8, 3, 6, 4, 7, 9, 10},
		},
	}
	for _, c := range cases {
		assert.Equalf(t, c.want, MergeSlice(c.inputs...), c.name)
	}
}

func TestUniqSlice(t *testing.T) {
	cases := []testMultipleSlice{
		{
			"nil",
			[][]int{nil},
			[]int{},
		},
		{
			"empty",
			[][]int{},
			[]int{},
		},
		{
			"empties",
			[][]int{{}, {}},
			[]int{},
		},
		{
			"no duplications",
			[][]int{{1, 3, 5}, {2, 4, 6}},
			[]int{1, 3, 5, 2, 4, 6},
		},
		{
			"duplications-1",
			[][]int{{1, 3, 5, 6}, {3, 5}},
			[]int{1, 3, 5, 6},
		},
		{
			"duplications-2",
			[][]int{{1}, {1, 5, 8}, {1, 6}, {1, 7, 5, 10}},
			[]int{1, 5, 8, 6, 7, 10},
		},
	}
	for _, c := range cases {
		assert.Equalf(t, c.want, UniqSlice(c.inputs...), c.name)
	}
}

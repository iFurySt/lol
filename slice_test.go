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

func TestMergeSlice(t *testing.T) {
	cases := []struct {
		name   string
		inputs [][]int
		want   []int
	}{
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

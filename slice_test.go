/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/7/19
 */

package lol

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
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

func TestMap(t *testing.T) {
	for _, c := range []struct {
		name   string
		inputs []int
		fn     func(int) int
		want   []int
	}{
		{
			"nil",
			nil,
			func(v int) int {
				return -v
			},
			[]int{},
		},
		{
			"empty",
			[]int{},
			func(v int) int {
				return -v
			},
			[]int{},
		},
		{
			"normal",
			[]int{1, 7, 3, 4},
			func(v int) int {
				return -v
			},
			[]int{-1, -7, -3, -4},
		},
		{
			"fn nil",
			[]int{1, 7, 3, 4},
			nil,
			[]int{1, 7, 3, 4},
		},
	} {
		assert.Equalf(t, c.want, Map(c.inputs, c.fn), c.name)
	}
}

func TestReduce(t *testing.T) {
	t.Parallel()

	for _, c := range []struct {
		name    string
		inputs  []int
		fn      func(int, int, int) int
		initial int
		want    int
	}{
		{
			"nil",
			nil,
			func(s, v, i int) int {
				return s + v
			},
			0,
			0,
		},
		{
			"empty",
			[]int{},
			func(s, v, i int) int {
				return s + v
			},
			1,
			1,
		},
		{
			"normal",
			[]int{1, 7, 3, 4},
			func(s, v, i int) int {
				return s + v
			},
			0,
			15,
		},
		{
			"fn nil",
			[]int{1, 7, 3, 4},
			nil,
			0,
			0,
		},
	} {
		assert.Equalf(t, c.want, Reduce(c.inputs, c.fn, c.initial), c.name)
	}

	for _, c := range []struct {
		name    string
		inputs  []int
		fn      func(float64, int, int) float64
		initial float64
		want    float64
	}{
		{
			"nil",
			nil,
			func(s float64, v, i int) float64 {
				return s * float64(v)
			},
			0,
			0,
		},
		{
			"empty",
			[]int{},
			func(s float64, v, i int) float64 {
				return s * float64(v)
			},
			1,
			1,
		},
		{
			"normal",
			[]int{1, 7, 3, 4},
			func(s float64, v, i int) float64 {
				return s * float64(v)
			},
			2,
			168,
		},
		{
			"fn nil",
			[]int{1, 7, 3, 4},
			nil,
			3,
			3,
		},
	} {
		assert.Equalf(t, c.want, Reduce(c.inputs, c.fn, c.initial), c.name)
	}

	for _, c := range []struct {
		name    string
		inputs  []string
		fn      func(string, string, int) string
		initial string
		want    string
	}{
		{
			"string",
			[]string{"1", "7", "3", "4"},
			func(s, v string, i int) string {
				return fmt.Sprintf("%s %d.%s", s, i, v)
			},
			"start:", "start: 0.1 1.7 2.3 3.4",
		},
	} {
		assert.Equalf(t, c.want, Reduce(c.inputs, c.fn, c.initial), c.name)
	}
}

func TestReduceRight(t *testing.T) {
	t.Parallel()

	for _, c := range []struct {
		name    string
		inputs  []int
		fn      func(int, int, int) int
		initial int
		want    int
	}{
		{
			"nil",
			nil,
			func(s, v, i int) int {
				return s + v
			},
			0,
			0,
		},
		{
			"empty",
			[]int{},
			func(s, v, i int) int {
				return s + v
			},
			1,
			1,
		},
		{
			"normal",
			[]int{1, 7, 3, 4},
			func(s, v, i int) int {
				return s + v
			},
			0,
			15,
		},
		{
			"fn nil",
			[]int{1, 7, 3, 4},
			nil,
			0,
			0,
		},
	} {
		assert.Equalf(t, c.want, ReduceRight(c.inputs, c.fn, c.initial), c.name)
	}

	for _, c := range []struct {
		name    string
		inputs  []int
		fn      func(float64, int, int) float64
		initial float64
		want    float64
	}{
		{
			"nil",
			nil,
			func(s float64, v, i int) float64 {
				return s * float64(v)
			},
			0,
			0,
		},
		{
			"empty",
			[]int{},
			func(s float64, v, i int) float64 {
				return s * float64(v)
			},
			1,
			1,
		},
		{
			"normal",
			[]int{1, 7, 3, 4},
			func(s float64, v, i int) float64 {
				return s * float64(v)
			},
			2,
			168,
		},
		{
			"fn nil",
			[]int{1, 7, 3, 4},
			nil,
			3,
			3,
		},
	} {
		assert.Equalf(t, c.want, ReduceRight(c.inputs, c.fn, c.initial), c.name)
	}

	for _, c := range []struct {
		name    string
		inputs  []string
		fn      func(string, string, int) string
		initial string
		want    string
	}{
		{
			"string",
			[]string{"1", "7", "3", "4"},
			func(s, v string, i int) string {
				return fmt.Sprintf("%s %d.%s", s, i, v)
			},
			"reverse:", "reverse: 3.4 2.3 1.7 0.1",
		},
	} {
		assert.Equalf(t, c.want, ReduceRight(c.inputs, c.fn, c.initial), c.name)
	}
}

func TestKeys(t *testing.T) {
	for _, c := range []struct {
		name   string
		inputs map[string]int
		want   []string
	}{
		{
			"nil",
			nil,
			[]string{},
		},
		{
			"empty",
			map[string]int{},
			[]string{},
		},
		{
			"string",
			map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			[]string{"one", "three", "two"},
		},
	} {
		res := Keys(c.inputs)
		sort.Strings(res)
		assert.Equalf(t, c.want, res, c.name)
	}
}

func TestValues(t *testing.T) {
	for _, c := range []struct {
		name   string
		inputs map[string]int
		want   []int
	}{
		{
			"nil",
			nil,
			[]int{},
		},
		{
			"empty",
			map[string]int{},
			[]int{},
		},
		{
			"string",
			map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
			},
			[]int{1, 2, 3},
		},
	} {
		res := Values(c.inputs)
		sort.Ints(res)
		assert.Equalf(t, c.want, res, c.name)
	}
}

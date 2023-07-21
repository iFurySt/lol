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
		{"nil", [][]int{nil}, []int{}},
		{"empty", [][]int{}, []int{}},
		{"empties", [][]int{{}, {}}, []int{}},
		{"same length", [][]int{{1, 3, 5}, {2, 4, 6}}, []int{1, 3, 5, 2, 4, 6}},
		{"different length", [][]int{{1, 3, 5, 6}, {2, 4}}, []int{1, 3, 5, 6, 2, 4}},
		{"many slices", [][]int{{1}, {2, 5, 8}, {3, 6}, {4, 7, 9, 10}},
			[]int{1, 2, 5, 8, 3, 6, 4, 7, 9, 10}},
	}
	for _, c := range cases {
		assert.Equalf(t, c.want, MergeSlice(c.inputs...), c.name)
	}
}

func TestUniqSlice(t *testing.T) {
	cases := []testMultipleSlice{
		{"nil", [][]int{nil}, []int{}},
		{"empty", [][]int{}, []int{}},
		{"empties", [][]int{{}, {}}, []int{}},
		{"no duplications", [][]int{{1, 3, 5}, {2, 4, 6}}, []int{1, 3, 5, 2, 4, 6}},
		{"duplications-1", [][]int{{1, 3, 5, 6}, {3, 5}}, []int{1, 3, 5, 6}},
		{"duplications-2", [][]int{{1}, {1, 5, 8}, {1, 6}, {1, 7, 5, 10}},
			[]int{1, 5, 8, 6, 7, 10}},
	}
	for _, c := range cases {
		assert.Equalf(t, c.want, UniqSlice(c.inputs...), c.name)
		assert.Equalf(t, c.want, Union(c.inputs...), c.name)
	}
}

func TestIntersection(t *testing.T) {
	cases := []testMultipleSlice{
		{"nil", [][]int{nil}, []int{}},
		{"empty", [][]int{}, []int{}},
		{"empties", [][]int{{}, {}}, []int{}},
		{"not-intersect-1", [][]int{{1, 3, 5}, {2, 4, 6}}, []int{}},
		{"not-intersect-2", [][]int{{1, 3, 5}, {2, 4, 6}, {7, 8}}, []int{}},
		{"normal-case1", [][]int{{1, 3, 5, 6}, {3, 5}}, []int{3, 5}},
		{"normal-case2", [][]int{{1}, {1, 5, 8}, {1, 6}, {1, 7, 5, 10}}, []int{1}},
	}
	for _, c := range cases {
		assert.Equalf(t, c.want, Intersection(c.inputs...), c.name)
	}
}

func TestDifference(t *testing.T) {
	cases := []struct {
		name       string
		a, b, want []int
	}{
		{"nil", nil, nil, []int{}},
		{"empty", []int{}, []int{}, []int{}},
		{"a-empty", []int{}, []int{2, 4, 6}, []int{}},
		{"b-empty", []int{3, 5, 1}, []int{}, []int{1, 3, 5}},
		{"normal-case1", []int{3, 5, 1}, []int{2, 4, 6}, []int{1, 3, 5}},
		{"normal-case2", []int{3}, []int{2, 3, 6}, []int{}},
	}
	for _, c := range cases {
		res := Difference(c.a, c.b)
		sort.Ints(res)
		assert.Equalf(t, c.want, res, c.name)
	}
}

func TestMap(t *testing.T) {
	for _, c := range []struct {
		name   string
		inputs []int
		fn     func(int) int
		want   []int
	}{
		{"nil", nil, func(v int) int { return -v }, []int{}},
		{"empty", []int{}, func(v int) int { return -v }, []int{}},
		{"normal", []int{1, 7, 3, 4}, func(v int) int { return -v }, []int{-1, -7, -3, -4}},
		{"fn nil", []int{1, 7, 3, 4}, nil, []int{1, 7, 3, 4}},
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
		{"nil", nil, func(s, v, i int) int { return s + v }, 0, 0},
		{"empty", []int{}, func(s, v, i int) int { return s + v }, 1, 1},
		{"normal", []int{1, 7, 3, 4}, func(s, v, i int) int { return s + v }, 0, 15},
		{"fn nil", []int{1, 7, 3, 4}, nil, 0, 0},
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
		{"nil", nil, func(s float64, v, i int) float64 { return s * float64(v) }, 0, 0},
		{"empty", []int{}, func(s float64, v, i int) float64 { return s * float64(v) },
			1, 1},
		{"normal", []int{1, 7, 3, 4}, func(s float64, v, i int) float64 { return s * float64(v) },
			2, 168},
		{"fn nil", []int{1, 7, 3, 4}, nil, 3, 3},
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
		{"string", []string{"1", "7", "3", "4"},
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

/**
 * Package lol
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/7/20
 */

package lol

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	// int
	for _, c := range []struct {
		name   string
		number int64
		want   int64
	}{
		{"0", 0, 0},
		{"-0", -0, 0},
		{"3", 3, 3},
		{"-3", -3, 3},
		{"Inf-8", math.MinInt8, -math.MinInt8},
		{"Inf-32", math.MinInt32, -math.MinInt32},
	} {
		assert.Equalf(t, c.want, Abs(c.number), c.name)
	}

	// float64
	for _, c := range []struct {
		name   string
		number float64
		want   float64
	}{
		{"3", float64(3), float64(3)},
		{"-3", float64(-3), float64(3)},
		{"Inf", math.Inf(1), math.Inf(1)},
		{"-Inf", math.Inf(-1), math.Inf(1)},
	} {
		assert.Equalf(t, c.want, Abs(c.number), c.name)
	}
}

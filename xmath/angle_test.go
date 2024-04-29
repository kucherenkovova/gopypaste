package xmath_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kucherenkovova/gopypaste/xmath"
)

func TestRadians(t *testing.T) {
	tests := []struct {
		deg float64
		rad float64
	}{
		{0, 0},
		{15, math.Pi / 12},
		{30, math.Pi / 6},
		{45, math.Pi / 4},
		{90, math.Pi / 2},
		{180, math.Pi},
		{270, 3 * math.Pi / 2},
		{360, 2 * math.Pi},
		{-90, -math.Pi / 2},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.deg)), func(t *testing.T) {
			got := xmath.Radians(tt.deg)
			assert.InDelta(t, tt.rad, got, 0.0000000001)
		})
	}
}

func TestDegrees(t *testing.T) {
	tests := []struct {
		deg float64
		rad float64
	}{
		{0, 0},
		{15, math.Pi / 12},
		{30, math.Pi / 6},
		{45, math.Pi / 4},
		{90, math.Pi / 2},
		{180, math.Pi},
		{270, 3 * math.Pi / 2},
		{360, 2 * math.Pi},
		{-90, -math.Pi / 2},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.deg)), func(t *testing.T) {
			got := xmath.Degrees(tt.rad)
			assert.InDelta(t, tt.deg, got, 0.0000000001)
		})
	}

}

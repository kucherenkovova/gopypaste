package xmath

import (
	"math"
)

func Radians(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func Degrees(rad float64) float64 {
	return rad * (180 / math.Pi)
}

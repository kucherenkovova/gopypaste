package xmaps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kucherenkovova/gopypaste/xmaps"
)

func TestStringKeys(t *testing.T) {
	input := map[string]int{
		"one":   100,
		"two":   200,
		"three": 300,
	}

	expected := []string{"one", "two", "three"}
	results := xmaps.Keys(input)

	assert.ElementsMatch(t, expected, results)
}

func TestIntKeys(t *testing.T) {
	input := map[int]string{
		100: "one",
		200: "two",
		300: "three",
	}

	expected := []int{100, 200, 300}
	results := xmaps.Keys(input)

	assert.ElementsMatch(t, expected, results)
}

func TestKeysEmptyInput(t *testing.T) {
	assert.Nil(t, xmaps.Keys(map[int]string{}))
	var nilMap map[int]string
	assert.Nil(t, xmaps.Keys(nilMap))
}

func TestStringValues(t *testing.T) {
	input := map[string]int{
		"one":   100,
		"two":   200,
		"three": 300,
	}

	expected := []int{100, 200, 300}
	results := xmaps.Values(input)

	assert.ElementsMatch(t, expected, results)
}

func TestFloat64Values(t *testing.T) {
	input := map[string]float64{
		"one":   111.1,
		"two":   222.2,
		"three": 333.3,
	}

	expected := []float64{111.1, 222.2, 333.3}
	results := xmaps.Values(input)

	assert.ElementsMatch(t, expected, results)
}

func TestValuesEmptyInput(t *testing.T) {
	assert.Nil(t, xmaps.Values(map[string]float64{}))
	var nilMap map[string]float64
	assert.Nil(t, xmaps.Values(nilMap))
}

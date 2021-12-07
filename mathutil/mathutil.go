// Package mathutil provides boilerplate mathematical operations not supported in the native math package
package mathutil

import (
	"errors"
	"sort"
)

var (
	ErrEmptySlice = errors.New("mathutil: slice parameter is empty")
)

// Mean calculates the mean of the given int slice and returns it
func Mean(in []int) (float64, error) {
	if len(in) == 0 {
		return -1, ErrEmptySlice
	}

	sum := 0

	for _, el := range in {
		sum += el
	}

	return float64(sum) / float64(len(in)), nil
}

// Median calculates the median of the given int slice and returns it
func Median(in []int) (float64, error) {
	if len(in) == 0 {
		return -1, ErrEmptySlice
	}

	sort.Ints(in)

	cen := len(in) / 2

	var med float64
	if len(in)%2 == 0 {
		// even num of ints, so take average of the center
		med = (float64(in[cen-1]) + float64(in[cen])) / 2.0
	} else {
		med = float64(in[cen])
	}

	return med, nil
}

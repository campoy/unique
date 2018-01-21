// Package unique provides primitives for sorting slices removing repeated elements.
package unique

import (
	"reflect"
	"sort"
)

// Slice sorts the slice pointed by the provided pointer given the provided
// less function and removes repeated elements.
// The function panics if the provided interface is not a pointer to a slice.
func Slice(slicePtr interface{}, less func(i, j int) bool) {
	v := reflect.ValueOf(slicePtr).Elem()
	if v.Len() == 0 {
		return
	}
	sort.Slice(v.Interface(), less)

	l := 1
	for i := 1; i < v.Len(); i++ {
		if !less(i, l-1) && !less(l-1, i) {
			continue
		}
		v.Index(l).Set(v.Index(i))
		l++
	}
	v.SetLen(l)
}

package dutch_flag

import (
	"reflect"
	"testing"
)

func testBase(t *testing.T, orig, expected []int) {
	sort(orig)
	if !areSliceEqual(orig, expected) {
		t.Fatal("error")
	}
	if !areSliceEqualAlt(orig, expected) {
		t.Fatal("error")
	}
}

func TestBasicSort(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 1, 2, 0, 0, 0, 1}
	sorted_array := []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2}
	testBase(t, arr, sorted_array)
}

func TestAllZeors(t *testing.T) {
	arr := []int{0, 0, 0}
	sorted_array := []int{0, 0, 0}
	testBase(t, arr, sorted_array)
}

func areSliceEqualAlt(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}

func areSliceEqual(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

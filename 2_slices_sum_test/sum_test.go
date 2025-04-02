package slicessumtest

import (
	"reflect"
	"testing"
)



func TestSlicesSum(t *testing.T){
	t.Run("tesing sum on single array", func(t *testing.T){
		nums:= []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		want := 15

		if got!=want {
			t.Errorf("got %d, want %d", got, want)
		}

	})

	t.Run("checkign for multiple arrays at once", func(t *testing.T){
		nums1 := []int{1, 2, 3, 4, 5}
		nums2 := []int{5, 6}

		got := SumAll(nums1, nums2)
		want := []int{15, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})
	t.Run("checkign for multiple arrays at once", func(t *testing.T){
		nums1 := []int{1, 2, 3, 4, 5}
		nums2 := []int{5, 6}

		got := SumAllUsingAppend(nums1, nums2)
		want := []int{15, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}

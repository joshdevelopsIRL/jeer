package jeer

import (
	"errors"
	"testing"
)

func TestJeer(t *testing.T) {
    Test[int](t).Actual(12).Expected(13).Run("should fail basic")

    var e error
    var err error 
    Test[error](t).Actual(e).Expected(err).Run("test errors empty")

    p := errors.New("This is Error A")
    l := errors.New("This is Error BB")
    Test[error](t).Actual(p).Expected(l).Run("test errors should fail")

    aList := []int{1,3,5,7,9}
    bList := []int{1,5,3,7,9}
    Test[int](t).IsList().Actual(bList...).Expected(aList...).Run("test int slice should fail")

    actual := []int{1,2,3}
    expected := []int{1,2,3}
    Test[int](t).IsList().Actual(actual...).Expected(expected...).Run("test int slice")
}

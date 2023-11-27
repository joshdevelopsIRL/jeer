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

    aList = []int{1,3,5,7,9}
    bList = []int{1,5,7,3,9}
    Test[int](t).IsList().AnyOrder().Actual(bList...).Expected(aList...).Run("test int slice should work (anyorder)")

    aList = []int{1,3,5,7,9}
    bList = []int{1,5,7,9}
    Test[int](t).IsList().AnyOrder().Actual(bList...).Expected(aList...).Run("test int slice should fail (anyorder)")

    actual := []int{1,2,3}
    expected := []int{1,2,3}
    Test[int](t).IsList().Actual(actual...).Expected(expected...).Run("test int slice")

    actual = []int{1,2,3}
    expected = []int{1,2,3}
    err = errors.New("this thing failed or something")
    Test[int](t).IsList().Actual(actual...).Expected(expected...).FailOn(err).Run("test int slice should fail, added FailOn(err)")

    actual = []int{1,2,3}
    expected = []int{3,1,2}
    err = nil
    Test[int](t).IsList().AnyOrder().Actual(actual...).Expected(expected...).FailOn(err).Run("test int slice should work, added FailOn(nil)")
    
    Test[string](t).Actual().Expected().Run("test fail both inputs single")
    Test[string](t).IsList().Actual().Expected().Run("test both inputs as empty lists should work")
}

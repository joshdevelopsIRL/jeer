package jeer

import "testing"

type Jeer[T comparable] struct {
    tester *testing.T
    actual T
    expected T
    name string
    err error
    isList bool
    actualList []T
    expectedList []T
}

func Test[T comparable] (t *testing.T) *Jeer[T] {
    return &Jeer[T]{
        tester: t,
        err: nil,
        isList: false,
    }
}

func (j *Jeer[T]) IsList() *Jeer[T] {
    j.isList = true
    return j
}

func (j *Jeer[T]) FailOn(err error) *Jeer[T] {
    j.err = err
    return j
}

func (j *Jeer[T]) Actual(v ...T) *Jeer[T] {
    if j.isList {
        j.actualList = v
    } else {
        if len(v) == 1 {
            j.actual = v[0]
        }
    }
    return j
}

func (j *Jeer[T]) Expected(v ...T) *Jeer[T] {
    if j.isList {
        j.expectedList = v
    } else {
        if len(v) == 1 {
            j.expected = v[0]
        }
    }
    return j
}

func (j *Jeer[T]) compareLists(a, b []T) bool {
    if len(a) == len(b) {
        for i := range a {
            if a[i] != b[i] {
                return false
            }
        }
        return true
    }
    return false
}

func (j *Jeer[T]) Run(name string) *Jeer[T] {
    j.name = name
    j.tester.Run(name, func(t *testing.T) {
        if j.err != nil {
            t.Fatalf("error '%v' | wanted '%v'", j.err, j.expected)
        }
        if j.isList {
            if !j.compareLists(j.actualList, j.expectedList) {
                t.Fatalf("got '%v' | wanted '%v'", j.actualList, j.expectedList)
            }
        } 
        if !j.isList {
            if j.actual != j.expected {
                t.Fatalf("got '%v' | wanted '%v'", j.actual, j.expected)
            }
        }
    })
    return j
}

package gosolist

import (
	"fmt"
	"testing"
)

func BenchmarkListGet(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}

func TestList(t *testing.T) {
	l := SortedList{c: func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}}
	l.Insert(1)
	l.Insert(34)
	l.Insert(23)
	l.Insert(5)
	l.Insert(534)
	l.Insert(423)
	fmt.Println(l.Values())
	fmt.Println(l.DeleteItem(34))
	fmt.Println(l.Values())
}

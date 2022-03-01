package gosolist

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestList(t *testing.T) {
	l := NewSortedList(IntsCompare, 50)
	for i := 0; i < 3000; i++ {
		l.Push(rand.Int() % 10000)
		if i%100 == 0 {
			fmt.Println(l.At(i))
		}
	}
	fmt.Println(l.Values())
}

package gosolist

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := NewSortedList(IntsCompare, 3)
	l.Push(1)
	l.Push(34)
	l.Push(23)
	l.Push(5)
	l.Push(534)
	l.Push(423)
	fmt.Println(l.Values())
	l.Push(76)
	l.Push(987)

	fmt.Println(l.At(6))
	l.Push(76)
	l.Push(987)
	l.Push(987)
	for i := 0; i < l.Size(); i++ {
		fmt.Println(l.At(i))
	}
}

//       11
//   9      2
// 4   5    2 0
//2 2 3 2 2 2 2
// 11 9 2 4 5 2 2 2 3 2 2

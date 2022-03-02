package gosolist

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func TestList(t *testing.T) {
	l := NewSortedList(IntCompare, 2000)
	most := int64(0)
	index := 0
	for i := 1; i < 60000; i++ {
		start := time.Now().UnixNano()
		a := rand.Int() % 10000000
		if !l.Has(a) {
			l.Push(a)
		}
		if rand.Int()%100 == 0 {
			temp := l.At(rand.Int() % l.Size())
			if !l.Has(temp) {
				t.Log("function 'has' is not correct")
			}
			if rand.Int()%10 == 9 {
				if !l.DeleteItem(temp) {
					t.Log("function deleteItem is not correct")
				}
				if l.Has(temp) {
					t.Log("function has is not correct")
				}
			}
		}
		if rand.Int()%200 == 0 {
			l.Delete(rand.Int() % l.Size())
		}
		end := time.Now().UnixNano()
		if end-start > most {
			most = end - start
			index = i
		}
	}
	t.Log(l.At(30000), index, most)
}

func TestFloor(t *testing.T) {
	l := NewSortedList(IntCompare, 3)
	var a []int
	for i := 0; i < 10; i++ {
		r := rand.Int() % 10000
		a = append(a, r)
		l.Push(r)
	}
	sort.Ints(a)
	if a[3] != l.At(3) {
		t.Fail()
	}
	fmt.Println(a, l.Values())
	fmt.Println(l.Floor(a[0]), l.Floor(a[0]-1), l.Ceil(a[0]-1))
	fmt.Println(l.Ceil(a[9]), l.Ceil(a[9]+1), l.Floor(a[9]-1))
}

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
				t.Fail()
			}
			if rand.Int()%10 == 9 {
				if !l.DeleteItem(temp) {
					t.Log("function deleteItem is not correct")
					t.Fail()
				}
				if l.Has(temp) {
					t.Log("function has is not correct")
					t.Fail()
				}
			}
		}
		end := time.Now().UnixNano()
		if end-start > most {
			most = end - start
			index = i
		}
	}
	t.Log(l.Size(), index, most)
}

func TestFloor(t *testing.T) {
	l := NewSortedList(IntCompare, 2)
	var a []int
	for i := 0; i < 20; i++ {
		r := rand.Int() % 10000
		a = append(a, r)
		l.Push(r)
	}
	sort.Ints(a)
	if a[3] != l.At(3) {
		t.Fail()
	}
	fmt.Println(a, l.Values())
	fmt.Println(l.Floor(a[1] + 1))
	fmt.Println(l.Floor(a[0]), l.Floor(a[0]-1), l.Ceil(a[0]-1))
	fmt.Println(l.Ceil(a[9]), l.Ceil(a[9]+1), l.Floor(a[9]+1))
}

func TestIndex(t *testing.T) {
	l := NewSortedList(IntCompare, 3)
	var a []int
	for i := 0; i < 50; i++ {
		r := rand.Int() % 10000
		a = append(a, r)
		l.Push(r)
	}
	sort.Ints(a)
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(l.Index(a[i] + 1))
		fmt.Println(l.Index(a[i]))
	}
}

func TestDelete(t *testing.T) {
	l := NewSortedList(IntCompare, 100)
	l.Push(1)
	l.Push(5)
	l.Push(9)
	fmt.Println(l.DeleteItem(9))
	l.Each(PrintEach)
}

func BenchmarkPush(b *testing.B) {
	l := NewSortedList(IntCompare, 2000)
	for n := 0; n < b.N; n++ {
		l.Push(n + n/2)
	}

}

func BenchmarkIndex(b *testing.B) {
	l := NewSortedList(IntCompare, 2000)
	for n := 0; n < 1000000; n++ {
		l.Push(rand.Int() % 100000)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		l.Index(n)
	}
}

func BenchmarkDelete(b *testing.B) {
	l := NewSortedList(IntCompare, 2000)
	for n := 0; n < 1000000; n++ {
		l.Push(rand.Int() % 100000)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		if l.Size() > 0 {
			l.DeleteItem(n)
		}
	}
}

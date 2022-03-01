package gosolist

import (
	"testing"
	"time"
)

func TestList(t *testing.T) {
	l := NewSortedList(IntsCompare, 2000)
	most := int64(0)
	index := 0
	for i := 10000; i < 60000; i++ {
		start := time.Now().UnixNano()
		l.Push(i / 10000)
		end := time.Now().UnixNano()
		if start-end > most {
			most = start - end
			index = i
		}
	}
	t.Log(l.At(30000), index, most)
}

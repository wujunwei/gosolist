package gosolist

import (
	"fmt"
	"sort"
)

type Compare func(a, b interface{}) bool

var IntsCompare Compare = func(a, b interface{}) bool {
	return a.(int) < b.(int)
}

var StringsCompare Compare = func(a, b interface{}) bool {
	return a.(string) < b.(string)
}

type ForEach func(index int, a interface{})

//PrintEach for debug usage
var PrintEach ForEach = func(index int, a interface{}) {
	fmt.Printf("index: %d, value: %v \n", index, a)
}

// Container is base interface that all data structures implement.
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}

func BisectRight(l []interface{}, c Compare, target interface{}) int {
	return sort.Search(len(l), func(i int) bool {
		return l[i] == target || !c(l[i], target)
	})
}

func InSort(l []interface{}, c Compare, a interface{}) []interface{} {
	index := BisectRight(l, c, a)
	if index == len(l) {
		return append(l, a)
	}
	return append(l[:index], append([]interface{}{a}, l[index:]...)...)
}
func RemoveSort(l []interface{}, c Compare, a interface{}) ([]interface{}, bool) {
	index := BisectRight(l, c, a)
	remove := l[index] == a
	if l[index] == a {
		l = append(l[:index], l[index+1:]...)
	}
	return l, remove
}

package gosolist

import (
	"fmt"
	"sort"
)

type Compare func(a, b interface{}) bool

var IntCompare Compare = func(a, b interface{}) bool {
	return a.(int) < b.(int)
}

var StringCompare Compare = func(a, b interface{}) bool {
	return a.(string) < b.(string)
}

type ForEach func(index int, a interface{})

//PrintEach for debug usage
var PrintEach ForEach = func(index int, a interface{}) {
	fmt.Printf("index: %d, value: %v \n", index, a)
}

func BisectRight(l []interface{}, c Compare, target interface{}) int {
	return sort.Search(len(l), func(i int) bool {
		return !c(l[i], target)
	})
}
func BisectLeft(l []interface{}, c Compare, target interface{}) int {
	return sort.Search(len(l), func(i int) bool {
		return l[i] == target || !c(l[i], target)
	})
}

func InSort(l []interface{}, c Compare, a interface{}) []interface{} {
	index := BisectRight(l, c, a)
	if index == len(l) {
		return append(l, a)
	}
	l = append(l, nil)
	copy(l[index+1:], l[index:len(l)-1])
	l[index] = a
	return l
}

func RemoveSort(l []interface{}, c Compare, a interface{}) ([]interface{}, bool) {
	index := BisectRight(l, c, a)
	remove := l[index] == a
	if l[index] == a {
		l = append(l[:index], l[index+1:]...)
	}
	return l, remove
}

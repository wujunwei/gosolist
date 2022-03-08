package gosolist

import (
	"fmt"
	"sort"
)

type Compare func(a, b interface{}) int

var IntCompare Compare = func(a, b interface{}) int {
	aa, bb := a.(int), b.(int)
	if aa > bb {
		return 1
	}
	if aa < bb {
		return -1
	}
	return 0
}

var StringCompare Compare = func(a, b interface{}) int {
	aa, bb := a.(string), b.(string)
	if aa > bb {
		return 1
	}
	if aa < bb {
		return -1
	}
	return 0
}

type ForEach func(index int, a interface{})

//PrintEach for debug usage
var PrintEach ForEach = func(index int, a interface{}) {
	fmt.Printf("index: %d, value: %v \n", index, a)
}

func BisectRight(l []interface{}, c Compare, target interface{}) int {
	return sort.Search(len(l), func(i int) bool {
		return c(l[i], target) > 0
	})
}
func BisectLeft(l []interface{}, c Compare, target interface{}) int {
	return sort.Search(len(l), func(i int) bool {
		return c(l[i], target) >= 0
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
	if index == len(l) || index == 0 {
		return l, false
	}
	remove := l[index-1] == a
	if remove {
		l = append(l[:index-1], l[index:]...)
	}
	return l, remove
}

package main

import (
	"fmt"
	"github.com/wujunwei/gosolist"
)

func main() {
	l := gosolist.NewSortedList(gosolist.IntCompare, gosolist.DefaultLoadFactor)
	l.Push(123)
	l.Push(46)
	l.Push(7)
	fmt.Println(l.At(2)) // 123
	l.Push(7)
	fmt.Println(l.At(2)) // 46
	l.Push(75)
	l.Push(45)
	fmt.Println(l.Values()) // 7 7 45 46 75 123
	l.Delete(2)             //delete the item(45) at 2

	l.Each(gosolist.PrintEach) // equal fmt.Println(l.Values()) 7 7 46 75 123
	l.Push(46)
	l.DeleteItem(46) //7 7 46 75 123

	fmt.Println(l.Floor(8), l.Ceil(8))   // 7  46
	fmt.Println(l.Floor(46), l.Ceil(46)) // 46 46
	fmt.Println(l.Floor(6), l.Ceil(124)) // nil nil

}

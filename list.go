package gosolist

const (
	DefaultLoadFactor = 1000
)

type SortedList struct {
	maxes   []interface{}
	lists   [][]interface{}
	indexes []int //index sum tree
	size    int
	c       Compare
}

func (l *SortedList) Insert(a interface{}) {
	l.size++
	if len(l.maxes) == 0 {
		l.maxes = append(l.maxes, a)
		l.lists = append(l.lists, []interface{}{a})
		return
	}
	pos := BisectRight(l.maxes, l.c, a)
	if pos == len(l.maxes) {
		pos--
		l.maxes[pos] = a
		l.lists[pos] = append(l.lists[pos], a)
	} else {
		l.lists[pos] = InSort(l.lists[pos], l.c, a)
	}
	l.extend()
}

func (l *SortedList) DeleteItem(a interface{}) bool {
	if l.size == 0 {
		return false
	}
	l.size--
	pos := BisectRight(l.maxes, l.c, a)
	if pos == len(l.maxes) {
		return false
	}
	var removed bool
	l.lists[pos], removed = RemoveSort(l.lists[pos], l.c, a)
	l.extend()
	return removed
}

func (l *SortedList) Delete(index int) {

}

func (l *SortedList) Values() []interface{} {
	res := make([]interface{}, l.Size())
	i := 0
	l.Each(func(_ int, a interface{}) {
		res[i] = a
		i++
	})
	return res
}

// todo
func (l *SortedList) At(index int) interface{} {
	return nil
}

func (l *SortedList) Each(f ForEach) {
	i := 0
	for _, list := range l.lists {
		for _, j := range list {
			f(i, j)
		}
	}
}

func (l *SortedList) Has(a interface{}) bool {
	if l.size == 0 {
		return false
	}
	pos := BisectRight(l.maxes, l.c, a)
	if pos == len(l.maxes) {
		return false
	}
	index := BisectRight(l.lists[pos], l.c, a)
	return l.lists[pos][index] == a
}

func (l *SortedList) Clear() {
	l.reset()
}

func (l *SortedList) Empty() bool {
	return l.Size() == 0
}

func (l *SortedList) Size() int {
	return l.size
}

// todo build index tree
func (l *SortedList) extend() {

}

func (l *SortedList) reset() {
	l.lists = [][]interface{}{}
	l.indexes = []int{}
	l.maxes = []interface{}{}
	l.size = 0
}

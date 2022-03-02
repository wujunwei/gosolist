package gosolist

const (
	DefaultLoadFactor = 1000
)

type SortedList struct {
	offset  int
	load    int
	maxes   []interface{}
	lists   [][]interface{}
	indexes []int //index sum tree
	size    int
	c       Compare
}

func (l *SortedList) Push(a interface{}) {
	l.size++
	if len(l.maxes) == 0 {
		l.maxes = append(l.maxes, a)
		l.lists = append(l.lists, []interface{}{a})
		return
	}
	pos := BisectRight(l.maxes, l.c, a)
	if pos > 0 && l.maxes[pos-1] == a {
		pos--
	}
	if pos == len(l.maxes) {
		pos--
		l.maxes[pos] = a
		l.lists[pos] = append(l.lists[pos], a)
	} else {
		l.lists[pos] = InSort(l.lists[pos], l.c, a)
	}
	l.fresh(pos)
}

func (l *SortedList) DeleteItem(a interface{}) bool {
	if l.size == 0 {
		return false
	}
	l.size--
	pos := BisectLeft(l.maxes, l.c, a)
	if pos == len(l.maxes) {
		return false
	}
	var removed bool
	l.lists[pos], removed = RemoveSort(l.lists[pos], l.c, a)
	if len(l.lists[pos]) == 0 {
		l.maxes = append(l.maxes[:pos], l.maxes[pos+1:]...)
		l.lists = append(l.lists[:pos], l.lists[pos+1:]...)
		l.resetIndex()
	} else {
		l.maxes[pos] = l.lists[pos][len(l.lists[pos])-1]
		l.updateIndex(pos, -1)
	}
	return removed
}

func (l *SortedList) Delete(index int) {
	if index >= l.size {
		return
	}
	var pos, in int
	if index == 0 {
		pos, in = 0, 0
	} else if index == l.size-1 {
		pos = len(l.lists) - 1
		in = len(l.lists[pos]) - 1
	} else {
		if len(l.indexes) == 0 {
			l.buildIndex()
		}
		pos, in = l.findPos(index)
	}
	l.size--
	l.lists[pos] = append(l.lists[pos][:in], l.lists[pos][in+1:]...)
	if len(l.lists[pos]) == 0 {
		l.maxes = append(l.maxes[:pos], l.maxes[pos+1:]...)
		l.lists = append(l.lists[:pos], l.lists[pos+1:]...)
		l.resetIndex()
	} else {
		l.maxes[pos] = l.lists[pos][len(l.lists[pos])-1]
		l.updateIndex(pos, -1)
	}
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

func (l *SortedList) At(index int) interface{} {
	if index >= l.size {
		return nil
	}
	if index == 0 {
		return l.lists[0][0]
	}
	if index == l.size-1 {
		return l.maxes[len(l.maxes)-1]
	}
	//if l.size-index <= len(l.lists[len(l.lists)-1]) {
	//	return l.lists[len(l.lists)-1][l.size-index-1]
	//}
	if len(l.indexes) == 0 {
		l.buildIndex()
	}
	pos, in := l.findPos(index)
	return l.lists[pos][in]
}

func (l *SortedList) Each(f ForEach) {
	i := 0
	for _, list := range l.lists {
		for _, j := range list {
			f(i, j)
			i++
		}
	}
}

func (l *SortedList) Has(a interface{}) bool {
	if l.size == 0 {
		return false
	}
	pos := BisectLeft(l.maxes, l.c, a)
	if pos == len(l.maxes) {
		return false
	}
	index := BisectLeft(l.lists[pos], l.c, a)
	return l.lists[pos][index] == a
}

func (l *SortedList) Floor(a interface{}) interface{} {
	if l.size == 0 {
		return nil
	}
	pos := BisectLeft(l.maxes, l.c, a)
	if pos == len(l.maxes) {
		return l.maxes[pos-1]
	}
	index := BisectLeft(l.lists[pos], l.c, a)
	if index == 0 && l.lists[pos][0] != a {
		if pos == 0 {
			return nil
		} else {
			return l.maxes[pos-1]
		}
	}
	return l.lists[pos][index]
}

func (l *SortedList) Ceil(a interface{}) interface{} {
	if l.size == 0 {
		return nil
	}
	pos := BisectLeft(l.maxes, l.c, a)
	if pos == len(l.maxes) {
		return nil
	}
	index := BisectRight(l.lists[pos], l.c, a)
	if index == len(l.lists[pos]) {
		return nil
	}
	if index > 0 && l.lists[pos][index-1] == a {
		return a
	}
	return l.lists[pos][index]
}

//todo
//func (l *SortedList) Index(a interface{}) (int, bool) {
//
//}

func (l *SortedList) Empty() bool {
	return l.size == 0
}

func (l *SortedList) Size() int {
	return l.size
}

func (l *SortedList) Clear() {
	l.resetIndex()
	l.lists = [][]interface{}{}
	l.maxes = []interface{}{}
	l.size = 0
}

// fresh update the index and rebuild basic array if the load is greater than load factor after insert
func (l *SortedList) fresh(pos int) {
	listPosLen := len(l.lists[pos])
	if listPosLen > l.load {
		halfLen := listPosLen >> 1
		half := append([]interface{}{}, l.lists[pos][halfLen:]...)
		l.lists[pos] = l.lists[pos][:halfLen]
		l.lists = append(l.lists[:pos+1], append([][]interface{}{half}, l.lists[pos+1:]...)...)
		// update max
		l.maxes[pos] = l.lists[pos][halfLen-1]
		l.maxes = append(l.maxes[:pos+1], append([]interface{}{l.lists[pos+1][len(l.lists[pos+1])-1]}, l.maxes[pos+1:]...)...)
		l.resetIndex()
	} else {
		l.maxes[pos] = l.lists[pos][listPosLen-1]
		l.updateIndex(pos, 1)
	}
}

func (l *SortedList) buildIndex() {
	n := len(l.lists)
	rowLens := roundUpOf2((n + 1) / 2)
	l.offset = rowLens*2 - 1
	indexLens := l.offset + n

	indexes := make([]int, indexLens)
	for i, list := range l.lists { // fill row0
		indexes[len(indexes)-n+i] = len(list)
	}

	last := indexLens - n - rowLens
	for rowLens > 0 {
		for i := 0; i < rowLens; i++ {
			if (last+i)*2+1 >= indexLens {
				break
			}
			if (last+i)*2+2 >= indexLens {
				indexes[last+i] = indexes[(last+i)*2+1]
				break
			}
			indexes[last+i] = indexes[(last+i)*2+1] + indexes[(last+i)*2+2]
		}
		rowLens >>= 1
		last -= rowLens
	}
	l.indexes = indexes
}

func (l *SortedList) updateIndex(pos, incr int) {
	if len(l.indexes) > 0 {
		child := l.offset + pos
		for child > 0 {
			l.indexes[child] += incr
			child = (child - 1) >> 1
		}
		l.indexes[0] += 1
	}
}

func (l *SortedList) findPos(index int) (int, int) {
	if index < len(l.lists[0]) {
		return 0, index
	}
	pos := 0
	child := 1
	lenIndex := len(l.indexes)

	for child < lenIndex {
		indexChild := l.indexes[child]
		if index < indexChild {
			pos = child
		} else {
			index -= indexChild
			pos = child + 1
		}
		child = (pos << 1) + 1
	}
	return pos - l.offset, index
}

func (l *SortedList) resetIndex() {
	l.indexes = []int{}
	l.offset = 0
}

func roundUpOf2(a int) int {
	i := 1
	for ; i < a; i <<= 1 {
	}
	return i
}

func NewSortedList(c Compare, loadFactor int) SortedList {
	if loadFactor <= 0 {
		loadFactor = DefaultLoadFactor
	}
	return SortedList{load: loadFactor, c: c}
}

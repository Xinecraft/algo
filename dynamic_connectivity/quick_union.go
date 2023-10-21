package dynamic_connectivity

type QuickUnion struct {
	list []int
}

func MakeQuickUnion(size int) QuickUnion {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	qu := QuickUnion{
		list: arr,
	}

	return qu
}

// O(N)
func (l *QuickUnion) Connected(p int, q int) bool {
	return l.root(p) == l.root(q)
}

// O(N)
func (l *QuickUnion) Union(p int, q int) {
	pRoot := l.root(p)
	qRoot := l.root(q)

	l.list[pRoot] = qRoot
}

func (l *QuickUnion) root(p int) int {
	idx := p
	for l.list[idx] != idx {
		idx = l.list[idx]
	}
	return idx
}

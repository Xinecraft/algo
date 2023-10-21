package dynamic_connectivity

// Quick Union Weighted (Improved)
type QuickUnionWeighted struct {
	list   []int
	weight []int
}

func MakeQuickUnionWeighted(size int) QuickUnionWeighted {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	quimp := QuickUnionWeighted{
		list:   arr,
		weight: make([]int, size),
	}
	return quimp
}

// O(LogN)
func (l *QuickUnionWeighted) Connected(p int, q int) bool {
	return l.root(p) == l.root(q)
}

// O(LogN)
func (l *QuickUnionWeighted) Union(p int, q int) {
	pRoot := l.root(p)
	qRoot := l.root(q)

	if l.weight[pRoot] < l.weight[qRoot] {
		l.list[pRoot] = qRoot
		l.weight[qRoot] = l.weight[qRoot] + l.weight[pRoot]
	} else {
		l.list[qRoot] = pRoot
		l.weight[pRoot] = l.weight[pRoot] + l.weight[qRoot]
	}
}

func (l *QuickUnionWeighted) root(p int) int {
	idx := p
	for l.list[idx] != idx {
		idx = l.list[idx]
	}
	return idx
}

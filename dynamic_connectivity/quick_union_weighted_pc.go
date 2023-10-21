package dynamic_connectivity

// Quick Union Weighted with Path Compression (Improved++)
type QuickUnionWeightedPathCompressed struct {
	list   []int
	weight []int
}

func MakeQuickUnionWeightedPathCompressed(size int) QuickUnionWeightedPathCompressed {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	quwpc := QuickUnionWeightedPathCompressed{
		list:   arr,
		weight: make([]int, size),
	}
	return quwpc
}

// O(LogN)
func (l *QuickUnionWeightedPathCompressed) Connected(p int, q int) bool {
	return l.root(p) == l.root(q)
}

// O(LogN)
func (l *QuickUnionWeightedPathCompressed) Union(p int, q int) {
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

func (l *QuickUnionWeightedPathCompressed) root(p int) int {
	idx := p
	for l.list[idx] != idx {
		l.list[idx] = l.list[l.list[idx]] // Path Compression: Make every other node in path point to its grandparent (thereby halving path length)
		idx = l.list[idx]
	}

	return idx
}

package dynamic_connectivity

type QuickFind struct {
	list []int
}

func MakeQuickFind(size int) QuickFind {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}
	qf := QuickFind{
		list: arr,
	}

	return qf
}

// O(1)
func (qf *QuickFind) Connected(p int, q int) bool {
	return qf.list[p] == qf.list[q]
}

// O(N)
func (qf *QuickFind) Union(p int, q int) {
	if qf.Connected(p, q) {
		return
	}

	pVal := qf.list[p]
	qVal := qf.list[q]

	for i := 0; i < len(qf.list); i++ {
		if qf.list[i] == pVal {
			qf.list[i] = qVal
		}
	}
}

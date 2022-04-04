package p50

var Merge = merge

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	e1 := m - 1
	e2 := n - 1
	m = m + n - 1

	for e1 >= 0 && e2 >= 0 {
		if nums1[e1] >= nums2[e2] {
			nums1[m] = nums1[e1]
			e1--
		} else {
			nums1[m] = nums2[e2]
			e2--
		}
		m--
	}

	for e2 >= 0 {
		nums1[m] = nums2[e2]
		m--
		e2--
	}
}

func MergeKSortedList(numsList [][]int) []int {
	totalLen := 0
	l := len(numsList)
	for i := 0; i < l; {
		if len(numsList[i]) == 0 {
			l--
			numsList[i], numsList[l] = numsList[l], numsList[i]
			continue
		}
		totalLen += len(numsList[i])
		i++
	}

	minHeap := SliceListMinHeap(numsList)
	minHeap.InitMinHeap()

	ans := make([]int, 0, totalLen)
	for {
		l = len(minHeap)
		ans = append(ans, minHeap[0][0])
		minHeap[0] = minHeap[0][1:]
		if len(minHeap[0]) == 0 {
			l--
			if l == 0 {
				break
			}
			minHeap[0] = minHeap[l]
			minHeap = minHeap[:l]
		}
		minHeap.Down(0)
	}

	return ans
}

type SliceListMinHeap [][]int

func (h SliceListMinHeap) Up() {
	p := len(h)
	var r int
	for {
		r = p/2 - 1
		if r < 0 || h[p][0] >= h[r][0] {
			break
		}

		h[p], h[r] = h[r], h[p]
		p = r
	}
}

func (h SliceListMinHeap) Down(p int) {
	length := len(h)
	var l, r int
	for {
		l = p*2 + 1
		r = l + 1
		if l >= length {
			break
		}

		if r < length && h[r][0] < h[l][0] {
			l = r
		}

		if h[p][0] < h[l][0] {
			break
		}

		h[p], h[l] = h[l], h[p]

		p = l
	}
}

func (h SliceListMinHeap) InitMinHeap() {
	for i := len(h)/2 - 1; i >= 0; i-- {
		h.Down(i)
	}
}

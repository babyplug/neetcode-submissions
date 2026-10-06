
type Item struct {
	num int
	count int 
}

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() any {
    old := *h
    n := len(old)

    item := old[n-1]
    *h = old[:n-1]

    return item
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int) // number/frequent

	for _, num := range nums {
		freq[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, count := range freq { 
		heap.Push(h, Item{
			num: num,
			count: count,
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

    result := make([]int, 0, k)

    for h.Len() > 0 {
        item := heap.Pop(h).(Item)
        result = append(result, item.num)
    }

    return result
}

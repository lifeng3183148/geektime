package main

import (
	"container/heap"	
)

func main() {
	maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3)
}

func maxSlidingWindow(nums []int, k int) []int {
    var (
        start = 0
        ans []int
        h InitHeap
    )

    for start < k {
        if start < len(nums) {
            heap.Push(&h, &Iterm{nums[start], start})
        } else {
            ans = append(ans, h.Top().val)
            return ans
        }
        start++
    }

    ans = append(ans, h.Top().val)

    for i := start; i < len(nums); i++ {
        heap.Push(&h, &Iterm{nums[i], i})
        for h.Top().index <= i-k {
            heap.Remove(&h, 0)
        }
        ans = append(ans, h.Top().val)
    }
    // heap.Init(h)
    // fmt.Printf("%+v\n", h)

    return ans
}

type Iterm struct {
    val int
    index int
}

// func (i Iterm) String() string {
//     return fmt.Sprintf("Iterm{val: %v}", i.val)
// }

type InitHeap []*Iterm

func (h InitHeap) Len() int {return len(h)}
func (h InitHeap) Less(i, j int) bool {return h[i].val > h[j].val}
func (h InitHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
    // h[i].index, h[j].index = i, j
}

func (h *InitHeap) Push(x interface{}) {
    iterm := x.(*Iterm)
    *h = append(*h, iterm)
}

func (h *InitHeap) Pop() interface{} {
    n := h.Len()
    x := (*h)[n-1]
    (*h)[n-1] = nil
    *h = (*h)[:n-1]
    return x
}

func (h *InitHeap) Top() *Iterm {
    if h.Len() > 0 {
        return (*h)[0]
    }

    return nil
}

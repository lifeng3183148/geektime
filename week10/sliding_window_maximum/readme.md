### 解题思路

使用go标准库中的heap包，生成一个大顶堆，每当华东窗口移动时，新元素直接push，再判断堆顶元素是否在范围内，不在则直接移除并找到在范围内的元素为止

### 代码实现

```go
func maxSlidingWindow(nums []int, k int) []int {
    var ans []int
    var h IntHeap
    for i := 0; i < k; i++ {
        h = append(h, &Value{nums[i], i})
    }

    heap.Init(&h)
    ans = append(ans, h[0].val)

    for j := k; j < len(nums); j++ {
        heap.Push(&h, &Value{nums[j], j})
        // 移除超出范围的
        for h.Len() > 0 && j - h[0].index >= k {
            heap.Remove(&h, 0)
        }
        ans = append(ans, h[0].val)
    }

    return ans
}

type IntHeap []*Value

type Value struct {
    val int
    index int
}

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(*Value))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

### 时间空间复杂度

时间复杂度为O(n*nlogn)

空间复杂度为O(n)

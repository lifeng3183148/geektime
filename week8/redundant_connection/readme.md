### 解题思路

借鉴并查集合并操作的思路

一个环可以看做是一个集合

最后一个加入集合，且已经存在的元素即为组成环的最后一条边，返回即可

### 代码实现

```go
func findRedundantConnection(edges [][]int) []int {
    var s = &Set{
        fa: make([]int, len(edges)+1),
    }

    for i := 0; i < len(edges)+1; i++ {
        s.fa[i] = i
    }

    var ans = make([]int, 2)

    for _, edge := range edges {
        // fmt.Println(s.fa)
        if s.unionSet(edge[0], edge[1]) {
            ans[0], ans[1] = edge[0], edge[1]
        }
    }

    return ans
}

type Set struct {
    fa []int
}

func (s *Set) find(x int) int {
    if x == s.fa[x] {
        return x
    }

    s.fa[x] = s.find(s.fa[x])

    return s.fa[x]
}

// 合并，如果不在一个集合中说明还不是环，在一个集合中，说明是环
func (s *Set) unionSet(x, y int) bool {
    x, y = s.find(x), s.find(y)
    if x != y {
        s.fa[x] = y
        return false
    } 
    return true
}
```

### 时间空间复杂度

时间复杂度认为并查集的合并复杂度为O(1)，整个的时间复杂度则为O(n)

空间复杂度由于使用到了fa切片，空间复杂度为O(n)
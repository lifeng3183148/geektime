### 解题思路

遍历整个二位数组，把点的值是'1'的都使用并查集做聚合

只是最后过滤答案的时候要把值与index不相等的和值原数组中值为'1'的记录下来，使用map去重

### 代码实现

```go
func numIslands(grid [][]byte) int {
    n := len(grid)
    m := len(grid[0])

    var s = &Set{
        fa: make([]int, n*m),
        m: m,
    }

    // 初始化
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            s.fa[s.num(i, j)] = s.num(i, j)
        }
    }

    var dx = []int{0, 1}
    var dy = []int{1, 0}

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid[i][j] == '1' {
                // 向右向下查找相邻的点是否是'1'
                for k := 0; k < 2; k++ {
                    di := i + dy[k]
                    dj := j + dx[k]
                    if (di < n) && (dj < m) && grid[di][dj] == '1' {
                        s.unionSet(s.num(i, j), s.num(di, dj))
                    }
                }
            }
        }
    }

    var ans = make(map[int]bool)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            // fmt.Println(s.num(i, j), s.find(s.num(i, j)))
            // 将值与index不相等的和值原数组中值为'1'的记录下来，使用map去重
            if s.fa[s.num(i, j)] != s.num(i, j) || grid[i][j] == '1' {
                val := s.find(s.num(i, j))
                if !ans[val]{
                    ans[val] = true
                }
            }
        }
    }

    return len(ans)
}



type Set struct {
    fa []int
    m int // 列数
}

func (s *Set) num(i, j int) int {
    return i * s.m + j
}

func (s *Set) find(x int) int {
    if s.fa[x] == x {
        return x
    }

    s.fa[x] = s.find(s.fa[x])

    return s.fa[x]
}

func (s *Set) unionSet(x, y int) {
    x, y = s.find(x), s.find(y)

    if x != y {
        s.fa[x] = y
    }
}
```

### 时间空间复杂度

时间复杂度认为并查集的合并复杂度为O(1)，整个的时间复杂度则为O(n x m)

空间复杂度由于使用到了fa切片，空间复杂度为O(n x m)
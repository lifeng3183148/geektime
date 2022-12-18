### 解题思路

使用BFS，从起始位置(0,0)开始向8个方向搜索，直到找到(n-1,m-1)的点为止，这个问题知道终点位置，使用双向BFS会更优

### 代码实现

```go
func shortestPathBinaryMatrix(grid [][]int) int {
    l := len(grid)
    w := len(grid[0])

    var globalStep int = 1e9
    
    visit := make(map[[2]int]bool)
    // 坐标      ↗  →  ↘  ↓  ↙   ←  ↖   ↑
    dx := []int{-1, 0, 1, 1, 1, 0, -1, -1}
    dy := []int{1, 1, 1, 0, -1, -1, -1, 0}
    queue := make([][3]int, 0)

    queue = append(queue, [3]int{0, 0, 1})
    // 设置访问过的点 
    visit[[2]int{0, 0}] = true

    for len(queue) != 0 {
        // 取出第一个元素
        x := queue[0][0]
        y := queue[0][1]
        step := queue[0][2]
        queue = queue[1:] // 去掉第一个元素

        // 判断是否合法
        if !isValid(l, w, x, y, grid) {
            continue
        }

        // 是否找到
        if x == l-1 && y == w-1 {
            globalStep = int(math.Min(float64(globalStep), float64(step)))
            continue
        }

        // 朝 ↗→↘↓↙ 5个方向前进
        for i := 0; i < 8; i++ {
            if x+dx[i] >= 0 && y+dy[i] >= 0 {
                if visit[[2]int{x+dx[i], y+dy[i]}] {
                    continue
                }
                queue = append(queue, [3]int{x+dx[i], y+dy[i], step+1})
                // 设置访问过的点 
                visit[[2]int{x+dx[i], y+dy[i]}] = true
            }
        }
    }

    if globalStep == 1e9 {
        return -1 
    } else {
        return globalStep
    }
}

// 超出范围和非0都不合法
func isValid(l, w, x, y int, grid [][]int) bool {
    if x >= l || y >= w{
        return false
    }

    if grid[x][y] != 0 {
        return false
    }

    return true
}
```

### 时间空间复杂度

时间复杂度为O(n*m)

使用了visit map记录是否访问过，空间复杂度为O(n*m)

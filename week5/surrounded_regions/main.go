package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func solve(board [][]byte)  {
    row := len(board)
    col := len(board[0])

    for i := 0;i < row; i++ {
        for j := 0; j < col; j++ {
            if board[i][j] == 'O' {
                scan := &Scan{
                    row: row,
                    col: col,
                    visited: make(map[string]bool),
                }

                if scan.dfs(board, i, j) {
                    for _, val := range scan.result {
                        board[val[0]][val[1]] = 'X'
                    }
                }
            }   
        }
    }
}

type Scan struct {
    result [][]int
    row, col int
    visited map[string]bool
}

func (s *Scan) dfs(board [][]byte, i, j int) bool {
    if (i < 0 || i > s.row-1) || (j < 0 || j > s.col-1) {
        return false
    }

    if board[i][j] == 'X' {
        return true
    }

    s.visited[fmt.Sprintf("%d%d", i, j)] = true
    s.result = append(s.result, []int{i,j})

    // 上
    var up, down, left, right bool = true, true, true, true
    if !s.visited[fmt.Sprintf("%d%d", i-1, j)] {
        up = s.dfs(board, i-1, j)
    }
    
    // 下
    if !s.visited[fmt.Sprintf("%d%d", i+1, j)] {
        down = s.dfs(board, i+1, j)
    }
    // 左
    if !s.visited[fmt.Sprintf("%d%d", i, j-1)] {
        left = s.dfs(board, i, j-1)
    }
    // 右
    if !s.visited[fmt.Sprintf("%d%d", i, j+1)] {
        right = s.dfs(board, i, j+1)
    }

    return up && down && left && right
}

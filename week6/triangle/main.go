package main

import (
	"fmt"
	"math"
)

func main() {
	min := minimumTotal([][]int{{2},{3,4},{6,5,7},{4,1,8,3}})
	fmt.Println(min)
}

func minimumTotal(triangle [][]int) int {
    var ans = make([][]int, len(triangle))
    for i := len(triangle)-1; i >= 0; i-- {
        if len(ans[i]) == 0 {
            ans[i] = make([]int, len(triangle[i]))
        }
        for j := 0; j < len(triangle[i]); j++ {
            if i == len(triangle)-1 {
                ans[i][j] = triangle[i][j]
            } else {
                ans[i][j] = int(math.Min(float64(triangle[i][j] + ans[i+1][j]), float64(triangle[i][j] + ans[i+1][j+1])))
            }
        }
    }

    return ans[0][0]
}

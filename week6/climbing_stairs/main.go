package main

import "fmt"

func main() {
	ans := climbStairs(8)
	fmt.Println(ans)
}

func climbStairs(n int) int {
    if n < 3 {
        return n
    }
    var ans = make([]int, n+1)
    ans[1], ans[2] = 1, 2
    for i := 3; i <= n; i++ {
        ans[i] = ans[i-1] + ans[i-2]
    }

    return ans[n]
}

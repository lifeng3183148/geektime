package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(13))
}

func numSquares(n int) int {
    var f = make([]int, n+1)
    f[0] = 0
    f[1] = 1
    for i := 2; i <= n; i++ {
        var min int = 1e9
        for j := 1; j*j <= i; j++ {
            min = int(math.Min(float64(min), float64(f[i-j*j])))
        }

        f[i] = min + 1
    }

    return f[n]
}

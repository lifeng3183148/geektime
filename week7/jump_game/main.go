package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{1, 2, 3, 4}))
}

func canJump(nums []int) bool {
    var n = len(nums)
    if n < 2 {
        return true
    }

    var f = make([]bool, n)
    f[n-1] = true


    for i := n-2; i >= 0; i-- {
        var result bool = false
        for j := 0; j <= nums[i]; j++ {
            if i+j < n {
                result = result || f[i+j]
            }
        }
        f[i] = result
    }

    return f[0]
}

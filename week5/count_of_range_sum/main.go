package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func countRangeSum(nums []int, lower int, upper int) int {
    var preSumI = make([]int, len(nums))
    for i, num := range nums {
        if i == 0 {
            preSumI[i] = num
        } else {
            preSumI[i] = num + preSumI[i-1]
        }
    }

    // var ans [][]int
    count := 0
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            value := 0
            if i == 0 {
                value = preSumI[j]
            } else {
                value = preSumI[j] - preSumI[i-1]
            }

            if value >= lower && value <= upper {
                // ans = append(ans, []int{i, j})
                count++
            }
        }
    }

    // fmt.Println(ans)

    return count
}

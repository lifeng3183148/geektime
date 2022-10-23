package main

func main() {
    findShortestSubArray([]int{1,2,2,3,1})
}

func findShortestSubArray(nums []int) int {
    type startEnd struct {
        min int
        max int
        count int
    }

    var m = make(map[int]*startEnd)
    var maxCount int = 0
    var maxCountMap = make(map[int][]int)

    for i := range nums {
        if _, ok := m[nums[i]]; ok {
            m[nums[i]].max = i
            m[nums[i]].count++
        } else {
            m[nums[i]] = &startEnd{
                min: i,
                max: i,
                count: 1,
            }
        }

        if m[nums[i]].count >= maxCount {
            maxCount =  m[nums[i]].count
            maxCountMap[maxCount] = append(maxCountMap[maxCount], nums[i])
        }
    }

    if len(maxCountMap) == 0 {
        return 0
    }

    var minLen int = m[maxCountMap[maxCount][0]].max + 1 - m[maxCountMap[maxCount][0]].min
    for j := range maxCountMap[maxCount] {
        if m[maxCountMap[maxCount][j]].max + 1 - m[maxCountMap[maxCount][j]].min < minLen {
            minLen = m[maxCountMap[maxCount][j]].max + 1 - m[maxCountMap[maxCount][j]].min
        }
    }
    return minLen
}

## 解法一

### 解题思路

两重循环，第一重遍历所有的数，第二重以外层循环遍历的数为起始，遍历至末尾，检查是否有和为K的子数组

### 代码实现

```go
func subarraySum(nums []int, k int) int {
    n := len(nums)
    count := 0
    for i := 0; i < n; i++ {
        sum := 0 // 每次遍历完一个数要清零
        for j := i; j < n; j++ {
            sum += nums[j]
            if sum == k {
                count++
            }
        }
    }
    return count
}
```

### 时间空间复杂度

时间复杂度为O(n!)，约等于O(n*n)

空间复杂度为O(1)

## 解法二

### 解题思路

利用前缀和，把两重循环优化成一重循环，利用map存下前缀和

### 代码实现

```go
func subarraySum(nums []int, k int) int {
    n := len(nums)
    perSum := make([]int, n)
    perSum[0] = nums[0]
    count := 0
    if perSum[0] == k {
        count++
    }

    var m = make(map[int]int)
    m[perSum[0]] = 1
    for i := 1; i < n; i++ {
        perSum[i] = perSum[i-1] + nums[i]
        // 前缀和为K count++
        if perSum[i] == k {
            count++
        }
        // 如果 cur前缀和 - before前缀和 == K count += v，v是之前值相同的前缀和出现的次数
        if v, ok := m[perSum[i]-k]; ok {
            count += v
        }
        // 如果 cur前缀和的值在map中已存在，计数++，否则记录下这个新值，计数计为1
        if _, ok := m[perSum[i]]; ok {
            m[perSum[i]]++
        }else {
            m[perSum[i]] = 1
        }

    }
    
    return count
}
```

### 时间空间复杂度

时间复杂度为O(n)

空间复杂度由于使用到了map，空间复杂度为O(n)
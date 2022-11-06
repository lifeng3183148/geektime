package main

func main() {
	obj := Constructor([]int{}, []int{});
	_ = obj.Q(0)
}

type TopVotedCandidate struct {
    times []int
    tops []int
}


func Constructor(persons []int, times []int) TopVotedCandidate {
    var tops = make([]int, 0, len(persons))
    var top = -1
    var ans = make(map[int]int)
    for _, p := range persons {
        ans[p]++
        if ans[p] >= ans[top] {
            top = p
        }
        tops = append(tops, top)
    }

    return TopVotedCandidate{
        times: times,
        tops: tops,
    }
}


func (this *TopVotedCandidate) Q(t int) int {
    left := 0
    right := len(this.times)-1

    for left < right {
        mid := (left+right+1)>>1
        if this.times[mid] <= t {
            left = mid
        } else {
            right = mid - 1
        }
    }

    if right == -1 {
        return -1
    }

    return this.tops[right]
}


package main

import (
	"strings"
	"strconv"
	"fmt"
)

func mian() {
    subdomainVisits([]string{["9001 discuss.leetcode.com"]})
}

func subdomainVisits(cpdomains []string) []string {
    var domainsMap = make(map[string]int)
    for i := range cpdomains {
        strs1 := strings.Split(cpdomains[i], " ")
        if len(strs1) < 2 {
            return nil
        }

        count, err := strconv.Atoi(strs1[0])
        if err != nil {
            return nil
        }

        strs2 := strings.Split(strs1[1], ".")
        if len(strs2) < 1 {
            return nil
        }

        domainsMap[strs1[1]] += count

        var sum string
        var j = 1
        for j < len(strs2) {
            for i := j; i < len(strs2); i++ {
                if i == len(strs2) - 1 {
                    sum += strs2[i]
                } else {
                    sum += strs2[i] + "."
                }
            }
            domainsMap[sum] += count
            // fmt.Println(domainsMap, len(domainsMap))
            sum = ""
            j++
        }
    }

    // fmt.Println(domainsMap["vdu.bfo.team"])

    var result = make([]string, 0, len(domainsMap))
    for k, v := range domainsMap {
        result = append(result, fmt.Sprintf("%d %s", v, k))
    }

    return result
}

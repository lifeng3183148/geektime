package main

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
    mergeTwoLists(nil, nil)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {return list2}
    if list2 == nil {return list1}

    var head = list2
    var list2Pre *ListNode = nil
    for list1 != nil && list2 != nil {
        if list1.Val >= list2.Val {
            list2Pre, list2 = list2, list2.Next
            // list2.Next, list1.Next, list1, list2 = list1, list2.Next, list1.Next, list1
        } else {
            if list2Pre != nil {
                list2Pre.Next, list1.Next, list1, list2Pre = list1, list2, list1.Next, list1
            } else {
                head = list1
                list1.Next, list1, list2Pre = list2, list1.Next, list1
            }
        }
    }

    if list2 == nil {
        list2Pre.Next = list1
    }

    return head
}

package easy

func HasCycle(head *ListNode) bool {
    cnt := 0
    for head != nil {
        if cnt > 10000 {
            return true
        }
        head = head.Next
        cnt++
    }
    return false
}
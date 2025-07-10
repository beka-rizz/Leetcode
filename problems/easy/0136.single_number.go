package easy

func SingleNumber(nums []int) int {
    resMap := make(map[int]bool)
    var res int
    for _, num := range nums {
        if _, ok := resMap[num]; ok {
            delete(resMap, num)
        } else {
            resMap[num] = true
        }
    }
    for k := range resMap {
        res = k
    }
    return res
}
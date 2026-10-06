func hasDuplicate(nums []int) bool {
    if len(nums) == 0 {
        return false
    }


    uniques := map[int]bool{}
    for _, num := range nums {
        if _, ok := uniques[num]; ok {
            return true
        }
        uniques[num] = true
    }

    return false
}

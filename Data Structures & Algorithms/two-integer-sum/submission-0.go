func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
    
    for i, num := range nums {
        if index, found := seen[target - num]; found {
            return []int{index, i}
        }
        seen[num] = i
    }

    return []int{}
}

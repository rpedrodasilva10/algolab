package algo

func TwoSumOptimized(nums []int, target int) []int {
    mem := make(map[int]int, len(nums))
    var rest int
    var answer []int
    for currentIndex, num := range nums {
        rest = target - num
        if foundIndex, exists := mem[rest]; exists {
            answer = append(answer, foundIndex, currentIndex)
            return answer
        }
        mem[num] = currentIndex
    }

    return answer
}

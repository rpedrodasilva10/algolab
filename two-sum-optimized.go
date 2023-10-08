package algo

// TwoSumOptimized is the optimized version of TwoSum
// It uses a map to store the index of the number and the number itself
// Then it iterates over the numbers and checks if the difference between the target and the current number is in the map
// If it is, it will return the index of the current number and the index of the difference
// If it is not, it will add the current number and its index to the map
// Time complexity: O(n)
// Space complexity: O(n)
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

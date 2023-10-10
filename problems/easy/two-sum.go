package easy

// TwoSum is the brute force version of the solution, reading the array twice, adding the pair and comparing them to target
// Time complexity: O(n^2)
// Space complexity: O(1)
func TwoSum(nums []int, target int) []int {
    var answer []int
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums); j++ {
            if i != j && target == nums[i]+nums[j] {
                answer = append(answer, i, j)
                return answer
            }
        }
    }

    return answer
}

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

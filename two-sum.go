package algo

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

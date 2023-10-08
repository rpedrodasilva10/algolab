package algo

func TwoSum(nums []int, target int) []int {
    /*
       Go through every and each element adding pairs to see if we reach target.
       Remember to avoid using the same index twice
    */
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

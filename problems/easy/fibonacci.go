package easy

// Fib returns the Nth fibonacci number
// Using recursion and always recalculating the sum of Nth minus one and Nth minus two.
// This way we can resolve the requested, but for large numbers (like the 50th number) is too slow because
// in this solution we are always calculating fibo numbers, even if we calculated them already in previous steps.
// Time complexity: O(2^n)
// Space complexity: O(n)
func Fib(n int) int {
    if n <= 2 {
        return 1
    }

    return Fib(n-1) + Fib(n-2)
}

// MemoizedFib returns the Nth fibonacci number, using a map as support for memoization.
// This is the same solution as the previous one but, using memoization, is much more efficient
// because avoids expensive recalculations using the memory map to retrieve previous calculated values
// Time complexity: O(n)
// Space complexity: O(n)
func MemoizedFib(n int, memo map[int]int) int {
    if n <= 2 {
        return 1
    }

    if fibNum, ok := memo[n]; ok {
        return fibNum
    } else {
        memo[n] = MemoizedFib(n-1, memo) + MemoizedFib(n-2, memo)
        return memo[n]
    }
}

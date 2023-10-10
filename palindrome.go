package algo

import (
    "fmt"
    "strconv"
)

// IsPalindrome is the brute force solution to check if a number is a palindrome
// We run all items inverting the string and compare with the original value
// Time complexity: O(n)
// Space complexity: O(n)
func IsPalindrome(x int) bool {
    strValue := strconv.Itoa(x)
    strArray := []rune(strValue)
    var invertedString string

    for j := len(strArray) - 1; j >= 0; j-- {
        invertedString += string(strArray[j])
    }

    return invertedString == strValue
}

// IsPalindromeOptimized is the optimized solution to check if a number is a palindrome
// We run half of the items inverting the string and compare with the original value
// Time complexity: O(n/2)
// Space complexity: O(n)
func IsPalindromeOptimized(x int) bool {
    strValue := strconv.Itoa(x)
    strArray := []rune(strValue)
    var invertedString string

    for j := len(strArray) - 1; j == len(strArray)/2; j-- {
        fmt.Println(j)
        invertedString += string(strArray[j])
    }

    return invertedString == strValue[1:len(strArray)/2]
}

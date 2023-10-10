package algo

import "strconv"

func IsPalindrome(x int) bool {
    strValue := strconv.Itoa(x)
    strArray := []rune(strValue)
    var invertedString string

    for j := len(strArray) - 1; j >= 0; j-- {
        invertedString += string(strArray[j])
    }

    return invertedString == strValue
}

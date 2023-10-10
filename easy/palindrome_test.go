package easy

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestIsPalindrome(t *testing.T) {

    tests := []struct {
        name     string
        inputNum int
        want     bool
    }{
        {
            name:     "should return true for 121 - positive number",
            inputNum: 121,
            want:     true,
        },
        {
            name:     "should return false for -121 - negative signed number",
            inputNum: -121,
            want:     false,
        },
        {
            name:     "should return false for 10",
            inputNum: 10,
            want:     false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.want, IsPalindrome(tt.inputNum), "IsPalindrome(%v)", tt.inputNum)
            assert.Equalf(t, tt.want, IsPalindromeOptimized(tt.inputNum), "IsPalindromeOptimized(%v)", tt.inputNum)
        })
    }
}

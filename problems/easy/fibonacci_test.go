package easy

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestFib(t *testing.T) {
    tests := []struct {
        name     string
        n        int
        expected int
    }{
        {
            name:     "should calculate the 1st fibonacci number",
            n:        1,
            expected: 1,
        },
        {
            name:     "should calculate the 7th fibonacci number",
            n:        7,
            expected: 13,
        },

        /* Too slow with this algo, check the memoized version
           {
              name:     "should calculate the 50th fibonacci number",
              n:        50,
              expected: 12586269025,
          },*/
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.expected, Fib(tt.n), "Fib(%v)", tt.n)
        })
    }
}
func TestMemoizedFib(t *testing.T) {
    tests := []struct {
        name     string
        n        int
        expected int
    }{
        {
            name:     "should calculate the 1st fibonacci number",
            n:        1,
            expected: 1,
        },
        {
            name:     "should calculate the 7th fibonacci number",
            n:        7,
            expected: 13,
        },

        {
            name:     "should calculate the 50th fibonacci number",
            n:        50,
            expected: 12586269025,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            assert.Equalf(t, tt.expected, MemoizedFib(tt.n, make(map[int]int)), "MemoizedFib(%v)", tt.n)
        })
    }
}

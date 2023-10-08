package algo

import (
    "reflect"
    "testing"
)

func Test_twoSum(t *testing.T) {
    type args struct {
        nums   []int
        target int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name: "should sum correctly",
            args: args{
                []int{2, 7, 11, 15},
                9,
            },
            want: []int{0, 1},
        },
        {
            name: "should sum correctly not using the same index twice",
            args: args{
                []int{3, 2, 4},
                6,
            },
            want: []int{1, 2},
        },
        {
            name: "should sum correctly using the same number with different index",
            args: args{
                []int{3, 3},
                6,
            },
            want: []int{0, 1},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := TwoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("twoSum() = %v, want %v", got, tt.want)
            }
        })
    }
}

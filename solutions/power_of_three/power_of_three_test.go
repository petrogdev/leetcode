package power_of_three

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans bool
}{

	{-1, false},
	{8, false},
	{2, false},
	{3, true},
	{81, true},
	{243, true},
}

func Test_isPowerOfThree(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, isIntPowerOfThree(tc.n), "Val:%v", tc)
	}
}
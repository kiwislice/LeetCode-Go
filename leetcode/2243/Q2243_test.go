package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q2243(t *testing.T) {
	assert := assert.New(t)

	result := digitSum("11111222223", 3)
	assert.Equal("135", result)

	result = digitSum("00000000", 3)
	assert.Equal("000", result)

	fmt.Println("===============================")
	result = digitSum("71818186138735364590516322993378229838446988388364431324753408563431136824898916288399", 85)
	assert.Equal("4169", result)

}

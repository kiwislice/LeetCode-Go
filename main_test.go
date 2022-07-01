package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	assert := assert.New(t)

	ar := make([]int, 0, 10)
	ar = push(ar)
	fmt.Println(ar)

	assert.Fail("end.")
}

func push(ar []int) []int {
	return append(ar, 0)
}

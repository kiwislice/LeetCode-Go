package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Q0399(t *testing.T) {
	fmt.Println("test start.")
	assert := assert.New(t)
	var result any

	result = calcEquation([][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}})
	assert.Equal([]float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000}, result)

	result = calcEquation([][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0},
		[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}})
	assert.Equal([]float64{3.75000, 0.40000, 5.00000, 0.20000}, result)

	result = calcEquation([][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}, []float64{3.0, 4.0, 5.0, 6.0},
		[][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}})
	assert.Equal([]float64{360.00000, 0.00833, 20.00000, 1.00000, -1.00000, -1.00000}, result)

	result = calcEquation([][]string{{"a", "e"}, {"b", "e"}}, []float64{4.0, 3.0},
		[][]string{{"a", "b"}, {"e", "e"}, {"x", "x"}})
	assert.Equal([]float64{1.33333, 1.00000, -1.00000}, result)

	fmt.Println("test finished.")
}

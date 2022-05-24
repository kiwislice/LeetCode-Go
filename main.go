package main

import (
	"fmt"
	"os"
)

func main() {
	makeQuestionDir(os.Args[1])
}

func makeQuestionDir(qNumber string) {
	questionNum := "000" + qNumber
	questionNum = questionNum[len(questionNum)-4:]
	qDir := "leetcode/" + questionNum
	os.Mkdir(qDir, 0777)

	qPath := qDir + "/Q" + questionNum + ".go"
	createFile(qPath, "package leetcode\n\n\n\n\n\n")

	testPath := qDir + "/Q" + questionNum + "_test.go"
	testText := `package leetcode

	import (
		"testing"
	
		"github.com/stretchr/testify/assert"
	)
	
	func Test_Q` + questionNum + `(t *testing.T) {
		assert := assert.New(t)
	
		// result := countBits(2)
		// assert.Equal([]int{0, 1, 1}, result)
	}
	`
	createFile(testPath, testText)
}

func createFile(filepath, text string) {
	fout, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()
	fout.WriteString(text)
}

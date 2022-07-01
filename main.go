package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	switch os.Args[1] {
	case "newDir":
		makeQuestionDir(os.Args[2])
	case "replaceArray":
		replaceArray(os.Args[2])
	}
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
		"fmt"
		"testing"
	
		"github.com/stretchr/testify/assert"
	)
	
	func Test_Q` + questionNum + `(t *testing.T) {
		fmt.Println("test start.")
		assert := assert.New(t)
		var result any
	
		// result = countBits(2)
		// assert.Equal([]int{0, 1, 1}, result)
		
		fmt.Println("test finished.")
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

func replaceArray(text string) {
	text = strings.ReplaceAll(text, "[", "{")
	text = strings.ReplaceAll(text, "]", "}")
	fmt.Println(text)
}

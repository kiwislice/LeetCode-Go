package leetcode

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func IntArray(s string) []int {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	split := strings.Split(s, ",")

	rtn := make([]int, 0, len(split))
	for i := range split {
		if split[i] == "" {
			continue
		}
		i, err := strconv.Atoi(split[i])
		if err != nil {
			panic(err)
		}
		rtn = append(rtn, i)
	}
	return rtn
}

func Int2Array(s string) [][]int {
	var data [][]int
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ReadJsonFile_Int2Array(path string) [][]int {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var data [][]int
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ReadJsonFile_IntArray(path string) []int {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var data []int
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ReadJsonFile_Map(path string) map[string]any {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var data map[string]any
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ToInt1d(o any) []int {
	ar := o.([]any)
	return Any1dToInt1d(ar)
}

func Any1dToInt1d(ar []any) []int {
	rtn := make([]int, len(ar))
	for i, a := range ar {
		rtn[i] = int(a.(float64))
	}
	return rtn
}

func ToInt2d(o any) [][]int {
	ar := o.([]any)
	rtn := make([][]int, len(ar))
	for i, a := range ar {
		rtn[i] = ToInt1d(a)
	}
	return rtn
}

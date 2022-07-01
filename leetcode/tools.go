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

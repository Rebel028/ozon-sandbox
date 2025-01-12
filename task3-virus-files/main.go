package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	tests "ozon-tests-runner"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	zipPath := "virus-files-go.zip"
	timeLimit := 1000 * time.Millisecond
	memoryLimit := 256 * 1024 * 1024 // 256 MB

	runner := tests.NewTestRunner(zipPath, timeLimit, uint64(memoryLimit))
	runner.RunTests(Solve)
}

func Solve(in *bufio.Reader, out *bufio.Writer) {
	inputData := parseInput(in)
	outputData := prepareOutput(inputData)
	outputStr := strings.Join(outputData, "\n")
	_, _ = fmt.Fprint(out, outputStr+"\n")
}

func parseInput(in *bufio.Reader) (inputData []Directory) {
	scanner := bufio.NewScanner(in)
	//scanner.Buffer(make([]byte, 3000000), 3000000)
	scanner.Scan()
	expectedCount, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		linesToScan, _ := strconv.Atoi(line)
		var jsonStr []byte
		for i := 0; i < linesToScan; i++ {
			scanner.Scan()
			jsonStr = append(jsonStr, scanner.Bytes()...)
		}
		var item Directory
		err := json.Unmarshal(jsonStr, &item)
		if err != nil {
			log.Fatal(err)
		}

		inputData = append(inputData, item)
	}

	if len(inputData) != expectedCount {
		log.Panicf("Expected %d lines, got %d", expectedCount, len(inputData))
	}
	return
}

func prepareOutput(data []Directory) (results []string) {
	for _, item := range data {
		result := countInfectedFiles(item)
		results = append(results, result)
	}
	return
}

func countInfectedFiles(item Directory) string {
	c := 0
	stack := []Directory{item}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.IsInfected() {
			c += addAllToInfected(current)
		} else {
			for _, child := range current.Folders {
				stack = append(stack, child)
			}
		}
	}
	return strconv.Itoa(c)
}

func addAllToInfected(dir Directory) (c int) {
	stack := []Directory{dir}
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		c += len(current.Files)
		for _, child := range current.Folders {
			stack = append(stack, child)
		}
	}
	return
}

type Directory struct {
	Dir     string      `json:"dir"`
	Files   []string    `json:"files"`
	Folders []Directory `json:"folders,omitempty"`
}

func (dir *Directory) IsInfected() bool {
	for _, file := range dir.Files {
		matched, _ := regexp.Match(".hack$", []byte(file))
		if matched {
			return true
		}
	}
	return false
}

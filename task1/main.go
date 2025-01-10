package main

import (
	"bufio"
	"fmt"
	"log"
	tests "ozon-tests-runner"
	"strconv"
	"strings"
	"time"
)

func Task1(in *bufio.Reader, out *bufio.Writer) {
	inputData := parseInput(in)
	outputData := prepareOutput(inputData)
	outputStr := strings.Join(outputData, "\n")
	_, _ = fmt.Fprint(out, outputStr+"\n")
}

func prepareOutput(inputData []string) (results []string) {
	for _, str := range inputData {
		result := removeDigit(str)
		results = append(results, result)
	}
	return
}

func removeDigit(input string) (result string) {
	length := len(input)
	if length == 1 {
		return "0"
	}

	digitToRemove := '9'
	for _, digit := range input {
		if digit == digitToRemove {
			continue
		}
		if digit < digitToRemove {
			digitToRemove = digit
			if digitToRemove == '0' {
				break
			}
		} else {
			break
		}
	}

	result = strings.Replace(input, string(digitToRemove), "", 1)
	return
}

func parseInput(in *bufio.Reader) (inputData []string) {
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 100001), 100001)
	scanner.Scan()
	expectedCount, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		inputData = append(inputData, line)
	}

	if len(inputData) != expectedCount {
		log.Panicf("Expected %d lines, got %d", expectedCount, len(inputData))
	}
	return
}

func main() {
	zipPath := "remove-digit.zip"
	timeLimit := 1000 * time.Millisecond
	memoryLimit := 256 * 1024 * 1024 // 256 MB

	runner := tests.NewTestRunner(zipPath, timeLimit, uint64(memoryLimit))
	runner.RunTests(Task1)
}

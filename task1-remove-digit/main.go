package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Solve(in, out)
}

func Solve(in *bufio.Reader, out *bufio.Writer) {
	inputData := parseInput(in)
	outputData := prepareOutput(inputData)
	outputStr := strings.Join(outputData, "\n")
	_, _ = fmt.Fprint(out, outputStr+"\n")
}

func prepareOutput(inputData chan string) (results []string) {
	for str := range inputData {
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

func parseInput(in *bufio.Reader) chan string {
	inputData := make(chan string)

	go func() {
		scanner := bufio.NewScanner(in)
		scanner.Buffer(make([]byte, 100001), 100001)
		scanner.Scan()
		//expectedCount, _ := strconv.Atoi(scanner.Text())
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			inputData <- line
		}
		//if len(inputData) != expectedCount {
		//	log.Panicf("Expected %d lines, got %d", expectedCount, len(inputData))
		//}
		defer close(inputData)
	}()

	return inputData
}

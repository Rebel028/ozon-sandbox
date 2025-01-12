package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

func validateResult(data testData) string {
	// sort array
	sort.Slice(data.TestArray, func(i, j int) bool {
		return data.TestArray[i] < data.TestArray[j]
	})
	joinedSortedArr := strings.Trim(fmt.Sprint(data.TestArray), "[]")
	if joinedSortedArr == data.TestResult {
		return "yes"
	} else {
		return "no"
	}
}

func parseInput(in *bufio.Reader) (inputData []testData) {
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 3000000), 3000000)
	scanner.Scan()
	expectedCount, _ := strconv.Atoi(scanner.Text())

	var item = testData{}
	scan := true
	for scan {
		var arrLen string
		var arr string
		scan, arrLen = getNextLine(scanner)
		if !scan {
			break
		}
		scan, arr = getNextLine(scanner)
		scan, item.TestResult = getNextLine(scanner)

		item.TestArray = readArray(arr)
		expectedLen, _ := strconv.Atoi(arrLen)
		if len(item.TestArray) != expectedLen {
			log.Printf("Invalid test array length %d, expected %d", len(item.TestArray), expectedLen)
		}

		inputData = append(inputData, item)
	}

	if len(inputData) != expectedCount {
		log.Panicf("Expected %d lines, got %d", expectedCount, len(inputData))
	}
	return
}

func readArray(arr string) (result []int) {
	for _, s := range strings.Split(arr, " ") {
		n, _ := strconv.Atoi(s)
		result = append(result, n)
	}
	return
}

func prepareOutput(data []testData) (results []string) {
	for _, str := range data {
		result := validateResult(str)
		results = append(results, result)
	}
	return
}

func getNextLine(scanner *bufio.Scanner) (success bool, line string) {
	success = scanner.Scan()
	line = scanner.Text()
	if line == "" {
		success = false
	}
	return
}

type testData struct {
	TestArray  []int
	TestResult string
}

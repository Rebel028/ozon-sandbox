package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

func parseInput(in *bufio.Reader) chan Directory {
	var inputData = make(chan Directory)
	go func() {
		scanner := bufio.NewScanner(in)
		scanner.Scan()
		expectedCount, _ := strconv.Atoi(scanner.Text())
		i := 0

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
			inputData <- item
			i++
		}
		if i != expectedCount {
			log.Panicf("Expected %d lines, got %d", expectedCount, len(inputData))
		}
		defer close(inputData)
	}()

	return inputData
}

func prepareOutput(data chan Directory) (results []string) {
	for item := range data {
		//result := countInfectedFiles(item)
		result := strconv.Itoa(countInfectedRecursive(item, false))
		results = append(results, result)
	}
	return
}

func countInfectedRecursive(dir Directory, parentInfected bool) (count int) {
	thisInfected := parentInfected || dir.IsInfected()
	if thisInfected {
		count = len(dir.Files)
	}
	for _, folder := range dir.Folders {
		count += countInfectedRecursive(folder, thisInfected)
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
		if filepath.Ext(file) == ".hack" {
			return true
		}
	}
	return false
}

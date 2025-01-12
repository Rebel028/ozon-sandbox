package ozon_tests_runner

import (
	"archive/zip"
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
	"sort"
	"strings"
	"time"
)

type TestRunner struct {
	zipFilePath string
	timeLimit   time.Duration
	memoryLimit uint64
}

func NewTestRunner(zipFilePath string, timeLimit time.Duration, memoryLimit uint64) *TestRunner {
	return &TestRunner{
		zipFilePath: zipFilePath,
		timeLimit:   timeLimit,
		memoryLimit: memoryLimit,
	}
}

func (tr *TestRunner) RunTests(solve func(*bufio.Reader, *bufio.Writer)) {
	zipReader, err := zip.OpenReader(tr.zipFilePath)
	if err != nil {
		log.Fatalf("failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	files := zipReader.File

	// Sort files by name
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name < files[j].Name
	})

	for i := 0; i < len(files); i += 2 {
		inputFile, outputFile := zipReader.File[i], zipReader.File[i+1]
		log.Printf("Reading files %s and %s", inputFile.Name, outputFile.Name)

		if !strings.HasSuffix(outputFile.Name, ".a") {
			inputFile, outputFile = outputFile, inputFile
		}

		inputData, err := tr.readFile(inputFile)
		if err != nil {
			log.Fatalf("failed to read input file: %v", err)
		}

		expectedOutput, err := tr.readFile(outputFile)
		if err != nil {
			log.Fatalf("failed to read output file: %v", err)
		}

		var resultBuffer bytes.Buffer
		in := bufio.NewReader(strings.NewReader(inputData))
		out := bufio.NewWriter(&resultBuffer)

		ctx, cancel := context.WithTimeout(context.Background(), tr.timeLimit)
		defer cancel()

		done := make(chan bool)
		go func() {
			solve(in, out)
			out.Flush()
			done <- true
		}()

		select {
		case <-ctx.Done():
			log.Printf("Test exceeded time limit for input:\n%s", inputData)
		case <-done:
			var memStats runtime.MemStats
			runtime.ReadMemStats(&memStats)
			log.Printf("Used %d bytes of memory", memStats.Alloc)
			if memStats.Alloc > tr.memoryLimit {
				log.Fatalf("Test exceeded memory limit %d: Used %d", tr.memoryLimit, memStats.Alloc)
			}

			actualOutput := resultBuffer.String()
			if actualOutput != expectedOutput {
				log.Fatalf("Test failed for input:\n%s\nExpected: %s\nActual: %s", inputData[:100], expectedOutput, actualOutput)
			}
		}
	}

	fmt.Println("All tests passed successfully.")
}

func (tr *TestRunner) readFile(file *zip.File) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	var data strings.Builder
	_, err = io.Copy(&data, f)
	if err != nil {
		return "", err
	}

	return data.String(), nil
}

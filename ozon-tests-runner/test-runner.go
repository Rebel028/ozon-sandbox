package ozon_tests_runner

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

type TestRunner struct {
	zipFilePath string
}

func NewTestRunner(zipFilePath string) *TestRunner {
	return &TestRunner{
		zipFilePath: zipFilePath,
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

		solve(in, out)
		out.Flush()

		actualOutput := resultBuffer.String()

		if actualOutput != expectedOutput {
			log.Fatalf("Test failed for input:\n%s\nExpected: %s\nActual: %s", inputData, expectedOutput, actualOutput)
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

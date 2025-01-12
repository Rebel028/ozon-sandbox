package main

import (
	tests "ozon-tests-runner"
	"testing"
	"time"
)

func Test_Solve2(t *testing.T) {
	zipPath := "validate-output.zip"
	timeLimit := 1000 * time.Millisecond
	memoryLimit := 256 * 1024 * 1024 // 256 MB

	runner := tests.NewTestRunner(t, zipPath, timeLimit, uint64(memoryLimit))
	runner.RunTests(Solve)
}

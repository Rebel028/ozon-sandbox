package main

import (
	tests "ozon-tests-runner"
	"testing"
	"time"
)

func Test_Solve1(t *testing.T) {
	zipPath := "remove-digit.zip"
	timeLimit := 1000 * time.Millisecond
	memoryLimit := 256 * 1024 * 1024 // 256 MB

	runner := tests.NewTestRunner(t, zipPath, timeLimit, uint64(memoryLimit))
	runner.RunTests(Solve)
}
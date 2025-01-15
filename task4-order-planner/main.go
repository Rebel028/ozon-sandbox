package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Arrival struct {
	index int
	time  int
}

type Truck struct {
	index    int
	start    int
	end      int
	capacity int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	Solve(in, out)
}

func Solve(in *bufio.Reader, out *bufio.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, 28), 1024)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for tc := 0; tc < t; tc++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		arrivals := make([]*Arrival, n)
		for i := 0; i < n; i++ {
			scanner.Scan()
			time, _ := strconv.Atoi(scanner.Text())
			arrivals[i] = &Arrival{index: i, time: time}
		}

		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())

		trucks := make([]*Truck, m)
		for j := 0; j < m; j++ {
			scanner.Scan()
			start, _ := strconv.Atoi(scanner.Text())
			scanner.Scan()
			end, _ := strconv.Atoi(scanner.Text())
			scanner.Scan()
			capacity, _ := strconv.Atoi(scanner.Text())
			trucks[j] = &Truck{index: j + 1, start: start, end: end, capacity: capacity}
		}

		result := planOrders(arrivals, trucks)
		fmt.Fprintln(out, result)
	}
}

func planOrders(arrivals []*Arrival, trucks []*Truck) string {
	sort.Slice(arrivals, func(i, j int) bool {
		return arrivals[i].time < arrivals[j].time
	})

	sort.Slice(trucks, func(i, j int) bool {
		if trucks[i].start != trucks[j].start {
			return trucks[i].start < trucks[j].start
		}
		return trucks[i].index < trucks[j].index
	})

	results := make([]int, len(arrivals))
	for i := range results {
		results[i] = -1 // Default to -1, meaning no truck found
	}

	for _, arrival := range arrivals {
		for j := range trucks {
			if trucks[j].start <= arrival.time && arrival.time <= trucks[j].end && trucks[j].capacity > 0 {
				results[arrival.index] = trucks[j].index
				trucks[j].capacity--
				break
			}
		}
	}

	var sb strings.Builder
	for _, res := range results {
		sb.WriteString(strconv.Itoa(res) + " ")
	}
	return sb.String()
}

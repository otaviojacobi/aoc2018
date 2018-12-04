package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func star1(values []int64) int64 {
	var sum int64
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	return sum
}

func star2(values []int64) int64 {
	var sum int64
	sums := make(map[int64]bool)
	for true {
		for i := 0; i < len(values); i++ {
			sum += values[i]
			_, exists := sums[sum]

			if exists {
				return sum
			}

			sums[sum] = true
		}
	}

	return 0
}

func main() {

	start := time.Now().UnixNano() / int64(time.Microsecond)
	buf, _ := ioutil.ReadFile("input.txt") // just pass the file name

	lines := strings.Split(string(buf), "\n")

	values := make([]int64, len(lines))
	for i := 0; i < len(lines); i++ {
		valStr := strings.TrimSuffix(lines[i], "\r")
		val, _ := strconv.ParseInt(valStr, 10, 64)
		values[i] = val
	}

	firstStar := star1(values)
	secondStar := star2(values)

	fmt.Printf("Star 1 answer: %v\n", firstStar)
	fmt.Printf("Star 2 answer: %v\n", secondStar)

	end := time.Now().UnixNano() / int64(time.Microsecond)

	fmt.Printf("Execution took: %v μs\n", end-start)

}

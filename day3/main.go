package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func splitLine(line string) (uint64, uint64, uint64, uint64, uint64) {
	var id uint64
	var lin uint64
	var col uint64
	var width uint64
	var length uint64

	first := strings.Split(line, "@")
	id, _ = strconv.ParseUint(strings.Trim(first[0], "# "), 10, 64)
	second := strings.Split(first[1], ":")
	pos := strings.Split(strings.Trim(second[0], " "), ",")
	vals := strings.Split(strings.Trim(second[1], " "), "x")

	lin, _ = strconv.ParseUint(pos[0], 10, 64)
	col, _ = strconv.ParseUint(pos[1], 10, 64)
	width, _ = strconv.ParseUint(vals[0], 10, 64)
	length, _ = strconv.ParseUint(vals[1], 10, 64)

	return id, lin, col, width, length
}

func sumSlice(M *[1200][1200]uint64, lin uint64, width uint64, col uint64, length uint64) uint64 {

	var sum uint64
	for i := uint64(0); i < width; i++ {
		for j := uint64(0); j < length; j++ {
			sum += M[lin+i][col+j]
		}
	}
	return sum
}

func star1(lines []string) (uint64, *[1200][1200]uint64) {

	M := [1200][1200]uint64{}

	var overlap uint64
	for i := 0; i < len(lines); i++ {

		_, lin, col, width, length := splitLine(lines[i])

		for w := uint64(0); w < width; w++ {
			for l := uint64(0); l < length; l++ {
				M[lin+w][col+l]++
			}
		}
	}

	for i := 0; i < 1200; i++ {
		for j := 0; j < 1200; j++ {
			if M[i][j] >= 2 {
				overlap++
			}
		}
	}

	return overlap, &M
}

func star2(lines []string, M *[1200][1200]uint64) uint64 {
	for i := 0; i < len(lines); i++ {
		id, lin, col, width, length := splitLine(lines[i])
		if sumSlice(M, lin, width, col, length) == width*length {
			return id
		}
	}
	return 0
}

func main() {
	start := time.Now().UnixNano() / int64(time.Microsecond)

	buf, _ := ioutil.ReadFile("input.txt") // just pass the file name
	lines := strings.Split(string(buf), "\r\n")

	firstStar, M := star1(lines)
	secondStar := star2(lines, M)

	fmt.Printf("Star 1 answer: %v\n", firstStar)
	fmt.Printf("Star 2 answer: %v\n", secondStar)

	end := time.Now().UnixNano() / int64(time.Microsecond)

	fmt.Printf("Execution took: %v μs\n", end-start)
}

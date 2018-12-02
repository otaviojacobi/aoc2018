package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func toMap(line string) map[byte]int64 {

	m := make(map[byte]int64)
	for i := 0; i < len(line); i++ {
		c := line[i]
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	return m
}

func getDiff(str1 string, str2 string) int64 {
	var sum int64
	for c := 0; c < len(str1); c++ {
		if str1[c] != str2[c] {
			sum++
		}
	}
	return sum
}

func formatEquals(str1 string, str2 string) string {
	out := ""
	for c := 0; c < len(str1); c++ {
		if str1[c] == str2[c] {
			out += string(str1[c])
		}
	}

	return out
}

func star1(lines []string) int64 {

	var exact2 int64
	var exact3 int64
	var seen2 bool
	var seen3 bool

	for i := 0; i < len(lines); i++ {
		m := toMap(lines[i])
		seen2 = false
		seen3 = false
		for key := range m {
			if m[key] == 2 && !seen2 {
				exact2++
				seen2 = true
			}
			if m[key] == 3 && !seen3 {
				exact3++
				seen3 = true
			}

		}
	}
	return exact2 * exact3
}

func star2(lines []string) string {

	minDiff := int64(1000)
	var str1 string
	var str2 string

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0:i]); j++ {
			diff := getDiff(lines[i], lines[j])
			if diff < minDiff {
				minDiff = diff
				str1 = lines[i]
				str2 = lines[j]
			}
		}
	}

	return formatEquals(str1, str2)
}

func main() {
	start := time.Now().UnixNano() / int64(time.Microsecond)

	buf, _ := ioutil.ReadFile("input.txt") // just pass the file name

	lines := strings.Split(string(buf), "\n")

	firstStar := star1(lines)
	secondStar := star2(lines)

	fmt.Printf("Star 1 answer: %v\n", firstStar)
	fmt.Printf("Star 2 answer: %v\n", secondStar)

	end := time.Now().UnixNano() / int64(time.Microsecond)

	fmt.Printf("Execution took: %v μs\n", end-start)
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func reacts(a byte, b byte) bool {
	return a != b && strings.ToUpper(string(a)) == strings.ToUpper(string(b))
}

func star1(polymer string) int {
	var changed bool
	changed = true
	for changed {
		changed = false
		for i := 0; i < len(polymer)-1; i++ {
			if reacts(polymer[i], polymer[i+1]) {
				polymer = polymer[:i] + polymer[i+2:]
				changed = true
			}
		}
	}

	return len(polymer)
}

func star2(polymer string) int {
	var mini int
	mini = 999999
	allChars := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(allChars); i++ {
		smallPolymer := strings.Replace(polymer, string(allChars[i]), "", 60000)
		smallPolymer = strings.Replace(smallPolymer, strings.ToUpper(string(allChars[i])), "", 60000)

		val := star1(smallPolymer)
		if val < mini {
			mini = val
		}
	}
	return mini
}

func main() {
	start := time.Now().UnixNano() / int64(time.Microsecond)

	buf, _ := ioutil.ReadFile("input.txt")
	polymer := string(buf)

	firstStar := star1(polymer)
	secondStar := star2(polymer)

	fmt.Printf("Star 1 answer: %v\n", firstStar)
	fmt.Printf("Star 2 answer: %v\n", secondStar)

	end := time.Now().UnixNano() / int64(time.Microsecond)

	fmt.Printf("Execution took: %v μs\n", end-start)
}

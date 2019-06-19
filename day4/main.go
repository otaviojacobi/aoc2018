package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func getLineProperties(line string) (string, string, string, string, string, string, uint64) {

	splittedLine := strings.Split(line, "]")

	ts := splittedLine[0][1:]
	dayAndHour := strings.Split(ts, " ")
	date := dayAndHour[0]
	hours := dayAndHour[1]

	dateSplitted := strings.Split(date, "-")
	year := dateSplitted[0]
	month := dateSplitted[1]
	day := dateSplitted[2]

	hoursSplitted := strings.Split(hours, ":")
	hour := hoursSplitted[0]
	minute := hoursSplitted[1]

	lastPart := splittedLine[1]

	var lineType string
	var id string
	if strings.Contains(lastPart, "up") {
		lineType = "up"
	} else if strings.Contains(lastPart, "asleep") {
		lineType = "asleep"
	} else if strings.Contains(lastPart, "begins") {
		lineType = "begins"
		id = strings.Split(strings.Split(lastPart, "#")[1], " ")[0]
	}

	idInt, _ := strconv.ParseUint(id, 10, 64)

	return year, month, day, hour, minute, lineType, idInt
}

func getLineValue(line string) uint64 {
	year, month, day, hour, minute, _, _ := getLineProperties(line)
	value, _ := strconv.ParseUint(year+month+day+hour+minute, 10, 64)
	return value
}

func earlier(line1 string, line2 string) bool {
	val1 := getLineValue(line1)
	val2 := getLineValue(line2)
	return val1 > val2
}

func sortShifts(lines []string) []string {
	for i := 0; i < len(lines)-1; i++ {
		for j := 0; j < len(lines)-1; j++ {
			if earlier(lines[j], lines[i]) {
				tmp := lines[j]
				lines[j] = lines[i]
				lines[i] = tmp
			}
		}
	}
	return lines
}

func star1(lines []string) (uint64, *[10000][60]uint64) {
	var idsMap [10000][60]uint64
	var curGuard uint64
	var curGuardStartSleep uint64
	sortedShifts := sortShifts(lines)
	for i := 0; i < len(sortedShifts); i++ {
		_, _, _, _, minute, lineType, id := getLineProperties(sortedShifts[i])

		if lineType == "begins" {
			curGuard = id
		} else if lineType == "asleep" {
			curGuardStartSleep, _ = strconv.ParseUint(minute, 10, 64)
		} else if lineType == "up" {
			curTime, _ := strconv.ParseUint(minute, 10, 64)

			for j := curGuardStartSleep; j < curTime; j++ {
				idsMap[curGuard][j] += 1
			}
		}
	}

	var argmax uint64
	var max uint64
	var i uint64
	for i = 0; i < 10000; i++ {
		var sum uint64
		for j := 0; j < 60; j++ {
			sum += idsMap[i][j]
		}

		if sum > max {
			argmax = i
			max = sum
		}
	}

	var j uint64
	var argmax2 uint64
	var max2 uint64
	for j = 0; j < 60; j++ {
		if idsMap[argmax][j] > max2 {
			argmax2 = j
			max2 = idsMap[argmax][j]
		}
	}

	return argmax * argmax2, &idsMap
}

func star2(idsMap *[10000][60]uint64) uint64 {
	var argmax uint64
	var argmax2 uint64
	var max uint64
	var i uint64
	var j uint64
	for i = 0; i < 10000; i++ {
		for j = 0; j < 60; j++ {
			if idsMap[i][j] > max {
				argmax = i
				argmax2 = j
				max = idsMap[i][j]
			}
		}
	}
	return argmax * argmax2
}

func main() {
	start := time.Now().UnixNano() / int64(time.Microsecond)

	buf, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(buf), "\n")

	firstStar, M := star1(lines)
	secondStar := star2(M)

	fmt.Printf("Star 1 answer: %v\n", firstStar)
	fmt.Printf("Star 2 answer: %v\n", secondStar)

	end := time.Now().UnixNano() / int64(time.Microsecond)

	fmt.Printf("Execution took: %v μs\n", end-start)
}

package banner

import (
	"fmt"
)


type Theme interface {
	Convert(byte) Ascii
}


type Ascii []string

func (one Ascii) Show() {
	for _, line := range one {
		fmt.Println(line)
	}
}

func (one Ascii) isEmpty() bool {
	for _, line := range one {
		for _, v := range line {
			if v != ' ' {
				return false
			}
		}
	}
	return true
}

func (one Ascii) simpleJoin(other Ascii) {
	for i, _ := range one {
		one[i] += other[i]
	}
}

func (one Ascii) getBlankNumbers(fromBack bool) []int {
	nums := []int{}
	for _, line := range one {
		num := 0
		for i := 0; i < len(line); i++ {
			idx := i
			if fromBack {
				idx = len(line) - i - 1
			}
			if line[idx] == ' ' {
				num++
			} else {
				break
			}
		}
		nums = append(nums, num)
	}
	return nums
}

func getSumMin(firsts, seconds []int) int {
	min := int(^uint(0) >> 1)
	for i, _ := range firsts {
		sum := firsts[i] + seconds[i]
		if min > sum {
			min = sum
		}
	}
	return min
}

type joinRecord struct {
	backPos		int
	forwardPos	int
	first		byte
	second		byte
}

func (j *joinRecord) hasEmpty() bool {
	return j.first == ' ' || j.second == ' '
}

func (j *joinRecord) isPair() bool {
	return j.first == '/' && j.second == '/' ||
		j.first == '\\' && j.second == '\\'
}

func (j *joinRecord) shouldAdjoin() bool {
	return j.hasEmpty() || j.isPair()
}


func (one Ascii) complexJoin(other Ascii, adjoin bool) {
	backNums := one.getBlankNumbers(true)
	forwardNums := other.getBlankNumbers(false)
	eraseNum := getSumMin(backNums, forwardNums)

	joinRecords := []joinRecord{}
	for i, _ := range one {
		var backPos, forwardPos int
		if eraseNum > backNums[i] {
			backPos = backNums[i]
			forwardPos = eraseNum - backPos
		} else {
			backPos = eraseNum
			forwardPos = 0
		}
		joinRecords = append(joinRecords, joinRecord{
			backPos: backPos,
			forwardPos: forwardPos,
			first: one[i][len(one[i]) - backPos - 1],
			second: other[i][forwardPos],
		})
	}

	hasPair := false
	shouldAdjoin := true
	for _, record := range joinRecords {
		if record.isPair() {
			hasPair = true
		}
		if !record.shouldAdjoin() {
			shouldAdjoin = false
		}
	}
	if !hasPair {
		shouldAdjoin = false
	}

	for i, _ := range one {
		record := joinRecords[i]
		backErase := len(one[i]) - record.backPos
		forwardErase := record.forwardPos
		if shouldAdjoin && adjoin {
			if one[i][backErase - 1] == ' ' {
				backErase -= 1
			} else {
				forwardErase += 1
			}
		}
		one[i] = one[i][:backErase] + other[i][forwardErase:]
	}
}

func (one Ascii) Append(other Ascii, adjoin bool) {
	if one.isEmpty() || other.isEmpty() {
		one.simpleJoin(other)
	} else {
		one.complexJoin(other, adjoin)
	}
}



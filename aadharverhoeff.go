package aadharverhoeff


import (
	"regexp"
)

const aadhar = "[0-9]{12}"

var (
	d = [][]int {
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 2, 3, 4, 0, 6, 7, 8, 9, 5},
		{2, 3, 4, 0, 1, 7, 8, 9, 5, 6},
		{3, 4, 0, 1, 2, 8, 9, 5, 6, 7},
		{4, 0, 1, 2, 3, 9, 5, 6, 7, 8},
		{5, 9, 8, 7, 6, 0, 4, 3, 2, 1},
		{6, 5, 9, 8, 7, 1, 0, 4, 3, 2},
		{7, 6, 5, 9, 8, 2, 1, 0, 4, 3},
		{8, 7, 6, 5, 9, 3, 2, 1, 0, 4},
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	}

	p = [][]int{
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{1, 5, 7, 6, 2, 8, 3, 0, 9, 4},
		{5, 8, 0, 3, 7, 9, 6, 1, 4, 2},
		{8, 9, 1, 6, 0, 4, 3, 5, 2, 7},
		{9, 4, 5, 3, 1, 2, 6, 8, 7, 0},
		{4, 2, 8, 6, 5, 7, 3, 9, 0, 1},
		{2, 7, 9, 3, 8, 0, 6, 4, 1, 5},
		{7, 0, 4, 6, 9, 1, 3, 2, 5, 8},
	}

	inv = []int{0, 4, 3, 2, 1, 5, 6, 7, 8, 9}
)

// ValidateVerhoeff validates Verhoeff checksum for given 12 digit Aadhar number 
func ValidateVerhoeff(numStr string) (valid bool) {
	// 12 digit number check
	matches, err := regexp.Match(aadhar, []byte(numStr))

	if err != nil || !matches {
		return
	}

	// Verhoeff checksum validation
	c := 0
	for i, num := range numStrToReverseIntArray(numStr) {
		c = d[c][p[(i % 8)][num]];
	}
	valid = (c == 0)
	return
}

func numStrToReverseIntArray(numStr string) (result []int) {
	result = make([]int, len(numStr))
	j := len(numStr) - 1
	for _, r := range numStr {
		result[j] = int(r - '0')
		j--
	}
	return
}
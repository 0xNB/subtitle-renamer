// Package analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

import (
	"fmt"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// LcSubstr finds the longest common substring in a list of strings given to it
func LcSubstr(strings []string) string {
	return ""
}

// Max returns the larger of x or y.
func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}


// LcSubStr computes the longest common substring of two strings using dynamic programming
func LcSubStr(s1 string, s2 string) (string, int) {

	m := len(s1) + 1
	n := len(s2) + 1

	// create two dimensional matrix
	var LCSuff [][]int = make([][]int, m)
	for i := 0; i < m; i++ {
		LCSuff[i] = make([]int, n);
	}

	var result int = 0
	var resultIdx = 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				LCSuff[i][j] = 0;
			} else if s1[i-1] == s2[j-1] {
				LCSuff[i][j] = LCSuff[i-1][j-1] +1;
				if LCSuff[i][j] > result {
					result = Max(result, LCSuff[i][j])
					resultIdx = i
				}
			} else {
				LCSuff[i][j] = 0
			}

		}
	}

	fmt.Sprintf("result length is %d", result)
	asRunes := []rune(s1)
	// convert to go runes and return the slice of those runes
	return string(asRunes[resultIdx - result : resultIdx]), result

}

func deduceCommonName(files []AnalyzedFile) (bool, string) {
	// TOOD impl
	return false, ""
}

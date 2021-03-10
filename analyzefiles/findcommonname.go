// Package analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

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

// LcSubStr computes the longest common substring of two strings using dynamic programming
func LcSubStr(s1 string, s2 string) (string, int) {

	m := len(s1) + 1
	n := len(s2) + 1

	var LCSuff []int = make([]int, (m+1)*(n+1))
	var result int = 0

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			
		}
	}

	// convert to go runes and return the slice of those runes
	runes := []rune(s1)
	return string(runes[lastIndex : lastIndex+resultLen]), resultLen

}

func deduceCommonName(files []AnalyzedFile) (bool, string) {
	// TOOD impl
	return false, ""
}

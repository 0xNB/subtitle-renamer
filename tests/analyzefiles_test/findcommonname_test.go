// AnalyzeFiles_Test is the testing package corresponding to the main readfiles go class
package analyzefiles_test

import (
	"path/filepath"
	"runtime"
	"subtitle-renamer/analyzefiles"
	"testing"
)

func TestFindCommonNames(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testFileDir := filepath.Dir(filename)
	t.Logf("Current test dir: %s", testFileDir)

	cases := []struct {
		string1          string
		string2          string
		longestSubstring string
	}{
		{
			string1:          "this is a teststring 123",
			string2:          "another teststring 345",
			longestSubstring: " testtring ",
		},
	}

	for _, c := range cases {
		result := analyzefiles.LcSubStr(c.string1, c.string2)
		if result != c.longestSubstring {
			t.Errorf("Longest substring didn't match result! Got [%s] Expected [%s]", result, c.longestSubstring)
		}
	}
}

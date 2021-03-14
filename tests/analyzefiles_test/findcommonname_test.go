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
			longestSubstring: " teststring ",
		},
		{
			string1: "12345",
			string2: "67889",
			longestSubstring: "",
		},
	}

	for _, c := range cases {
		result, len := analyzefiles.LcSubStr(c.string1, c.string2)
		if result != c.longestSubstring {
			t.Errorf("Longest substring didn't match result! Got [%s] Expected [%s], length was %d", result, c.longestSubstring, len)
		}
	}
}

func TestFindCommonNamesWrapAround(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testFileDir := filepath.Dir(filename)
	t.Logf("Current test dir: %s", testFileDir)

	cases := []struct {
		string1          string
		string2          string
		longestSubstring string
	}{
		{
			string1:          "episode_name01goodshit",
			string2:          "episode_name12goodshit",
			longestSubstring: "goodshitepisode_name",
		},
		{
			string1: "12345",
			string2: "67889",
			longestSubstring: "",
		},
	}

	for _, c := range cases {
		result, len := analyzefiles.LcSubStrWrap(c.string1, c.string2)
		if result != c.longestSubstring {
			t.Errorf("Longest substring didn't match result! Got [%s] Expected [%s], length was %d", result, c.longestSubstring, len)
		}
	}
}

func TestDeduceCommonName(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testFileDir := filepath.Dir(filename)
	t.Logf("Current test dir: %s", testFileDir)

	cases := []struct {
		subfiles         []analyzefiles.SubFile
		longestSubstring  string
		shouldBeValid bool
	}{
		{
			subfiles: []analyzefiles.SubFile{
				analyzefiles.SubFile {
						Name: "this is episode zero",
						FullPath: "goodshitEpisode01",
						OSFileInfo: nil,
					},
				analyzefiles.SubFile {
					Name: "this is episode zero",
					FullPath: "goodshitEpisode02",
					OSFileInfo: nil,
				 },
			},
			longestSubstring: "goodshitEpisode",
			shouldBeValid: true,
		},
	}

	for _, c := range cases {
		valid, result := analyzefiles.DeduceCommonName(c.subfiles)
		if valid != c.shouldBeValid {
			t.Errorf("Longest substring didn't match result! Got [%s] Expected [%s], length was %d", valid, c.longestSubstring, result)
		}
	}
}


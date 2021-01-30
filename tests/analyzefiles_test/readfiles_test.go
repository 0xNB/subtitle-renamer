// AnalyzeFiles_Test is the testing package corresponding to the main readfiles go class
package analyzefiles_test

import (
	"path/filepath"
	"runtime"
	"subtitle-renamer/analyzefiles"
	"testing"
)

func containsString(strings []string, s string) bool {
	for _, cur := range strings {
		if cur == s {
			return true
		}
	}
	return false
}

func ContainsFileName(fileNames []string, a analyzefiles.AnalyzedFile, t *testing.T) {
	if !containsString(fileNames, a.FileName()) {
		t.Errorf("Couldn't find required filename %s!", a.FileName())
	}
}

func TestReadFiles(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testFileDir := filepath.Dir(filename)
	t.Logf("Current test dir: %s", testFileDir)

	cases := []struct {
		inFolderName string
		fileNames    []string
	}{
		{
			inFolderName: "/demo_dir",
			fileNames:    []string{"terminator.srt", "test.mkv", "titanic.mkv"},
		},
	}
	for _, c := range cases {
		res := analyzefiles.ReadFilesFromDir(testFileDir + c.inFolderName)
		if len(res) != len(c.fileNames) {
			t.Error("Length of returned filenames doesn't match test length!")
		}
		for _, analyzedFile := range res {
			ContainsFileName(c.fileNames, analyzedFile, t)
		}
	}
}

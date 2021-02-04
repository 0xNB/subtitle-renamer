// AnalyzeFiles_Test is the testing package corresponding to the main readfiles go class
package analyzefiles_test

import (
	"subtitle-renamer/analyzefiles"
	"testing"
)

func TestReadFiles(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testFileDir := filepath.Dir(filename)
	t.Logf("Current test dir: %s", testFileDir)

	cases := []struct {
		inFolderName string
		fileInfos    []fileInfo
	}{
		{
			inFolderName: "/demo_dir",
			fileInfos: []fileInfo{
				{
					fileName: "terminator.srt",
					fileType: analyzefiles.SUB,
				},
				{fileName: "test.mkv",
					fileType: analyzefiles.VID,
				},
				{fileName: "titanic.mkv",
					fileType: analyzefiles.VID,
				},
			},
		},
		{
			inFolderName: "/demo_dir_unknown",
			fileInfos: []fileInfo{
				{
					fileName: "",
					fileType: analyzefiles.NIL,
				},
				{
					fileName: "",
					fileType: analyzefiles.NIL,
				},
			},
		},
	}

	for _, c := range cases {
		res := analyzefiles.ReadFilesFromDir(testFileDir + c.inFolderName)
		if len(res) != len(c.fileInfos) {
			t.Fatal("Length of returned fileInfo doesn't match test length!")
		}
		for _, analyzedFile := range res {
			EqualsNameAndFileType(c.fileInfos, analyzedFile, t)
			t.Logf("Found file %s, with corret type", analyzedFile.FileName())
		}
	}
}

// The analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

import (
	"io/ioutil"
	"log"
	"os"
)

const (
	MKV = iota
	ASS = iota
	MP4 = iota
)

// SubFile is the 
type SubFile struct {
	name string
	fileInfo os.FileInfo
}

// SubFile can return its own type, the underlying file or name
type AnalyzedFile interface {
	FileName() string
	FileType() string
	FileInfo() os.FileInfo
}

// ScanFiles takes an absolute path or relative path and outputs a
func ScanFiles(folderPath string) {

	scannedFiles := make([]SubFile, 0, 100)

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Couldn't read directory", err)
	}

	for _, file := range files {
		append(scannedFiles, {

		})
	}
}

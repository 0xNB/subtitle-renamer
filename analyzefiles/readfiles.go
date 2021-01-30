// The analyzefiles package gives use some utility methods to read files from disk in the current folder
// and determine their file type, sorting them in categories defined by analyze files.
package analyzefiles

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"slice"
	"strings"
)

const (
	MKV = iota
	ASS = iota
	MP4 = iota
	SRT = iota
)

// AnalyzedFile can return its own type, the underlying file or name
type AnalyzedFile interface {
	FileName() string
	FileType() string
	FileInfo() os.FileInfo
}

func DetermineFileType(info os.FileInfo) AnalyzedFile {
	
	fileExt := strings.ToLower(filepath.Ext(name))
	
	switch { 
		case slice.ContainsString(SupportedVideoFileTypes(), fileExt): 
			// supported video file extension detected
			return VidFile{
				name: info.Name(),
				fileInfo: info,
			}
		case slice.ContainsString(SupportedSubFileTypes(), fileExt):
			// supported sub file extension detected
			return SubFile{
				name: info.Name(),
				fileInfo: info,
			}
		default 
	}	

	
}

// ScanFiles takes an absolute path or relative path and outputs a
func ScanFiles(folderPath string) {

	scannedFiles := make([]File, 0, 100)

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal("Couldn't read directory", err)
	}

	for _, file := range files {
		
		append(scannedFiles, {
			
		})
	}
}

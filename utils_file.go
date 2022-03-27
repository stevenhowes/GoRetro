package GoRetro

/*
 * --------------------
 * File Handler
 * --------------------
 * All IO should pass through this rather than direct file access to allow
 * the use of archive files etc in future.
 */

import (
	"fmt"
	"os"
)

type vFile struct {
	Data []byte
	Size int
}

var FileList map[string]*vFile

func GetFile(filename string) (*vFile, error) {
	if val, ok := FileList[filename]; ok {
		CacheHitsFile++
		return val, nil
	}

	Data, err := os.ReadFile(Config.DataDir + filename)
	vf := vFile{
		Size: len(Data),
		Data: Data,
	}

	fmt.Printf("File Caching %s\n", filename)
	FileList[filename] = &vf
	return &vf, err
}

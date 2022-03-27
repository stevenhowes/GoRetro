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

	"github.com/stevenhowes/PakGo"
)

type vFile struct {
	Data []byte
	Size int
}

var FileList map[string]*vFile

var pak PakGo.PakFile

func PakLoad(filename string) error {
	var err error

	fmt.Printf("Loading PAK %s\n", filename)

	pak, err = PakGo.PakLoad(Config.DataDir + filename)
	if err != nil {
		panic(err)
	}

	return err
}

func GetFile(filename string) (*vFile, error) {
	if val, ok := FileList[filename]; ok {
		CacheHitsFile++
		return val, nil
	}

	Data, err := os.ReadFile(Config.DataDir + filename)
	if err != nil {
		Data, err = pak.ReadFile(filename)
	}

	vf := vFile{
		Size: len(Data),
		Data: Data,
	}

	fmt.Printf("File Caching %s at %d bytes\n", filename, len(Data))

	FileList[filename] = &vf
	return &vf, err
}

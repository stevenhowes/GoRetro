package GoRetro

/*
 * --------------------
 * File Handler
 * --------------------
 * All IO should pass through this rather than direct file access to allow
 * the use of archive files etc in future.
 */

import "os"

type vFile struct {
	Data []byte
	Size int
}

func GetFile(filename string) *vFile {
	Data, _ := os.ReadFile(Config.DataDir + filename)
	vf := vFile{
		Size: len(Data),
		Data: Data,
	}
	return &vf
}

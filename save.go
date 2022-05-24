package filemanager

import "io/fs"

/*
We should have functions that:
- Allow us to read and save a file in chunks, the reading will not be performed instantly in the memory, but it will be read
  in chunks... and also save in chunks!
- Allow us to save files from memory
*/

type SaveFileOptions struct {
	// Name -> is optional, if not indicated, the file name will be taken
	Name string
	// Save in database -> in chunks
	SaveInDB    bool
	DBChunkSize int64
	// by indicating full file path, it will be read from there!
	FullFilePath string
	// by indicating File, it will be read from there
	File fs.File
}

// Save -> saves the input file to the location (destination)
func (f *File) Save() {

}

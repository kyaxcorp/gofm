package filemanager

// NewFile -> create a new file in the DB
func (fm *FileManager) NewFile() *File {
	return &File{
		// Set the file manager as reference
		fileManager: fm,
	}
}

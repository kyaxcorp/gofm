package gofm

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FindFileOptions struct {
	ID   uuid.UUID
	Name string
}

// FindFile -> it will search for the described file in the database
// if nothing is found, an error will be returned
func (fm *FileManager) FindFile(o FindFileOptions) (*File, error) {
	var file *File
	dbResult := fm.db().Where(o).First(file)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return nil, fm.setError(ErrFmFileNotFound, dbResult.Error)
		} else {
			return nil, fm.setError(ErrFmDatabaseClientError, dbResult.Error)
		}
	}
	return file, nil
}

type FindFilesOptions struct {
	ID   uuid.UUID
	Name string
}

// FindFiles -> TODO:
func (fm *FileManager) FindFiles(o FindFilesOptions) *File {
	var file File
	fm.db().Where(o).First(&file)
	return nil
}

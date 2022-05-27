package filemanager

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FindOptions struct {
	ID   uuid.UUID
	Name string
}

// setError -> return the fmError for code simplicity
func (fm *FileManager) setError(fmError, intError error) error {
	fm.fmError = fmError
	fm.internalError = intError
	return fm.fmError
}

// FindFile -> it will search for the described file in the database
// if nothing is found, an error will be returned
func (fm *FileManager) FindFile(o FindOptions) (*File, error) {
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

// FindFiles -> TODO:
func (fm *FileManager) FindFiles(o FindOptions) *File {
	var file File
	fm.db().Where(o).First(&file)
	return nil
}

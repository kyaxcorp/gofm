package gofm

import (
	"errors"
	"gorm.io/gorm"
)

type FindFileOptions struct {
	//ID   uuid.UUID
	ID   UUID
	Name string
}

// FindFile -> it will search for the described file in the database
// if nothing is found, an error will be returned
func (fm *FileManager) FindFile(o FindFileOptions) (*File, error) {
	var file File
	dbResult := fm.db().
		Where(o).
		Where("fm_instance = ?", fm.Name).
		First(&file)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return nil, fm.setError(ErrFmFileNotFound, dbResult.Error)
		} else {
			return nil, fm.setError(ErrFmDBClientQueryFailed, dbResult.Error)
		}
	}
	return &file, nil
}

type FindFilesOptions struct {
	ID   UUID
	Name string
}

// FindFiles -> TODO:
func (fm *FileManager) FindFiles(o FindFilesOptions) *File {
	var file File
	fm.db().Where(o).First(&file)
	return nil
}

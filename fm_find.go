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
// this function may be deprecated because retrieving files can be done natively from the gorm from the app side
// where multiple parameters can be set by the user
func (fm *FileManager) FindFile(o FindFileOptions) (*File, error) {
	var file File
	dbResult := fm.DB().
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

func (fm *FileManager) StartFind(o FindFileOptions) *gorm.DB {
	return fm.DB().Where("fm_instance = ?", fm.Name)
}

type FindFilesOptions struct {
	ID   UUID
	Name string
}

// FindFiles -> TODO:
func (fm *FileManager) FindFiles(o FindFilesOptions) *File {
	var file File
	fm.DB().Where(o).First(&file)
	return nil
}

package gofm

import (
	"errors"
	gofileDriver "github.com/kyaxcorp/gofile/driver"
	"gorm.io/gorm"
	"reflect"
)

// FindFile -> it will search for the described file in the database
// if nothing is found, an error will be returned
func (fm *FileManager) FindFile(fileModel interface{}, cb func(db *gorm.DB) *gorm.DB) error {
	fileModelVal := reflect.ValueOf(fileModel)

	// first of all, check if the FileModel is a pointer of a struct!
	if fileModelVal.Kind() != reflect.Ptr {
		return ErrFileModelShouldBePointerOfAStruct
	}
	fieldModelIndirect := reflect.Indirect(fileModelVal)

	db := fm.DB().Where("fm_instance = ?", fm.Name)

	// Exec the Callback
	db = cb(db)

	dbResult := db.
		First(fileModel)
	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return fm.setError(ErrFmFileNotFound, dbResult.Error)
		} else {
			return fm.setError(ErrFmDBClientQueryFailed, dbResult.Error)
		}
	}
	// set the fileManager back to the fileModel!
	// why it needs it? because it will call some functions of reading the files from the locations
	if _model, ok := fileModel.(interface{ SetFileManager(fm *FileManager) }); ok {
		_model.SetFileManager(fm)
	}

	locations, _err := structGetFieldVal(fieldModelIndirect, "Locations")
	if _err != nil {
		return _err
	}
	locationsMeta := locations.(FileLocationsMeta)
	// let's also index the current locations for faster finding...
	// the fileModel it's a PTR
	locationsIndexed := make(map[LocationName]FileLocationMeta)
	filesIndexed := make(map[LocationName]gofileDriver.FileInterface)
	var defaultFile gofileDriver.FileInterface

	for _, locMeta := range locationsMeta {
		// let's get the location from the file manager instance

		if loc, ok := fm.LocationsIndexed[locMeta.LocationName]; ok {
			file, _err := loc.Driver.FileFromInfo(locMeta.FileInfo)
			// TODO: should we return  the error?! or we simply continue!
			if _err != nil {
				// We continue because the location that was provided before doesn't exist in the file manager
				continue
			}
			if defaultFile == nil {
				defaultFile = file
			}

			filesIndexed[LocationName(locMeta.LocationName)] = file
		}

		// Index the location
		locationsIndexed[LocationName(locMeta.LocationName)] = locMeta
	}

	if _model, ok := fileModel.(interface {
		SetLocationsIndexed(locations map[LocationName]FileLocationMeta)
	}); ok {
		_model.SetLocationsIndexed(locationsIndexed)
	}

	if _model, ok := fileModel.(interface {
		SetFilesIndexed(files map[LocationName]gofileDriver.FileInterface)
	}); ok {
		_model.SetFilesIndexed(filesIndexed)
	}

	// set the default file
	if _model, ok := fileModel.(interface {
		SetDefaultFile(file gofileDriver.FileInterface)
	}); ok {
		_model.SetDefaultFile(defaultFile)
	}

	// ar fi bine aici sa intoarcem si locatiile ca interfete pentru ca functiile sa se cheme direct...

	// TODO: we should test this...
	//structSetFieldVal(fieldModelIndirect,"locationsIndexed",locationsIndexed)

	return nil
}

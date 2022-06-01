package gofm

import (
	"errors"
	"github.com/google/uuid"
	"github.com/kyaxcorp/gofile/driver"
	"github.com/kyaxcorp/gofile/driver/filesystem"
	"github.com/kyaxcorp/gofile/driver/filesystem/helper"
	"os"
	"path/filepath"
	"reflect"
)

/*
We should have functions that:
- Allow us to read and save a file in chunks, the reading will not be performed instantly in the memory, but it will be read
  in chunks... and also save in chunks!
- Allow us to save files from memory
*/

// Save -> saves the input file to the location (destination)
func (f *NewFile) Save() error {

	fileModelVal := reflect.ValueOf(f.FileModel)

	// first of all, check if the FileModel is a pointer of a struct!
	if fileModelVal.Kind() != reflect.Ptr {
		return ErrFileModelShouldBePointerOfAStruct
	}

	fieldModelIndirect := reflect.Indirect(fileModelVal)

	currentLocation := filesystem.Location{DirPath: ""}
	var fileMetaData driver.FileInfoInterface
	var srcFile driver.FileInterface
	var _err error
	if f.InputFilePath != "" {
		//log.Println(f.InputFilePath)
		srcFile, _err = currentLocation.FindFile(f.InputFilePath)
	} else if f.GraphQLFile != nil {
		// save to a tmp path and then delete it
		graphFile := f.GraphQLFile
		fileData := make([]byte, graphFile.Size)
		// Read the file, ano now let's save it
		_, _err = graphFile.File.Read(fileData)
		if _err != nil {
			return _err
		}

		// Generate a random folder id (name)
		tmpFolderID, _err := uuid.NewRandom()
		if _err != nil {
			return _err
		}
		// create the temporary folder where to save the file
		tmpFolder := os.TempDir() + filepath.FromSlash("/") + tmpFolderID.String()
		if !helper.FolderExists(tmpFolder) {
			_err = helper.MkDir(tmpFolder, 0751)
			if _err != nil {
				return _err
			}
		}
		// Delete the folder and the files inside it after copying the destination locations
		defer helper.FolderDelete(tmpFolder)
		// Write the file to that generated folder
		tmpFileFullPath := tmpFolder + filepath.FromSlash("/") + graphFile.Filename
		_err = os.WriteFile(tmpFileFullPath, fileData, 0751)
		if _err != nil {
			return _err
		}
		srcFile, _err = currentLocation.FindFile(tmpFileFullPath)
	} else {
		return ErrFmNoInputFile
	}

	if _err != nil {
		return _err
	}
	fileMetaData = srcFile.Info()

	/*fileMetaData, _err := filesystem.NewFileInfo(f.InputFilePath)
	if _err != nil {
		return nil, _err
	}*/

	updatedAt := fileMetaData.UpdatedAt()
	createdAt := fileMetaData.CreatedAt()

	fileName := fileMetaData.Name()

	if !structFieldExists(fieldModelIndirect, "Name") {
		return errors.New("field Name doesn't exist in the file model")
	}

	// Check if the user has set any value for Name
	nameVal := fieldModelIndirect.FieldByName("Name").String()
	if nameVal != "" {
		fileName = nameVal
	} else {
		fileName = fileMetaData.Name()
	}

	// Set Name
	/*_err = structSetFieldVal(fieldModelIndirect, "Name", fileName)
	if _err != nil {
		return _err
	}*/

	// Generate a new file id
	id, _err := uuid.NewRandom()
	if _err != nil {
		return _err
	}

	//_err = structSetFieldVal(fieldModelIndirect, "ID", UUID(id))
	//if _err != nil {
	//	return _err
	//}

	structValues := map[string]interface{}{
		"ID":            UUID(id),
		"FMInstance":    f.fileManager.Name,
		"Name":          fileName,
		"FullName":      fileMetaData.FullName(),
		"OriginalName":  fileMetaData.FullName(),
		"Size":          fileMetaData.Size(),
		"Extension":     fileMetaData.Extension(),
		"ContentType":   fileMetaData.ContentType(),
		"FileUpdatedAt": &updatedAt,
		"FileCreatedAt": &createdAt,
	}

	// Set struct values
	for fieldName, fieldVal := range structValues {
		_err = structSetFieldVal(fieldModelIndirect, fieldName, fieldVal)
		if _err != nil {
			return _err
		}
	}

	// Set file manager
	if _model, ok := f.FileModel.(interface{ SetFileManager(fm *FileManager) }); ok {
		_model.SetFileManager(f.fileManager)
	}

	// TODO: call SetFileManager

	/*file := &File{
		ID:            UUID(id),
		FMInstance:    f.fileManager.Name,
		Name:          fileName,
		Description:   f.Description,
		CategoryID:    f.CategoryID,
		FullName:      fileMetaData.FullName(),
		OriginalName:  fileMetaData.FullName(),
		Size:          fileMetaData.Size(),
		Extension:     fileMetaData.Extension(),
		ContentType:   fileMetaData.ContentType(),
		FileUpdatedAt: &updatedAt,
		FileCreatedAt: &createdAt,

		// helpers
		fileManager: f.fileManager,
	}*/

	// we will first copy the files to the locations (destinations) and after that make an insert into the DB if success!

	if len(f.Locations) == 0 {
		return ErrFmNewFileLocationMissing
	}

	var fileLocationsMeta FileLocationsMeta

	locationFileName := id.String() + "_" + fileMetaData.FullName()

	// Copy in all mentioned locations
	for _, destLoc := range f.Locations {
		if loc, ok := f.fileManager.LocationsIndexed[destLoc.LocationName]; ok {
			//

			newFile, _err := loc.Driver.CopyFile(srcFile, driver.FileDestination{
				// the path may be empty...
				FileName: locationFileName,
				FilePath: destLoc.FilePath,
				DirPath:  destLoc.DirPath,
				// that contains path
			})
			if _err != nil {
				// Failed to copy...
				return _err
			}

			fileLocationsMeta = append(fileLocationsMeta, FileLocationMeta{
				LocationName: destLoc.LocationName,
				FileInfo:     newFile.Info().ToStruct(),
			})
		}
	}

	// in DB ar fi bine sa se salveze location instance name si + FileInfo
	// apoi o sa fie nevoie de restabilit Fisierul pe baza la FileInfo cu ajutorul functiei FindFile
	//

	_err = structSetFieldVal(fieldModelIndirect, "Locations", fileLocationsMeta)
	if _err != nil {
		return _err
	}
	//file.Locations = fileLocationsMeta

	//log.Println("saving file")
	dbResult := f.db().Save(f.FileModel)
	if dbResult.Error != nil {
		// TODO: set the error...
		return dbResult.Error
	}

	return nil
}

// TODO: we can do later a BackgroundSave which in case of failure (because of storage failure or interconnection failure)
// 		will retry the saving in specific locations
// BgSave -> background save
func (f *NewFile) BgSave() (*File, error) {
	return nil, nil
}

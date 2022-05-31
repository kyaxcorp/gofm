package gofm

import (
	"github.com/google/uuid"
	"github.com/kyaxcorp/gofile/driver"
	"github.com/kyaxcorp/gofile/driver/filesystem"
	"log"
)

/*
We should have functions that:
- Allow us to read and save a file in chunks, the reading will not be performed instantly in the memory, but it will be read
  in chunks... and also save in chunks!
- Allow us to save files from memory
*/

// Save -> saves the input file to the location (destination)
func (f *NewFile) Save() (*File, error) {
	currentLocation := filesystem.Location{DirPath: ""}

	log.Println(f.InputFilePath)
	srcFile, _err := currentLocation.FindFile(f.InputFilePath)
	if _err != nil {
		return nil, _err
	}
	fileMetaData := srcFile.Info()

	/*fileMetaData, _err := filesystem.NewFileInfo(f.InputFilePath)
	if _err != nil {
		return nil, _err
	}*/

	updatedAt := fileMetaData.UpdatedAt()
	createdAt := fileMetaData.CreatedAt()

	fileName := fileMetaData.Name()
	if f.Name != "" {
		fileName = f.Name
	}

	// Generate a new file id
	id, _err := uuid.NewRandom()
	if _err != nil {
		return nil, _err
	}

	file := &File{
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
	}

	// we will first copy the files to the locations (destinations) and after that make an insert into the db if success!

	if len(f.Locations) == 0 {
		return nil, ErrFmNewFileLocationMissing
	}

	var fileLocationsMeta FileLocationsMeta

	// Copy in all mentioned locations
	for _, destLoc := range f.Locations {
		if loc, ok := f.fileManager.LocationsIndexed[destLoc.LocationName]; ok {
			//

			newFile, _err := loc.Driver.CopyFile(srcFile, driver.FileDestination{
				// the path may be empty...
				FilePath: destLoc.FilePath,
				DirPath:  destLoc.DirPath,
				// that contains path
			})
			if _err != nil {
				// Failed to copy...
				return nil, _err
			}

			// let's get info about the new file
			// TODO: we should save

			fileLocationsMeta = append(fileLocationsMeta, FileLocationMeta{
				LocationName: destLoc.LocationName,
				FileInfo:     newFile.Info().ToStruct(),
			})
		}
	}

	// TODO: in DB ar fi bine sa se salveze location instance name si + FileInfo
	// TODO: apoi o sa fie nevoie de restabilit Fisierul pe baza la FileInfo cu ajutorul functiei FindFile
	//

	file.Locations = fileLocationsMeta

	log.Println("saving file")
	dbResult := f.db().Save(file)
	if dbResult.Error != nil {
		// TODO: set the error...
		return nil, dbResult.Error
	}

	return file, nil
}

// TODO: we can do later a BackgroundSave which in case of failure (because of storage failure or interconnection failure)
// 		will retry the saving in specific locations
// BgSave -> background save
func (f *NewFile) BgSave() (*File, error) {
	return nil, nil
}

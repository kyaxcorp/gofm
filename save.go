package filemanager

/*
We should have functions that:
- Allow us to read and save a file in chunks, the reading will not be performed instantly in the memory, but it will be read
  in chunks... and also save in chunks!
- Allow us to save files from memory
*/

// Save -> saves the input file to the location (destination)
func (f *NewFile) Save() (*File, error) {

	/*
		1. Detect content-type
		2. Set File Manager Instance Name
		3. Get the physical file!
		4. Check if everything is ok with the physical file
		5.
	*/

	fileMetaData, _err := GetFileMetaData(f.File)
	if _err != nil {
		return nil, _err
	}

	file := &File{
		FMInstance:   f.fileManager.Name,
		Name:         fileMetaData.Name,
		Description:  f.Description,
		CategoryID:   f.CategoryID,
		FullName:     fileMetaData.BaseName,
		OriginalName: fileMetaData.BaseName,
		Size:         fileMetaData.Size,
		Extension:    fileMetaData.Extension,
		ContentType:  fileMetaData.ContentType,
	}

	// 1. Save the file to the storage
	// 2. after saving the file to the storage should
	// Use Transaction!?
	// we should first generate a new id

	// we will first copy the files to the locations (destinations) and after that make an insert into the db if success!

	dbResult := f.db().Save(file)
	if dbResult.Error != nil {
		// TODO: set the error...
		return nil, dbResult.Error
	}

	// Now let's try saving the file
	// if failed, let's delete the inserted record in db!?
	// why don't we use transactions!? because transaction are slow -> if using for cluster, they'll be even slower...
	// 1.  we don't do so many operations over here...
	// 2. files can be quite large, and saving to the new location sometimes can take some time, but if starting the
	// 	  transaction

	return file, nil
}

// TODO: we can do later a BackgroundSave which in case of failure (because of storage failure or interconnection failure)
// 		will retry the saving in specific locations
// BgSave -> background save
func (f *NewFile) BgSave() (*File, error) {
	return nil, nil
}

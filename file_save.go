package gofm

func (f *File) Save() (*File, error) {
	dbResult := f.db().Save(f)
	if dbResult.Error != nil {
		return f, dbResult.Error
	}
	return f, nil
}

package gofm

import "gorm.io/gorm"

func (fm *FileManager) db() *gorm.DB {
	if fm.DBClient == nil {
		panic("DBClient is nil")
	}
	return fm.DBClient
}

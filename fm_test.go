package filemanager

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
)

func getTestDBClient() (*gorm.DB, error) {
	dsn := "host=172.29.24.104 user=filemanager password=filemanager dbname=filemanager port=26257 sslmode=require TimeZone=Europe/Chisinau"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})

}

func TestFileManager(t *testing.T) {
	log.Println("connecting to db...")

	db, _err := getTestDBClient()
	if _err != nil {
		t.Error(_err)
		return
	}

	fm := GetInstance(&FileManager{
		DBClient:      db,
		DBAutoMigrate: true,
		Name:          "instance name",
	})

	file := fm.NewFile()
	file.Name = "my super File"
	file.Save()
}

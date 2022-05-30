package gofm

import (
	"github.com/kyaxcorp/gofile/driver/filesystem"
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

	loc := Location{
		Name: "local",
		Driver: &filesystem.Location{
			DirPath: ".",
		},
	}

	db, _err := getTestDBClient()
	if _err != nil {
		t.Error(_err)
		return
	}

	// Define file manager locations
	fm := GetInstance(&FileManager{
		DBClient:      db,
		DBAutoMigrate: true,
		Name:          "instance name",
		Locations:     Locations{loc},
	})

	file := fm.NewFile()
	//file.Name = "my super InputFile"
	//file.Description = "my description"
	file.InputFilePath = "./LICENSE"
	file.Locations = []NewFileLocation{
		{
			// Define the location instance name
			LocationName: "local",
		},
	}
	file.Save()
}

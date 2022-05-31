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

	// define a storage location
	loc := Location{
		Name: "local",
		Driver: &filesystem.Location{
			DirPath: "./newpath/loc1",
		},
	}
	loc2 := Location{
		Name: "local2",
		Driver: &filesystem.Location{
			DirPath: "./newpath/loc2",
		},
	}

	// Get the db client
	db, _err := getTestDBClient()
	if _err != nil {
		t.Error(_err)
		return
	}

	// Define the file manager
	fm := GetInstance(&FileManager{
		DBClient:      db,
		DBAutoMigrate: true,
		Name:          "instance1",
		// Define file manager locations
		Locations: Locations{loc, loc2},
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
		{
			// Define the location instance name
			LocationName: "local2",
		},
	}
	f, _err := file.Save()
	if _err != nil {
		t.Error(_err)
		return
	}
	log.Println("id", f.ID)
}

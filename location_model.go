package gofm

import (
	"github.com/kyaxcorp/gofile/driver"
)

type Location struct {
	// this is used driver name
	//DriverType string
	// Name -> the name is the identity
	Name string

	// TODO: let's describe here how the files will be saved!
	// save as YYYY/mm/dd folders and UUID
	// save as YYYY/mm/dd folders and OriginalFileName
	// save as YYYY/mm/dd folders and OriginalFileName + Incremented one

	fileManager *FileManager `gorm:"-"`

	// the driver is the one that creates the bridge between the client and the destination server (location)
	// this is the instance with which the will interact
	Driver driver.LocationInterface
}

// TableName -> Get the Database table name from the file manager
func (f *Location) TableName() string {
	return ""
	//return f.fileManager.GetLocationsDBTableName()
}

/*
// Scan -> this is the func which is called when it's necessary to read from the Database to
// the defined variable with this type
func (l *Location) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	// Define a temporary variable
	var data Location
	// JSON Decode into the temporary variable
	err := json.Unmarshal(bytes, &data)
	// Check if it's ok...

	// we should set the data from DB to the existing structure... and should be dynamic?!...
	if err == nil {
		// if ok, set back the obtained value
		*l = data
	} else {
		*l = Location{}
	}
	return err
}

// Value -> this is the func which is called when it's necessary to send to Database,
// the value should be automatically converted to JSON
func (l Location) Value() (driver.Value, error) {
	// Convert the data into Json Bytes
	return json.Marshal(l)
}

func (Location) GormDataType() string {
	return "bytes"
}

// GormDBDataType gorm db data type
// This is the Database data type which is sent to the DB
func (Location) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
*/

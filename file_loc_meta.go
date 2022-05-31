package gofm

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	gofileDriver "github.com/kyaxcorp/gofile/driver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type FileLocationMeta struct {
	// this is the instance Name
	LocationName string
	FileInfo     gofileDriver.FileInfo
}

type FileLocationsMeta []FileLocationMeta

// Scan -> this is the func which is called when it's necessary to read from the Database to
// the defined variable with this type
func (l *FileLocationsMeta) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	// Define a temporary variable
	var data FileLocationsMeta
	// JSON Decode into the temporary variable
	err := json.Unmarshal(bytes, &data)
	// Check if it's ok...

	// we should set the data from DB to the existing structure... and should be dynamic?!...
	if err == nil {
		// if ok, set back the obtained value
		*l = data
	} else {
		*l = FileLocationsMeta{}
	}
	return err
}

// Value -> this is the func which is called when it's necessary to send to Database,
// the value should be automatically converted to JSON
func (l FileLocationsMeta) Value() (driver.Value, error) {
	// Convert the data into Json Bytes
	return json.Marshal(l)
}

func (FileLocationsMeta) GormDataType() string {
	return "bytes"
}

// GormDBDataType gorm db data type
// This is the Database data type which is sent to the DB
func (FileLocationsMeta) GormDBDataType(db *gorm.DB, field *schema.Field) string {
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

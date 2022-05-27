package gofm

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Locations []Location

// Scan -> this is the func which is called when it's necessary to read from the Database to
// the defined variable with this type
func (l *Locations) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	// Define a temporary variable
	var data Locations
	// JSON Decode into the temporary variable
	err := json.Unmarshal(bytes, &data)
	// Check if it's ok...

	// we should set the data from DB to the existing structure... and should be dynamic?!...
	if err == nil {
		// if ok, set back the obtained value
		*l = data
	} else {
		*l = Locations{}
	}
	return err
}

// Value -> this is the func which is called when it's necessary to send to Database,
// the value should be automatically converted to JSON
func (l Locations) Value() (driver.Value, error) {
	// Convert the data into Json Bytes
	return json.Marshal(l)
}

func (Locations) GormDataType() string {
	return "bytes"
}

// GormDBDataType gorm db data type
// This is the Database data type which is sent to the DB
func (Locations) GormDBDataType(db *gorm.DB, field *schema.Field) string {
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

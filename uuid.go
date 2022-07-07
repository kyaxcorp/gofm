package gofm

import (
	"database/sql/driver"
	"fmt"
	"github.com/google/uuid"
)

type UUID uuid.UUID

// Scan implements sql.Scanner so UUIDs can be read from databases transparently.
// Currently, database types that map to string and []byte are supported. Please
// consult database-specific driver documentation for matching types.
func (u *UUID) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case string:
		// if an empty UUID comes from a table, we return a null UUID
		if src == "" {
			return nil
		}

		// see Parse for required string format
		pu, err := uuid.Parse(src)
		if err != nil {
			return fmt.Errorf("Scan: %v", err)
		}

		*u = UUID(pu)

	case []byte:
		// if an empty UUID comes from a table, we return a null UUID
		if len(src) == 0 {
			return nil
		}

		// assumes a simple slice of bytes if 16 bytes
		// otherwise attempts to parse
		if len(src) != 16 {
			return u.Scan(string(src))
		}
		copy((*u)[:], src)

	default:
		return fmt.Errorf("Scan: unable to scan type %T into UUID", src)
	}

	return nil
}

func (u *UUID) UnmarshalJSON(b []byte) error {
	s := string(b)
	_u, _err := uuid.Parse(s)
	if _err != nil {
		return _err
	}
	*u = UUID(_u)
	return nil
}

// Value implements sql.Valuer so that UUIDs can be written to databases
// transparently. Currently, UUIDs map to strings. Please consult
// database-specific driver documentation for matching types.
func (u UUID) Value() (driver.Value, error) {
	return uuid.UUID(u).String(), nil
}

func (u UUID) String() string {
	return uuid.UUID(u).String()
}

func (u UUID) URN() string {
	return uuid.UUID(u).URN()
}

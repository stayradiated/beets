package beets

import (
	"database/sql"
	"fmt"
)

type JSONNullString struct {
	sql.NullString
}

func (j JSONNullString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", j.String)), nil
}

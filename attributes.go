package beets

import (
	"database/sql"
)

type Attribute struct {
	Key   string
	Value string
}

type AttributeList []Attribute

const attributeColumns = `
	key, value
`

func (a *Attribute) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&a.Key, &a.Value,
	)
}

func (a AttributeList) Get(key string) string {
	for _, attribute := range a {
		if attribute.Key == key {
			return attribute.Value
		}
	}
	return ""
}

func ParseRowsAsAttributes(rows *sql.Rows) (AttributeList, error) {
	attributes := make(AttributeList, 0)
	for rows.Next() {
		attribute := Attribute{}

		if err := attribute.Scan(rows); err != nil {
			return nil, err
		}

		attributes = append(attributes, attribute)
	}
	return attributes, nil
}

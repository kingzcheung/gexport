package gexport

import "strings"

type FieldTyper interface {
	FieldType(in string) string
}

var typStrings = []string{"char", "varchar", "tinytext", "text", "mediumtext", "longtext"}
var typDates = []string{"date", "time", "datetime", "timestamp"}

func (s *SqlStruct) FieldType(in string) string {
	in = strings.ToLower(in)
	if strings.HasPrefix(in, "tinyint") {
		return "int8"
	}
	if strings.HasPrefix(in, "smallint") {
		return "int16"
	}
	if strings.HasPrefix(in, "mediumint") {
		return "int"
	}
	if strings.HasPrefix(in, "int") {
		return "int"
	}

	if strings.HasPrefix(in, "bigint") {
		return "int64"
	}

	for _, typString := range typStrings {
		if strings.HasPrefix(in, typString) {
			return "string"
		}
	}
	for _, typDate := range typDates {
		if strings.HasPrefix(in, typDate) {
			return "time.Time"
		}
	}

	return "string"
}

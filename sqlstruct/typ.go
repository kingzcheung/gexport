package sqlstruct

import "strings"

type FieldTyper interface {
	FieldType(in string) string
}

var typStrings = []string{"char", "varchar", "tinytext", "text", "mediumtext", "longtext"}
var typDates = []string{"date", "time", "datetime", "timestamp"}

func (s *SqlStruct) FieldType(in string) string {
	in = strings.ToLower(in)
	var out string
	if strings.HasPrefix(in, "tinyint") {
		out = "int8"
	}
	if strings.HasPrefix(in, "smallint") {
		out = "int16"
	}
	if strings.HasPrefix(in, "mediumint") {
		out = "int"
	}
	if strings.HasPrefix(in, "int") {
		out = "int"
	}

	if strings.HasPrefix(in, "bigint") {
		out = "int64"
	}

	if strings.Index(in, "unsigned") > 0 {
		out = "u" + out
	}

	for _, typString := range typStrings {
		if strings.HasPrefix(in, typString) {
			out = "string"
		}
	}
	for _, typDate := range typDates {
		if strings.HasPrefix(in, typDate) {
			out = "time.Time"
		}
	}

	return out
}

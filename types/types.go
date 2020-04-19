package types

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	//数值类型
	TINYINT   = "TINYINT"
	SMALLINT  = "SMALLINT"
	MEDIUMINT = "MEDIUMINT"
	INT       = "INT"
	INTEGER   = "INTEGER"
	BIGINT    = "BIGINT"
	FLOAT     = "FLOAT"
	DOUBLE    = "DOUBLE"
	DECIMAL   = "DECIMAL"

	//日期和时间类型
	DATE      = "DATE"
	TIME      = "TIME"
	YEAR      = "YEAR"
	DATETIME  = "DATETIME"
	TIMESTAMP = "TIMESTAMP"

	//字符串类型
	CHAR       = "CHAR"
	VARCHAR    = "VARCHAR"
	TINYBLOB   = "TINYBLOB"
	TINYTEXT   = "TINYTEXT"
	BLOB       = "BLOB"
	TEXT       = "TEXT"
	MEDIUMBLOB = "MEDIUMBLOB"
	MEDIUMTEXT = "MEDIUMTEXT"
	LONGBLOB   = "LONGBLOB"
	LONGTEXT   = "LONGTEXT"
)

func GoType(fieldType string) string {
	fmt.Println(fieldType)
	var goT bytes.Buffer

	ft := strings.Split(fieldType, " ")

	if len(ft) == 2 {
		goT.WriteString("u")
	}
	fieldWithLen := ft[0]
	field := strings.Split(fieldWithLen, "(")[0]

	switch strings.ToUpper(field) {
	case TINYINT:
		goT.WriteString("int8")
	case SMALLINT:
		goT.WriteString("int16")
	case MEDIUMINT, INT, INTEGER:
		goT.WriteString("int")
	case BIGINT:
		goT.WriteString("int64")
	case FLOAT, DECIMAL, DOUBLE:
		goT.WriteString("int64")
	case DATE, TIME, YEAR, DATETIME, TIMESTAMP:
		goT.WriteString("time.Time")
	case CHAR, VARCHAR, TINYTEXT, LONGTEXT, TEXT, MEDIUMTEXT:
		goT.WriteString("string")
	case BLOB, TINYBLOB, MEDIUMBLOB, LONGBLOB:
		//暂时
		goT.WriteString("string")
	default:
		goT.WriteString(fieldType)
	}

	return goT.String()
}

package gexport

import (
	"fmt"
	"github.com/kingzcheung/gexport/types"
	"github.com/xwb1989/sqlparser"
)

type Sql struct {
	structName string
}

func (s *Sql) SetStructName(structName string) {
	s.structName = structName
}

func NewSql() *Sql {
	return &Sql{}
}

func (s *Sql) Parse(sql string) ([]string, error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return nil, err
	}
	createStmt, ok := stmt.(*sqlparser.DDL)
	if !ok {
		return nil, fmt.Errorf("sql error")
	}
	res, err := s.parseCreateSql(createStmt)
	return []string{res}, err
}

func (s *Sql) parseCreateSql(stmt *sqlparser.DDL) (string, error) {

	var name = stmt.NewName.Name.String()

	if s.structName != "" {
		name = s.structName
	}
	gs := NewGoStruct(name)
	gs.Start()
	for _, col := range stmt.TableSpec.Columns {
		var (
			tags      []Tag
			fieldName string
		)
		fieldName = col.Name.String()
		// 添加json标签
		tags = append(tags, CreateJsonTag(fieldName))
		// 添加form标签
		tags = append(tags, CreateFormTag(fieldName))

		// gorm tag
		field := map[string]string{
			"column": fieldName,
			"type":   reformatType(col.Type.Type, col.Type.Length),
		}
		if col.Type.Autoincrement {
			field["autoIncrement"] = ""
		}

		tags = append(tags, Tag{
			Name:  "gorm",
			Field: field,
		})

		gs.Field(col.Name.String(), types.GoType(col.Type.Type), tags...)
	}
	gs.End()
	gs.WithTableFunc()
	return gs.String(), nil
}

func reformatType(t string, v *sqlparser.SQLVal) string {
	if t == "varchar" {
		return fmt.Sprintf("%s(%s)", t, string(v.Val))
	}
	return t
}

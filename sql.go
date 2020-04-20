package gexport

import (
	"fmt"
	"github.com/kingzcheung/gexport/types"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"strings"
)

type Sql struct {
	parser     *parser.Parser
	structName string
}

func (s *Sql) SetStructName(structName string) {
	s.structName = structName
}

func NewSql() *Sql {
	return &Sql{parser: parser.New()}
}

func (s *Sql) Parse(sql string) ([]string, error) {
	stmts, _, err := s.parser.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	var res []string

	if strings.HasPrefix(strings.ToLower(sql), "create") {
		res, err = s.parseCreateSql(stmts)
	}

	return res, err
}

func (s *Sql) parseCreateSql(stmts []ast.StmtNode) ([]string, error) {
	var structRes []string
	for _, stmt := range stmts {
		sc, ok := stmt.(*ast.CreateTableStmt)
		if !ok {
			return nil, fmt.Errorf("SQL 错误")
		}
		var name = sc.Table.Name.String()

		if s.structName != "" {
			name = s.structName
		}
		gs := NewGoStruct(name)
		gs.Start()
		for _, col := range sc.Cols {
			var (
				tags      []Tag
				fieldName string
			)
			fieldName = col.Name.String()
			//添加json标签
			tags = append(tags, CreateJsonTag(fieldName))
			//添加form标签
			tags = append(tags, CreateFormTag(fieldName))

			//gorm tag
			tags = append(tags, Tag{
				Name: "gorm",
				Fields: map[string]string{
					"column": fieldName,
					"type":   col.Tp.String(),
				},
			})
			gs.Field(col.Name.String(), types.GoType(col.Tp.InfoSchemaStr()), tags...)
		}
		gs.End()
		structRes = append(structRes, gs.String())
	}
	return structRes, nil
}

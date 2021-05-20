package gexport

import (
	"bytes"
	"fmt"
	"github.com/kingzcheung/gexport/tmpl"
	"text/template"
)

type Typ int

const (
	SQL Typ = iota
	JSON
)

type GExport struct {
	parse Parser
	sql   string
}

func New(sql string, typ Typ) *GExport {
	ge := new(GExport)
	ge.sql = sql
	switch typ {
	case SQL:
		s := NewSql()
		s.SetHasJson(true)
		s.SetHasXml(true)
		s.SetHasGorm(true)
		ge.parse = s

	default:
		panic("未实现")
	}
	return ge
}

func (e *GExport) Export() ([]byte, error) {

	parse, err := e.parse.Parse(e.sql)
	if err != nil {
		return nil, err
	}

	t, err := template.ParseFS(tmpl.TmpFs, "struct.tmpl")
	if err != nil {
		return nil, err
	}

	var bf = bytes.NewBuffer(nil)

	err = t.ExecuteTemplate(bf, "struct.tmpl", map[string]interface{}{
		"structName": parse.StructName,
		"Fields":     parse.Fields,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(bf.String())

	return nil, nil
}

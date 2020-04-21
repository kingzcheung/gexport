package gexport

import "fmt"

type ExportType string

const (
	SQL  ExportType = "sql"
	JSON ExportType = "json"
)

type Gexport struct {
	//The original string can be json or sql
	raw string
	//Parser from https://github.com/pingcap/parser
	parser     StructParser
	output     []string
	err        error
	StructName string
}

func (g *Gexport) Output() []string {
	return g.output
}

func New(raw string, t ...ExportType) *Gexport {
	var exportType ExportType = "sql"
	if len(t) > 0 {
		exportType = t[0]
	}
	gx := &Gexport{raw: raw}
	gx.newParser(exportType)
	return gx
}

func (g *Gexport) newParser(t ExportType) {
	switch t {
	case SQL:
		g.parser = NewSql()

	case JSON:
		g.parser = NewJson()
	default:
		g.parser = NewSql()
	}
}

func (g *Gexport) Parse() *Gexport {
	g.parser.SetStructName(g.StructName)
	g.output, g.err = g.parser.Parse(g.raw)
	if len(g.output) == 0 {
		g.err = fmt.Errorf("data cannot be parsed")
	}
	return g
}

func (g *Gexport) Error() error {
	return g.err
}

func (g *Gexport) String() string {
	if len(g.output) == 0 {
		return ""
	}
	return g.output[0]
}

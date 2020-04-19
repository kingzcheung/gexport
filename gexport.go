package gexport

type ExportType string

const (
	SQL  ExportType = "sql"
	JSON ExportType = "json"
)

type Gexport struct {
	raw    string
	parser Parser
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

func (g *Gexport) String() string {
	res, _ := g.parser.Parse(g.raw)
	return res[0]
}

package gexport

type StructParser interface {
	Parse(raw string) ([]string, error)
	SetStructName(structName string)
}

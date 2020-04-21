package gexport

type StructParser interface {
	//Parse the input data and return the structure list
	Parse(raw string) ([]string, error)
	//Set the structure name
	SetStructName(structName string)
}

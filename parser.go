package gexport

type Parser interface {
	// Parse the input data and return the structure list
	Parse(raw string) (*Struct, error)
}

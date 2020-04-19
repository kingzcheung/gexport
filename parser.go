package gexport

type Parser interface {
	Parse(raw string) ([]string, error)
}

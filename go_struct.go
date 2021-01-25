package gexport

import (
	"bytes"
	"fmt"
	"sort"
)

var DefaultStructName = "RootGeneratedName"

type GoStruct struct {
	Name    string
	rawName string
	buf     bytes.Buffer
	// 是否内联
	IsInline bool
}

func NewGoStruct(name ...string) *GoStruct {
	var gs = new(GoStruct)
	if len(name) > 0 {
		gs.Name = NameCamelCase(name[0])
		gs.rawName = name[0]
	} else {
		gs.Name = DefaultStructName
	}
	return gs
}

func (g *GoStruct) Start() {
	g.buf.WriteString("struct {\n")
}

func (g *GoStruct) WithTableFunc() {
	if g.rawName == "" {
		return
	}

	g.buf.WriteString(`func (`)
	g.buf.WriteString(g.Name)
	g.buf.WriteString(") TableName() string {\n")
	g.buf.WriteString(`return "`)
	g.buf.WriteString(g.rawName)
	g.buf.WriteString("\"\n}")
}

// Field return struct field.
func (g *GoStruct) Field(name string, fileType string, tags ...Tag) {
	g.buf.WriteString("\t")
	g.buf.WriteString(NameCamelCase(name))
	g.buf.WriteString(" ")
	g.buf.WriteString(fileType)
	g.buf.WriteString(" ")

	// add tags
	if len(tags) > 0 {
		g.buf.WriteString("`")
		for _, tag := range tags {
			g.buf.WriteString(tag.Name)
			g.buf.WriteString(":")
			g.buf.WriteString("\"")
			fLen := len(tag.Field)
			var i int
			// 保证顺序输出
			var keys []string
			for k := range tag.Field {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, key := range keys {
				g.buf.WriteString(key)
				if tag.Field[key] != "" {
					g.buf.WriteString(":")
					g.buf.WriteString(tag.Field[key])
				}
				if i < fLen-1 {
					g.buf.WriteString(";")
				}
				i++
			}
			g.buf.WriteString("\"")
			g.buf.WriteString(" ")
		}
		g.buf.WriteString("`")
	}

	g.buf.WriteString("\n")
}

func (g *GoStruct) End() {
	g.buf.WriteString("}\n")
}

func (g *GoStruct) TName() string {
	return fmt.Sprintf("type %s ", g.Name)
}
func TName(name string) string {
	return fmt.Sprintf("type %s ", name)
}
func (g *GoStruct) String() string {
	return g.TName() + g.buf.String()
}

func (g *GoStruct) StringNotType() string {
	return g.buf.String()
}

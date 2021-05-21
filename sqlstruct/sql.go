package sqlstruct

import (
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/parser/test_driver"
)

type SqlStruct struct {
	structName string
	hasJson    bool
	hasGorm    bool
	hasXml     bool
}

func (s *SqlStruct) SetStructName(structName string) {
	s.structName = structName
}

func (s *SqlStruct) SetHasJson(hasJson bool) {
	s.hasJson = hasJson
}

func (s *SqlStruct) SetHasGorm(hasGorm bool) {
	s.hasGorm = hasGorm
}

func (s *SqlStruct) SetHasXml(hasXml bool) {
	s.hasXml = hasXml
}

func NewSql() *SqlStruct {
	return &SqlStruct{}
}

func (s *SqlStruct) Parse(sql string) (*Struct, error) {
	p := parser.New()
	stmt, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		return nil, err
	}
	ct, ok := stmt.(*ast.CreateTableStmt)
	if !ok {
		return nil, fmt.Errorf("it is not createtable sql: %v", sql)
	}

	st := new(Struct)
	st.StructName = s.FieldName(ct.Table.Name.String())
	for _, col := range ct.Cols {
		sf := new(StructField)
		sf.FieldName = s.FieldName(col.Name.String())
		sf.FieldType = s.FieldType(col.Tp.String())
		//fmt.Printf("%+v\n", col.Tp)

		if s.hasJson {
			sf.Tags = append(sf.Tags, &Tag{
				TagName: "json",
				TagValue: map[string]string{
					col.Name.String(): "",
				},
			})
		}

		if s.hasXml {
			sf.Tags = append(sf.Tags, &Tag{
				TagName: "xml",
				TagValue: map[string]string{
					col.Name.String(): "",
				},
			})
		}
		if s.hasGorm {
			sf.Tags = append(sf.Tags, &Tag{
				TagName: "gorm",
				TagValue: map[string]string{
					"column": col.Name.String(),
				},
			})
		}

		//for _, option := range col.Options {
		//
		//	fmt.Printf("%+v\n", option)
		//	if option.Expr != nil {
		//		fmt.Println(option.Expr.Text())
		//	}
		//}
		st.Fields = append(st.Fields, sf)
	}
	//fmt.Printf("%+v\n",ct.Cols[0])
	//fmt.Printf("%+v\n",ct.Table.Name.String())

	return st, nil
}

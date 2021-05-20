package gexport

import (
	"fmt"
	"github.com/kingzcheung/gexport/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	src string
	ok  bool
	out string
}

func TestSql_Parse(t *testing.T) {
	data, _ := testdata.TestData.ReadFile("simple3.sql")
	ns := NewSql()
	ns.hasJson = true
	res, err := ns.Parse(string(data))
	if err != nil {
		assert.Error(t, err)
	}

	for _, field := range res.Fields {
		fmt.Printf("%+v\n", field)

	}

}

package gexport

import (
	"github.com/kingzcheung/gexport/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	src string
	ok  bool
	out string
}

func TestSqlStruct_Parse2(t *testing.T) {
	readFile, err := testdata.TestData.ReadFile("simple1.sql")
	assert.NoError(t, err)
	s := NewSql()
	_, _ = s.Parse(string(readFile))
}

func TestSqlStruct_Parse(t *testing.T) {
	type args struct {
		filename string
		wantName string
	}

	files := []args{
		{"simple1.sql", "User"},
		{"simple2.sql", "A"},
		{"simple3.sql", "Tracking"},
		{"simple4.sql", "AbUser"},
	}

	for _, file := range files {
		readFile, err := testdata.TestData.ReadFile(file.filename)
		assert.NoError(t, err)
		s := NewSql()
		parse, err := s.Parse(string(readFile))
		assert.NoError(t, err)
		// test struct name
		assert.Equal(t, file.wantName, parse.StructName)
	}
}

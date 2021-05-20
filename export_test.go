package gexport

import (
	"fmt"
	"github.com/kingzcheung/gexport/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGExport_Export(t *testing.T) {
	//type fields struct {
	//	parse Parser
	//	sql   string
	//}
	//tests := []struct {
	//	name    string
	//	fields  fields
	//	want    []byte
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		e := &GExport{
	//			parse: tt.fields.parse,
	//			sql:   tt.fields.sql,
	//		}
	//		got, err := e.Export()
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("Export() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("Export() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	data, _ := testdata.TestData.ReadFile("simple3.sql")
	ge := New(string(data), SQL)

	_, err := ge.Export()
	assert.NoError(t, err)
	if err != nil {

		fmt.Println(err)
	}
}

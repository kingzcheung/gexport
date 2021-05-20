package gexport

import "testing"

func TestSql_FieldName(t *testing.T) {
	type fields struct {
		structName string
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{name: "", fields: fields{structName: ""}, args: args{"a_bc_df"}, want: "ABcDf"},
		{name: "", fields: fields{structName: ""}, args: args{"upper_camel_case"}, want: "UpperCamelCase"},
		{name: "", fields: fields{structName: ""}, args: args{"_upper_camel_case"}, want: "UpperCamelCase"},
		{name: "", fields: fields{structName: ""}, args: args{"_upper_camel_case_"}, want: "UpperCamelCase_"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SqlStruct{
				structName: tt.fields.structName,
			}
			if got := s.FieldName(tt.args.in); got != tt.want {
				t.Errorf("FieldName() = %v, want %v", got, tt.want)
			}
		})
	}
}

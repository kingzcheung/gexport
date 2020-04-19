package gexport

type Tag struct {
	Name   string
	Fields []map[string]string
}

//CreateJsonTag 生成json tag
func CreateJsonTag(field string) Tag {
	return CreateFieldTag("json", field)
}

//CreateFormTag 生成表单字段的tag
func CreateFormTag(field string) Tag {
	return CreateFieldTag("form", field)
}

func CreateXmlTag(field string) Tag {
	return CreateFieldTag("xml", field)
}

//CreateFieldTag 生成特定的字段 tag
func CreateFieldTag(name string, field string) Tag {
	var fields []map[string]string
	fields = append(fields, map[string]string{field: ""})
	return Tag{
		Name:   name,
		Fields: fields,
	}
}

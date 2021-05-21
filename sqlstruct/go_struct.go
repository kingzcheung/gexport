package sqlstruct

type Tag struct {
	TagName  string
	TagValue map[string]string
}

type Struct struct {
	StructName string
	Fields     []*StructField
}

type FieldName string

type FieldType string

type StructField struct {
	FieldName  string // 字段名
	FieldType  string // 字段类型
	Tags       []*Tag //字段标签
	Annotation string //字段注释
}

//func (fn FieldName) UpperCamelCase() string {
//	bs := []byte(fn)
//	for i, b := range bs {
//		if i == 0 {
//			if b >= 97 && b <= 122 {
//				bs[0] = b-32
//			}
//		}
//		if b == '_' && i < len(bs) - 1 {
//			// 小写变大写
//			bs[i+1] = bs[i+1] - 32
//
//			// 删除下划线 _
//			bs = append(bs[:i],bs[i+1:]... )
//		}
//	}
//	return string(bs)
//}

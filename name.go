package gexport

type FieldNamer interface {
	FieldName(in string) string
}

func (s *SqlStruct) FieldName(in string) string {
	bs := []byte(in)
	for i, b := range bs {
		if i == 0 {
			if b >= 97 && b <= 122 {
				bs[0] = b - 32
			}
		}
		if b == '_' && i < len(bs)-1 {
			// 小写变大写
			bs[i+1] = bs[i+1] - 32

			// 删除下划线 _
			bs = append(bs[:i], bs[i+1:]...)
		}
	}
	return string(bs)
}

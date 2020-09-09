package gexport

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	src string
	ok  bool
	out string
}

func TestSql_Parse(t *testing.T) {
	as := assert.New(t)
	// sql := `CREATE TABLE test (id int(11) NOT NULL AUTO_INCREMENT COMMENT '编号ID')`

	data := []testCase{
		{`create table a (id int(11) not null COMMENT '编号ID')`, true, "type A struct {\n\tId int `json:\"id\" form:\"id\" gorm:\"column:id;type:int(11)\" `\n}"},
		{`create table a (user_id varchar(255) not null COMMENT '')`, true, "type A struct {\n\tUserId string `json:\"user_id\" form:\"user_id\" gorm:\"column:user_id;type:varchar(255)\" `\n}"},
		{`create table a (user_id varchar(255) not null COMMENT '',username varchar(50) not null comment '用户')`, true, "type A struct {\n\tUserId string `json:\"user_id\" form:\"user_id\" gorm:\"column:user_id;type:varchar(255)\" `\n\tUsername string `json:\"username\" form:\"username\" gorm:\"column:username;type:varchar(50)\" `\n}"},
	}

	for _, testC := range data {
		ns := NewSql()
		p, err := ns.Parse(testC.src)
		if err != nil {
			as.Error(err)
		}
		as.Equal(p[0], testC.out)
	}

}

package main

import (
	"fmt"
	"github.com/kingzcheung/gexport/sqlstruct"
)

func example() {

	sql := `CREATE TABLE gd_goods_record (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  goods_sn varchar(20) NOT NULL COMMENT '商品编号',
  type tinyint(2) NOT NULL DEFAULT '1' COMMENT '记录类型1：物流信息，2：商品事件',
  create_time int(10) NOT NULL COMMENT '创建时间',
  system_user_id int(10) NOT NULL COMMENT '记录者id',
  old_value varchar(100) NOT NULL DEFAULT '-' COMMENT '旧值',
  new_value varchar(100) NOT NULL DEFAULT '-' COMMENT '新值',
  system_user varchar(50) NOT NULL COMMENT '修改者',
  event tinyint(2) NOT NULL DEFAULT '0' COMMENT '事件 ig_audit_event',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品物流，事件记录表';`

	gx := sqlstruct.New(sqlstruct.SQL)

	fmt.Println(gx.Export(sql))
}

func main() {
	example()
	// json_example()
}

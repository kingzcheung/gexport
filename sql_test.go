package gexport

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSql_Parse(t *testing.T) {
	as := assert.New(t)
	sql := `CREATE TABLE ig_reg_log (
	  log int(11) NOT NULL AUTO_INCREMENT COMMENT '编号ID',
	  user_id int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
	  reg_ip varchar(20) NOT NULL COMMENT '注册IP',
	  add_time int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',
	  party_reg tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1-手机验证码注册；2-邮箱注册; 3-第三方注册',
	  party_style tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '1-facekbook 2-wechat 3-QQ 4-Twitter 5-支付宝 7-google 8-line 9-ebbly 10-小程序',
	  reg_id varchar(255) DEFAULT '0' COMMENT '注册标识',
	  reg_source tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '注册来源(详情见ig_entrance表)',
	  reg_info varchar(255) DEFAULT '0' COMMENT '注册信息',
	  PRIMARY KEY (log),
	  KEY user_id (user_id),
	  KEY party_style (party_style),
	  KEY reg_source (reg_source)
	) ENGINE=InnoDB AUTO_INCREMENT=4777 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='注册日志'`

	ns := NewSql()
	p, err := ns.Parse(sql)
	if err != nil {
		as.Error(err)
	}
	fmt.Println(p)
}

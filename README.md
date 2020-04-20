#  gexport json/sql 转换struct 工具

根据 json/sql 手写 struct 是很痛苦的事情，通过 `gexport` 可快速生成 go struct，并且支持 `gorm` 标签

### 下载

```bash
go get github.com/kingzcheung/gexport
```

### 使用说明

#### 作为命令行使用

```bash
go build -o gexport/gexport gexport/main.go
```

gexport 通过 `stdin` 作为数据输入源来生成`struct`，因此支持多样的调用方式：

1. 只输入`gexport` ,默认会进入用户输入状态，通过输入json或者sql，再按 `cmd + d` 即可生成结果:

```bash
$ gexport 
{"foo":"bar"}
type RootGeneratedName struct {
        Foo string `json:"foo" `
}

```

2. 通过管道作为输入源:

```bash
cat data.json | gexport #通过cat命令获取文本并传给gexport
curl -s https://your.api.com/data.json | gexport #通过curl获取网络json数组并传给gexport
```

3. 通过输入重定向作为输入源:

```bash
gexport < data.sql #通过输入重定向把文件内容输入gexport
gexport < data.json > data.go # 也可以把结果重定向到文件
gexport --outfile=data.file < data.json # 或者通过--outfile 选项输出到文件
```

4. 自定义 `struct name`

```bash
gexport --name='User' <data.sql #通过--name 自定义struct name
```



#### 作为类库使用:

```go
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

	gx := gexport.New(sql, "sql")

	fmt.Println(gx.Parse().Output())
```

更多使用示例请参考 `_examples`目录。


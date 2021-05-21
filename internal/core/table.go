package core

import (
	"gorm.io/gorm"
)

type Table struct {
	Name string `json:"name"`
	Res  string `json:"res"`
}

type CreateTable struct {
	TableName   string `json:"table_name" gorm:"column:Table"`
	CreateTable string `json:"create_table" gorm:"column:Create Table"`
	StructRes   string `json:"struct_res" gorm:"-"`
}

func FetchTable(db *gorm.DB) []*Table {
	var tables []*Table

	db.Raw("show table status").Scan(&tables)

	return tables
}

func ShowCreateTable(db *gorm.DB, name string) *CreateTable {
	var ct CreateTable
	db.Raw("show create table " + name).Scan(&ct)
	return &ct
}

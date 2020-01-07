/**
 *
 * @Description:
 * @Author: https://www.bajins.com
 * @File: database.go
 * @Version: 1.0.0
 * @Time: 2020/1/7/007 12:54
 * @Project: directory-lister-echo
 * @Package:
 * @Software: GoLand
 */
package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// https://godoc.org/github.com/mattn/go-sqlite3#hdr-Supported_Types
// https://gorm.io/zh_CN/docs/conventions.html
type User struct {
	// gorm.Model 是一个包含了ID, CreatedAt, UpdatedAt, DeletedAt四个字段的GoLang结构体。
	gorm.Model
	//Id         int        `gorm:"type:integer;PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;NOT NULL;unique_index"`
	//CreateTime *time.Time `gorm:"type:datetime;DEFAULT:(DATETIME('NOW', 'LOCALTIME'))"`
	//ModifyTime *time.Time `gorm:"type:datetime"`
	Name   string `gorm:"type:text;UNIQUE;NOT NULL;unique_index"`
	Email  string `gorm:"type:text;UNIQUE;NOT NULL;unique_index"`
	Status bool   `gorm:"type:integer;UNIQUE;NOT NULL;DEFAULT:1"`
}

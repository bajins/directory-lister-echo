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
	"time"
)

type User struct {
	gorm.Model
	Id         int    `gorm:"type:integer;PRIMARY_KEY;AUTO_INCREMENT;UNIQUE;NOT NULL;unique_index"`
	Name       string `gorm:"type:varchar(100);UNIQUE;NOT NULL;unique_index"`
	Email      string `gorm:"type:varchar(100);UNIQUE;NOT NULL;unique_index"`
	CreateTime *time.Time
	ModifyTime *time.Time
	Status     int `gorm:"type:integer;UNIQUE;NOT NULL;DEFAULT:1"`
}

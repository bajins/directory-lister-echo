/**
 *
 * @Description:
 * @Author: https://www.bajins.com
 * @File: main_test.go
 * @Version: 1.0.0
 * @Time: 2020/1/7/007 15:37
 * @Project: directory-lister-echo
 * @Package:
 * @Software: GoLand
 */
package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
)

func TestDB(t *testing.T) {
	// https://gorm.io/zh_CN/docs/conventions.html
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// AutoMigrate为给定模型运行自动迁移，只会添加缺少的字段，不会删除/更改当前数据
	db = db.AutoMigrate(&User{})
	if db.Error != nil {
		panic("failed to create table")
	}

	// 创建
	db.Create(&User{Name: "L1212", Email: "test@test.com"})
	db.Commit()

	u := db.Exec("select * from user")
	t.Log(u.Rows())

	// 读取
	var user User
	db.Select(&user)
	t.Log(user)
	db.First(&user, 1) // 查询id为1的product
	t.Log(user)
	db.First(&user, "name = ?", "L1212") // 查询code为l1212的product
	t.Log(user)
}

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
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

	// 创建
	db.Create(&User{Name: "L1212", Email: "test@test.com"})

	// 读取
	var user User
	db.First(&user, 1)                   // 查询id为1的product
	db.First(&user, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&user).Update("Price", 2000)

	// 删除 - 删除product
	db.Delete(&user)
}

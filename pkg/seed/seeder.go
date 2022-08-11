// Package seed 处理数据库填充相关逻辑
package seed

import (
	"gorm.io/gorm"
)

type Seeder struct {
	Func SeederFunc
	Name string
}

type SeederFunc func(*gorm.DB)

//按顺序执行的Seeder 数组
//支持一些必须按顺序执行的 seeder, 例如 topic 创建时必须依赖于user,所以TopicSeeder 应该在UserSeeder后执行
var orderedSeederNames []string

//存放所有 Seeder
var seeders []Seeder

//Add 注册到 seeders 数据中
func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Func: fn,
		Name: name,
	})
}

//SetRunOrder 设置 按顺序执行的Seeder 数组
func SetRunOrder(names []string) {
	orderedSeederNames = names
}
